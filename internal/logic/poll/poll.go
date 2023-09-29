package poll

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"

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

func (s *sPoll) GetPollsBySentenceUUIDs(ctx context.Context, sentenceUUIDs []string) (polls []entity.Poll, err error) {
	err = dao.Poll.Ctx(ctx).Where(dao.Poll.Columns().SentenceUuid, sentenceUUIDs).Scan(&polls)
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
		affectedRows, err := dao.Pending.Ctx(ctx).TX(tx).Unscoped().
			Where(dao.Pending.Columns().Id, pending.Id).
			Data(g.Map{
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
	service.Cache().ClearCacheAfterPollUpdated(ctx, 0, uint(poll.Id), poll.SentenceUuid)
	return poll, nil
}

//nolint:gocyclo
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
	if in.PolledFilter != consts.PollListPolledFilterAll {
		filter := ""
		if in.PolledFilter == consts.PollListPolledFilterReviewed {
			filter = "NOT"
		}
		sql := fmt.Sprintf(`
SELECT
	poll.*
FROM
	%s poll
	LEFT JOIN %s log ON poll.%s = log.%s  AND log.%s = %d
WHERE
	log.%s IS %s NULL AND
	poll.%s = 1
`,
			dao.Poll.Table(),
			dao.PollLog.Table(),
			dao.Poll.Columns().Id,
			dao.PollLog.Columns().PollId,
			dao.PollLog.Columns().UserId,
			in.UserID,
			dao.PollLog.Columns().Id,
			filter,
			dao.Poll.Columns().Status,
		)
		query = query.Raw(sql)
	} else if in.StatusStart == in.StatusEnd || in.StatusEnd == 0 { // 之后才判断状态
		query = query.Where(dao.Poll.Columns().Status, in.StatusStart)
	} else {
		query = query.WhereBetween(dao.Poll.Columns().Status, in.StatusStart, in.StatusEnd)
	}
	eg, egCtx := errgroup.WithContext(ctx)
	query.Ctx(egCtx)
	eg.Go(func() error {
		var e error
		q := query.Clone()
		count, e = q.Count()
		g.Log().Debugf(ctx, "count: %d", count)
		return e
	})
	eg.Go(func() error {
		q := query.Clone()
		if in.WithCache {
			name := ""
			if in.PolledFilter == consts.PollListPolledFilterAll {
				name = fmt.Sprintf(
					"poll:list:status:%d:%d:page:%d:limit:%d:%s",
					in.Page,
					in.PageSize,
					in.StatusStart,
					in.StatusEnd,
					in.Order,
				)
			} else {
				name = fmt.Sprintf(
					"poll:list:reviewed_filter:%d:user:%d:status:%d:%d:page:%d:limit:%d:%s",
					in.PolledFilter,
					in.UserID,
					in.Page,
					in.PageSize,
					in.StatusStart,
					in.StatusEnd,
					in.Order,
				)
			}
			q = q.Cache(gdb.CacheOption{
				Duration: time.Minute * 10,
				Name:     name,
				Force:    false,
			})
		}
		e := q.Order(dao.Poll.Columns().CreatedAt, in.Order).Page(in.Page, in.PageSize).Scan(&result)
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
		if in.WithPollRecords {
			eg.Go(func() error {
				records, e := service.Poll().GetPollLogsByPollID(egCtx, value.Id)
				if e != nil {
					return e
				} else if len(records) == 0 {
					collection[index].Records = make([]model.PollRecord, 0)
					return nil
				}
				users := service.BizCtx().GetUser(ctx)
				isAdmin := users != nil && users.Role == consts.UserRoleAdmin
				collection[index].Records = make([]model.PollRecord, 0, len(records))
				for recordIndex, record := range records {
					if uint(record.UserId) == in.UserID {
						collection[index].PolledData = &model.PolledData{
							Point:     record.Point,
							Method:    consts.PollMethod(record.Type),
							CreatedAt: (*vtime.Time)(record.CreatedAt),
							UpdatedAt: (*vtime.Time)(record.UpdatedAt),
						}
					}
					if isAdmin || record.Comment != "" {
						collection[index].Records = append(collection[index].Records, *s.MustConvertPollLogToPollRecord(ctx, &records[recordIndex], isAdmin))
					}
				}
				return nil
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
func (s *sPoll) Poll(ctx context.Context, in *model.PollInput) error {
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
			list := make([]do.PollMarkRelation, 0, len(in.Marks))
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
			// fmt.Printf("%+v\n", list)
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
				affectedRows, err := dao.PollUsers.Ctx(egCtx).TX(tx).Where(dao.PollUsers.Columns().UserId, in.UserID).
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
	// Remove caches
	service.Cache().ClearCacheAfterPollUpdated(ctx, in.UserID, in.PollID, in.SentenceUUID)
	return nil
}

func (s *sPoll) CancelPollByID(ctx context.Context, in *model.CancelPollInput) error {
	if in == nil {
		return gerror.New("input is empty")
	}
	err := g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		eg, egCtx := errgroup.WithContext(ctx)
		// 更新投票信息
		eg.Go(func() error {
			affectedRows, e := dao.Poll.Ctx(egCtx).TX(tx).Where(dao.Poll.Columns().Id, in.PollID).
				Data(g.Map{
					translatePollMethodToField(in.PolledData.Method): gdb.Raw(
						fmt.Sprintf("%s-%d",
							translatePollMethodToField(in.PolledData.Method),
							in.PolledData.Point,
						),
					),
				}).
				UpdateAndGetAffected()
			if e != nil {
				return e
			} else if affectedRows == 0 {
				return gerror.New("failed to update poll because poll not found")
			}
			return nil
		})
		// 更新用户投票信息
		if in.PolledData.Method != consts.PollMethodNeedCommonUserPoll {
			eg.Go(func() error {
				affectedRows, e := dao.PollUsers.Ctx(egCtx).TX(tx).Where(dao.PollUsers.Columns().UserId, in.UserID).
					Data(g.Map{
						dao.PollUsers.Columns().Points: gdb.Raw(
							fmt.Sprintf("%s-%d",
								dao.PollUsers.Columns().Points,
								in.PolledData.Point,
							),
						),
						translatePollMethodToField(in.PolledData.Method): gdb.Raw(
							fmt.Sprintf("%s-%d",
								translatePollMethodToField(in.PolledData.Method),
								in.PolledData.Point,
							),
						),
					}).
					UpdateAndGetAffected()
				if e != nil {
					return e
				} else if affectedRows == 0 {
					return gerror.New("failed to update poll because poll not found")
				}
				return nil
			})
		}
		// 删除投票标记
		eg.Go(func() error {
			_, e := dao.PollMarkRelation.Ctx(egCtx).TX(tx).Where(dao.PollMarkRelation.Columns().UserId, in.UserID).
				Where(dao.PollMarkRelation.Columns().PollId, in.PollID).
				Delete()
			return e
		})
		// 删除投票日记
		eg.Go(func() error {
			_, e := dao.PollLog.Ctx(egCtx).TX(tx).Where(dao.PollLog.Columns().UserId, in.UserID).
				Where(dao.PollLog.Columns().PollId, in.PollID).
				Where(dao.PollLog.Columns().Type, in.PolledData.Method).
				Delete()
			return e
		})
		return eg.Wait()
	})
	if err != nil {
		return gerror.Wrap(err, "cancel poll failed")
	}
	service.Cache().ClearCacheAfterPollUpdated(ctx, in.UserID, in.PollID, in.SentenceUUID)
	return nil
}

func (s *sPoll) CountUserUnreviewedPoll(ctx context.Context, uid uint) (int, error) {
	record, err := dao.Poll.Ctx(ctx).
		Cache(gdb.CacheOption{
			Duration: time.Minute * 10,
			Name:     fmt.Sprintf("user:poll:unreviewed:uid:%d:count", uid),
		}).
		Raw(
			fmt.Sprintf(`
SELECT
	COUNT(1)
FROM
	%s poll
	LEFT JOIN %s log ON poll.%s = log.%s  AND log.%s = %d
WHERE
	log.%s IS NULL AND
	poll.%s = 1`,
				dao.Poll.Table(),
				dao.PollLog.Table(),
				dao.Poll.Columns().Id,
				dao.PollLog.Columns().PollId,
				dao.PollLog.Columns().UserId,
				uid,
				dao.PollLog.Columns().Id,
				dao.Poll.Columns().Status,
			),
		).Value()
	if err != nil {
		return 0, err
	}
	return record.Int(), nil
}

func (s *sPoll) GetPollDetailByID(ctx context.Context, pollID uint, withPolledData bool) (*model.PollListElement, error) {
	poll, err := service.Poll().GetPollByID(ctx, int(pollID))
	if err != nil {
		return nil, gerror.Wrap(err, "获取投票失败")
	} else if poll == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}
	user := service.BizCtx().GetUser(ctx)
	// fetch logs and sentence
	var (
		logs       []entity.PollLog
		sentence   *model.HitokotoV1WithPoll
		marks      []int
		polledData *model.PolledData
	)
	eg, egCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		var e error
		logs, e = service.Poll().GetPollLogsByPollID(egCtx, poll.Id)
		if e != nil {
			return gerror.WrapCode(gcode.CodeOperationFailed, e, "获取投票日志失败")
		}
		return nil
	})
	eg.Go(func() error {
		var e error
		sentence, e = service.Hitokoto().GetHitokotoV1SchemaByUUID(egCtx, poll.SentenceUuid)
		if e != nil {
			return gerror.WrapCode(gcode.CodeOperationFailed, e, "获取句子失败")
		}
		return nil
	})
	eg.Go(func() error {
		var e error
		marks, e = service.Poll().GetPollMarksByPollID(egCtx, uint(poll.Id))
		if e != nil {
			return gerror.WrapCode(gcode.CodeOperationFailed, e, "获取投票标签失败")
		}
		return nil
	})
	if withPolledData {
		eg.Go(func() error {
			var e error
			polledData, e = service.User().GetUserPolledDataWithPollID(egCtx, service.BizCtx().GetUser(egCtx).Id, uint(poll.Id))
			if e != nil {
				return gerror.WrapCode(gcode.CodeOperationFailed, e, "获取投票信息失败")
			}
			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		return nil, err
	}
	var records []model.PollRecord
	if len(logs) > 0 && (user.Role == consts.UserRoleAdmin || poll.Status != int(consts.PollStatusOpen)) {
		records = make([]model.PollRecord, len(logs))
		for i, log := range logs {
			u, e := service.User().GetUserByID(ctx, uint(log.UserId))
			if e != nil {
				return nil, gerror.WrapCode(gcode.CodeOperationFailed, e, "获取用户信息失败")
			}
			records[i] = model.PollRecord{
				User: &model.UserPublicInfo{
					ID:        uint(log.UserId),
					Name:      u.Name,
					EmailHash: gmd5.MustEncryptString(u.Email),
				},
				Point:     log.Point,
				Method:    consts.PollMethod(log.Type),
				Comment:   log.Comment,
				CreatedAt: (*vtime.Time)(log.CreatedAt),
				UpdatedAt: (*vtime.Time)(log.UpdatedAt),
			}
		}
	} else {
		records = []model.PollRecord{}
	}
	return &model.PollListElement{
		PollElement: model.PollElement{
			ID:                 uint(poll.Id),
			SentenceUUID:       poll.SentenceUuid,
			Sentence:           sentence,
			Status:             consts.PollStatus(poll.Status),
			Approve:            poll.Accept,
			Reject:             poll.Reject,
			NeedModify:         poll.NeedEdited,
			NeedCommonUserPoll: poll.NeedUserPoll,
			CreatedAt:          (*vtime.Time)(poll.CreatedAt),
			UpdatedAt:          (*vtime.Time)(poll.UpdatedAt),
		},
		Marks:      marks,
		PolledData: polledData,
		Records:    records,
	}, nil
}
