package poll

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/util/gconv"

	vtime "github.com/hitokoto-osc/reviewer/utility/time"
	"golang.org/x/sync/errgroup"

	"github.com/hitokoto-osc/reviewer/internal/model"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/google/uuid"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/hitokoto-osc/reviewer/internal/consts"

	"github.com/hitokoto-osc/reviewer/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/do"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

type sPoll struct{}

func init() {
	service.RegisterPoll(New())
}

func New() service.IPoll {
	return &sPoll{}
}

func (s *sPoll) GetPollByID(ctx context.Context, pid int) (poll *entity.Poll, err error) {
	err = dao.Poll.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Minute * 10,
		Name:     "poll:id:" + gconv.String(pid),
		Force:    false,
	}).Where(do.Poll{Id: pid}).Scan(&poll)
	return
}

// GetPollBySentenceUUID 根据 Sentence UUID 获取最新发起的投票
func (s *sPoll) GetPollBySentenceUUID(ctx context.Context, sentenceUUID string) (poll *entity.Poll, err error) {
	err = dao.Poll.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Minute * 10,
		Name:     "poll:sentence_uuid:" + sentenceUUID,
		Force:    false,
	}).Where(do.Poll{SentenceUuid: sentenceUUID}).OrderDesc(dao.Poll.Columns().CreatedAt).Scan(&poll)
	return
}

func (s *sPoll) CountOpenedPoll(ctx context.Context) (int, error) {
	return dao.Poll.Ctx(ctx).Where(dao.Poll.Columns().Status, consts.PollStatusOpen).Count()
}

func (s *sPoll) CreatePollByPending(ctx context.Context, pending *entity.Pending) (*entity.Poll, error) {
	if pending == nil {
		return nil, gerror.New("pending is nil")
	}
	var poll *entity.Poll
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if pending.Uuid == "" {
			uuidInstance, err := uuid.NewRandom()
			if err != nil {
				return nil
			}
			pending.Uuid = uuidInstance.String()
		}
		// 修改 pending 状态
		affectedRows, err := dao.Pending.Ctx(ctx).TX(tx).Where(dao.Pending.Columns().Id, pending.Id).Data(g.Map{
			dao.Pending.Columns().PollStatus: consts.PollStatusOpen,
			dao.Pending.Columns().Uuid:       pending.Uuid,
		}).UpdateAndGetAffected()
		if err != nil {
			return err
		} else if affectedRows == 0 {
			return gerror.New("pending affectedRows is 0")
		}
		// 创建投票
		poll = &entity.Poll{
			SentenceUuid:   pending.Uuid,
			Status:         int(consts.PollStatusOpen),
			IsExpandedPoll: 0,
			Accept:         0,
			Reject:         0,
			NeedEdited:     0,
			NeedUserPoll:   0,
			PendingId:      pending.Id,
			CreatedAt:      gtime.Now(),
			UpdatedAt:      gtime.Now(),
		}
		insertedID, err := dao.Poll.Ctx(ctx).TX(tx).Data(poll).InsertAndGetId()
		if err != nil {
			return err
		}
		poll.Id = int(insertedID)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return poll, nil
}

func (s *sPoll) GetPollList(ctx context.Context, in *model.GetPollListInput) (*model.GetPollListOutput, error) {
	if in == nil {
		return nil, gerror.New("input is nil")
	}
	if in.Order == "" {
		in.Order = "asc"
	}
	query := dao.Poll.Ctx(ctx)
	var (
		result []entity.Poll
		count  int
	)
	// fetch poll list
	query = query.WhereBetween(dao.Poll.Columns().Status, in.StatusStart, in.StatusEnd).
		Order(dao.Poll.Columns().CreatedAt, in.Order).
		Page(in.Page, in.PageSize)
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		q := query.Clone()
		if in.WithCache {
			q = q.Cache(gdb.CacheOption{
				Duration: time.Minute * 10,
				Name:     fmt.Sprintf("poll:list:status:%d:%d:page:%d:limit:%d:%s", in.Page, in.PageSize, in.StatusStart, in.StatusEnd, in.Order),
				Force:    false,
			})
		}
		e := q.Scan(&result)
		return e
	})
	eg.Go(func() error {
		var e error
		q := query.Clone()
		count, e = q.Count(&count)
		return e
	})
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	// fetch user poll data, marks and sentence info
	eg, egCtx = errgroup.WithContext(ctx)
	collection := make([]model.PollListElement, len(result))
	for i, v := range result {
		index, value := i, v
		collection[index] = model.PollListElement{
			PollElement: model.PollElement{
				ID:                 uint(value.Id),
				SentenceUUID:       value.SentenceUuid,
				Status:             consts.PollStatus(value.Status),
				Approve:            value.Accept,
				Reject:             value.Reject,
				NeedModify:         value.NeedEdited,
				NeedCommonUserPoll: value.NeedUserPoll,
				CreatedAt:          (*vtime.Time)(v.CreatedAt),
				UpdatedAt:          (*vtime.Time)(v.UpdatedAt),
			},
		}
		// 获取句子
		eg.Go(func() error {
			var e error
			collection[index].Sentence, e = service.Hitokoto().GetHitokotoV1SchemaByUUID(egCtx, value.SentenceUuid)
			return e
		})
		if in.WithMarks {
			eg.Go(func() error {
				var e error
				collection[index].Marks, e = s.GetPollMarksByPollID(egCtx, uint(value.Id))
				return e
			})
		}
		if in.WithUserPolledData {
			eg.Go(func() error {
				var e error
				collection[index].PolledData, e = service.User().GetUserPolledDataWithPollID(egCtx, in.UserID, uint(value.Id))
				return e
			})
		}
	}
	err := eg.Wait()
	if err != nil {
		return nil, err
	}
	out := &model.GetPollListOutput{
		Collection: collection,
		Total:      count,
		Page:       in.Page,
		PageSize:   in.PageSize,
	}
	return out, nil
}

