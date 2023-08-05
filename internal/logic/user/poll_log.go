package user

import (
	"context"
	"time"

	"github.com/hitokoto-osc/reviewer/internal/model/do"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hitokoto-osc/reviewer/internal/dao"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/internal/service"
)

func (s *sUser) GetUserPollLogByUserID(ctx context.Context, userID uint) (res []entity.PollLog, err error) {
	err = dao.PollLog.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Second * 3, // 3s
		Name:     "poll_log:uid:" + gconv.String(userID),
		Force:    false,
	}).Where(do.PollLog{UserId: userID}).Scan(&res)
	return
}

func (s *sUser) GetUserPollLog(ctx context.Context) (res []entity.PollLog, err error) {
	bizctx := service.BizCtx().Get(ctx)
	if bizctx == nil || bizctx.User == nil {
		err = gerror.New("bizctx or bizctx.User is nil")
		return
	}
	user := bizctx.User
	return s.GetUserPollLogByUserID(ctx, user.Id)
}
