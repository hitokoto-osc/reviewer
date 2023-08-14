package model

type DailyReportNotificationMessage struct {
	CreatedAt         string                           `json:"created_at"`         // 报告生成时间
	To                string                           `json:"to"`                 // 接收人地址
	UserName          string                           `json:"user_name"`          // 接收人名称
	SystemInformation pollDailyReportSystemInformation `json:"system_information"` // 系统信息
	UserInformation   pollDailyReportUserInformation   `json:"user_information"`   // 用户信息
}

type pollDailyReportSystemInformation struct {
	Total             int `json:"total"`               // 平台当前剩余的投票数目
	ProcessTotal      int `json:"process_total"`       // 平台处理了的投票数目（过去 24 小时）
	ProcessAccept     int `json:"process_accept"`      // 平台处理为入库的投票数目（过去 24 小时）
	ProcessReject     int `json:"process_reject"`      // 平台处理为驳回的投票数目（过去 24 小时）
	ProcessNeedEdited int `json:"process_need_edited"` // 平台处理为亟待修改的投票数目（过去 24 小时）
}

type pollDailyReportUserInformation struct {
	Polled         pollDailyReportMessageUserInformationPolled `json:"polled"`           // 用户参与了的投票数目（过去 24 小时）
	Waiting        int                                         `json:"waiting"`          // 等待其他用户参与的投票数目（基于已投票的数目）
	Accepted       int                                         `json:"accepted"`         // 已入库的投票数目（基于已投票的数目）
	Rejected       int                                         `json:"rejected"`         // 已驳回的投票数目（基于已投票的数目）
	InNeedEdited   int                                         `json:"in_need_edited"`   // 已进入亟待修改状态的投票数目（基于已投票的数目）
	WaitForPolling int                                         `json:"wait_for_polling"` // 基于剩余投票数目，计算出来的等待投票数目。
}

type pollDailyReportMessageUserInformationPolled struct {
	Total      int `json:"total"`       // 投票参与的总数
	Accept     int `json:"accept"`      // 投批准票的投票数目
	Reject     int `json:"reject"`      // 投拒绝票的投票数目
	NeedEdited int `json:"need_edited"` // 投需要修改的投票数目
}