// Poll 投票
func Poll(ctx context.Context, in *model.PollInput) error {
	if in == nil {
		return gerror.New("input is empty")
	}
	if in.PollID == 0 {
		return gerror.New("poll id is empty")
	}
	if in.UserID == 0 {
		user := service.BizCtx().GetUser(ctx)
		if user == nil {
			return gerror.New("user is empty")
		}
		in.UserID = user.Id
		in.IsAdmin = user.Role == consts.UserRoleAdmin
	}
	updatedAt := gtime.Now()
	createdAt := updatedAt
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		eg, egCtx := errgroup.WithContext(ctx)
		// 提交投票标记
		if len(in.Marks) > 0 {
			list := make([]do.PollMarkRelation, len(in.Marks))
			for _, v := range in.Marks {
				list = append(list, do.PollMarkRelation{
					UserId:       int(in.UserID),
					PollId:       int(in.PollID),
					SentenceUuid: in.SentenceUUID,
					MarkId:       v,
					UpdatedAt:    updatedAt,
					CreatedAt:    createdAt,
				})
			}
			eg.Go(func() error {
				_, e := dao.PollMarkRelation.Ctx(egCtx).TX(tx).Data(list).Insert()
				return e
			})
		}
		// 更新投票数据
		eg.Go(func() error {
			affectedRows, err := dao.Poll.Ctx(egCtx).TX(tx).Where(dao.Poll.Columns().Id, in.PollID).
				Data(g.Map{
					translatePollMethodToField(in.Method): gdb.Raw(fmt.Sprintf("%s+%d", translatePollMethodToField(in.Method), in.Point)),
					dao.Poll.Columns().UpdatedAt:          updatedAt,
				}).
				UpdateAndGetAffected()
			if err != nil {
				return err
			} else if affectedRows == 0 {
				return gerror.New("failed to update poll because poll not found")
			}
			return nil
		})
		// 更新用户投票数据
		if in.Method != consts.PollMethodNeedCommonUserPoll {
			eg.Go(func() error {
				affectedRows, err := dao.PollUsers.Ctx(egCtx).TX(tx).Where(dao.PollUsers.Columns().Id, in.UserID).
					Data(g.Map{
						dao.PollUsers.Columns().Points:        gdb.Raw(fmt.Sprintf("%s+%d", dao.PollUsers.Columns().Points, in.Point)),
						translatePollMethodToField(in.Method): gdb.Raw(fmt.Sprintf("%s+%d", translatePollMethodToField(in.Method), in.Point)),
						dao.PollUsers.Columns().UpdatedAt:     updatedAt,
					}).
					UpdateAndGetAffected()
				if err != nil {
					return err
				} else if affectedRows == 0 {
					return gerror.New("failed to update user because user not found")
				}
				return nil
			})
		}

		// 记录用户日记
		eg.Go(func() error {
			isAdmin := 0
			if in.IsAdmin {
				isAdmin = 1
			}
			_, err := dao.PollLog.Ctx(ctx).TX(tx).Data(do.PollLog{
				PollId:       int(in.PollID),
				UserId:       int(in.UserID),
				Point:        in.Point,
				SentenceUuid: in.SentenceUUID,
				Type:         int(in.Method),
				Comment:      in.Comment,
				IsAdmin:      isAdmin,
				CreatedAt:    createdAt,
				UpdatedAt:    updatedAt,
			}).Insert()
			return err
		})
		err := eg.Wait()
		return err
	})
	if err != nil {
		return gerror.Wrap(err, "poll failed")
	}
	return nil
}
