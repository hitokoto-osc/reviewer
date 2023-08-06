package user

import (
	context "context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model/do"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
)

func (s *sUser) GetPollUserByUserID(ctx context.Context, uid uint) (user *entity.PollUsers, err error) {
	err = dao.PollUsers.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Hour, // 缓存一小时
		Name:     "user:poll:uid:" + gconv.String(uid),
		Force:    false,
	}).Where(do.PollUsers{UserId: uid}).Scan(&user)
	return
}

func (s *sUser) CreatePollUser(ctx context.Context, uid uint) (err error) {
	_, err = dao.PollUsers.Ctx(ctx).Insert(do.PollUsers{
		UserId:     uid,
		Points:     0,
		Accept:     0,
		Reject:     0,
		NeedEdited: 0,
		Score:      0,
	})
	return
}
