package poll

import (
	"context"
	"time"

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

func (s *sPoll) GetPollBySentenceUUID(ctx context.Context, uuidStr string) (poll *entity.Poll, err error) {
	err = dao.Poll.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Minute * 10,
		Name:     "poll:uuid:" + uuidStr,
		Force:    false,
	}).Where(do.Poll{SentenceUuid: uuidStr}).Scan(&poll)
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
