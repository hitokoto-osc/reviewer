package poll

import (
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
	"github.com/hitokoto-osc/reviewer/internal/service"
	"github.com/hitokoto-osc/reviewer/utility/time"
)

var pollMethodsMap = map[consts.PollMethod]string{
	consts.PollMethodApprove:            dao.Poll.Columns().Accept,
	consts.PollMethodReject:             dao.Poll.Columns().Reject,
	consts.PollMethodNeedModify:         dao.Poll.Columns().NeedEdited,
	consts.PollMethodNeedCommonUserPoll: dao.Poll.Columns().NeedUserPoll,
}

func translatePollMethodToField(in consts.PollMethod) string {
	return pollMethodsMap[in]
}

var pointsMap = map[consts.UserRole]consts.UserPollPoints{
	consts.UserRoleUser:     consts.UserPollPointsNormal,
	consts.UserRoleReviewer: consts.UserPollPointsReviewer,
	consts.UserRoleAdmin:    consts.UserPollPointsAdmin,
}

func (s *sPoll) GetPointsByRole(role consts.UserRole) consts.UserPollPoints {
	return pointsMap[role]
}

var pollStatusToMethodMap = map[consts.PollStatus]consts.PollMethod{
	consts.PollStatusApproved:   consts.PollMethodApprove,
	consts.PollStatusRejected:   consts.PollMethodReject,
	consts.PollStatusNeedModify: consts.PollMethodNeedModify,
}

func (s *sPoll) translatePollStatusToMethod(in consts.PollStatus) consts.PollMethod {
	return pollStatusToMethodMap[in]
}

func (s *sPoll) ConvertPollLogToPollRecord(ctx context.Context, in *entity.PollLog, isAdmin bool) (out *model.PollRecord, err error) {
	if in == nil {
		err = gerror.New("nil poll log")
		return
	}
	u, e := service.User().GetUserByID(ctx, uint(in.UserId))
	if e != nil {
		err = gerror.Wrap(e, "获取用户信息失败")
		return
	}
	if isAdmin {
		out = &model.PollRecord{
			User:      &model.UserPublicInfo{ID: u.Id, Name: u.Name, EmailHash: gmd5.MustEncryptString(u.Email)},
			Point:     in.Point,
			Method:    consts.PollMethod(in.Type),
			Comment:   in.Comment,
			CreatedAt: (*time.Time)(in.CreatedAt),
			UpdatedAt: (*time.Time)(in.UpdatedAt),
		}
	} else {
		out = &model.PollRecord{
			User:      &model.UserPublicInfo{ID: u.Id, Name: u.Name, EmailHash: gmd5.MustEncryptString(u.Email)},
			Comment:   in.Comment,
			CreatedAt: (*time.Time)(in.CreatedAt),
			UpdatedAt: (*time.Time)(in.UpdatedAt),
		}
	}
	return
}

func (s *sPoll) MustConvertPollLogToPollRecord(ctx context.Context, in *entity.PollLog, isAdmin bool) (out *model.PollRecord) {
	var err error
	out, err = s.ConvertPollLogToPollRecord(ctx, in, isAdmin)
	if err != nil {
		panic(err)
	}
	return out
}
