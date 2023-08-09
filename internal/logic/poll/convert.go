package poll

import (
	"github.com/hitokoto-osc/reviewer/internal/consts"
	"github.com/hitokoto-osc/reviewer/internal/dao"
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
