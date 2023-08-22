package user

import (
	"context"
	"time"

	"github.com/hitokoto-osc/reviewer/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/hitokoto-osc/reviewer/internal/dao"
)

// GetReviewersAndAdmins 获取所有需要发送通知的用户
// 目前只是一次性获取管理员和审核员
func (s *sUser) GetReviewersAndAdmins(ctx context.Context) (users []entity.Users, err error) {
	err = dao.Users.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: time.Minute * 10, // 缓存十分钟
		Name:     "users:reviewersAndAdmins",
		Force:    false,
	}).
		WhereOr(dao.Users.Columns().IsReviewer, 1). // 审核员
		WhereOr(dao.Users.Columns().IsAdmin, 1).    // 管理员
		Scan(&users)
	return
}
