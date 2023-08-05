package consts

// PollMethod 审核员投票的类型
type PollMethod int

const (
	PollMethodApprove            PollMethod = 1 // 赞同
	PollMethodReject             PollMethod = 2 // 驳回
	PollMethodNeedModify         PollMethod = 3 // 需要修改
	PollMethodNeedCommonUserPoll PollMethod = 4 // 需要普通用户参与投票
)

// PollStatus 投票状态
type PollStatus int

const (
	PollStatusOpen              PollStatus = 1   // 投票正常开放
	PollStatusProcessing        PollStatus = 2   // 处理中，停止投票
	PollStatusSuspended         PollStatus = 100 // 暂停投票
	PollStatusClosed            PollStatus = 101 // 关闭投票
	PollStatusOpenForCommonUser PollStatus = 102 // 开放给普通用户投票
	PollStatusApproved          PollStatus = 200 // 赞同
	PollStatusRejected          PollStatus = 201 // 驳回
	PollStatusNeedModify        PollStatus = 202 // 需要修改
)

// PollMarkLevel 投票标记的等级，用于前端显示亮色提醒
type PollMarkLevel string

const (
	PollMarkLevelInfo    PollMarkLevel = "info"
	PollMarkLevelWarning PollMarkLevel = "warning"
	PollMarkLevelDanger  PollMarkLevel = "danger"
)

// PollMarkProperty 投票标记的属性，目前无实际作用
type PollMarkProperty int

const (
	PollMarkPropertyApprove PollMarkProperty = 1 // 批准时可以使用的标记
	PollMarkPropertyReject  PollMarkProperty = 2 // 驳回时可以使用的标记
	PollMarkPropertyModify  PollMarkProperty = 3 // 需要修改时可以使用的标记
)
