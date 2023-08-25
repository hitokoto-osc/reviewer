package consts

const (
	AppName      = "@hitokoto/reviewer"
	Version      = "v2.0.0"
	APIVersionV1 = "v1.1.0"
	ContextKey   = "ContextKey" // 上下文变量存储键名，前后端系统共享

	TimeFormat = "c" // 包装 gtime.Time 使用的格式化字符串

	ReviewInactivityThreshold = 30 // 30 天

	PollWinnerScore      = 3 // 投票胜利者获得的积分
	PollParticipantScore = 1 // 投票参与者获得的积分

	PollRulingUserID                         = 4756
	PollRulingUsername                       = "众裁委员会"
	PollRulingInitThreshold                  = 5
	PollRulingNormalThreshold                = 10 // 10 票
	PollRulingNormalRate                     = 0.6
	PollRulingNeedForCommonUserPollThreshold = 20 // 20 票
	PollRulingNeedForCommonUserPollRate      = 0.5
	PollOverdueThreshold                     = 7 // 7 天

	ReviewerInactiveDays = 30 // 30 天
)
