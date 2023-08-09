package poll

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/dao"
	"github.com/hitokoto-osc/reviewer/internal/model"
	"github.com/hitokoto-osc/reviewer/internal/model/entity"
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

func (s *sPoll) ConvertPollLogToPollRecord(in *entity.PollLog, isAdmin bool) (out *model.PollRecord, err error) {
	if in == nil {
		err = gerror.New("nil poll log")
		return
	}
	if isAdmin {
		out = &model.PollRecord{
			UserID:    uint(in.UserId),
			Point:     in.Point,
			Method:    consts.PollMethod(in.Type),
			Comment:   in.Comment,
			CreatedAt: (*time.Time)(in.CreatedAt),
			UpdatedAt: (*time.Time)(in.UpdatedAt),
		}
	} else {
		out = &model.PollRecord{
			UserID:    uint(in.UserId),
			Comment:   in.Comment,
			CreatedAt: (*time.Time)(in.CreatedAt),
			UpdatedAt: (*time.Time)(in.UpdatedAt),
		}
	}
	return
}

func (s *sPoll) MustConvertPollLogToPollRecord(in *entity.PollLog, isAdmin bool) (out *model.PollRecord) {
	var err error
	out, err = s.ConvertPollLogToPollRecord(in, isAdmin)
	if err != nil {
		panic(err)
	}
	return out
}
