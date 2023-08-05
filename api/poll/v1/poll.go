package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/internal/consts"
)

type PollData struct {
	Point  int               `json:"point" dc:"投票点数"`
	Method consts.PollMethod `json:"method" dc:"投票方式"`
}

type PollSchema struct {
	SentenceUUID       string            `json:"sentence_uuid" dc:"句子 UUID"`
	Status             consts.PollStatus `json:"status" dc:"投票状态"`
	Accept             int               `json:"accept" dc:"赞同票数"`
	Reject             int               `json:"reject" dc:"反对票数"`
	NeedModify         int               `json:"need_edited" dc:"需要修改票数"`
	NeedCommonUserPoll int               `json:"need_common_user_poll" dc:"需要普通用户投票"`
	CreatedAt          string            `json:"created_at" dc:"创建时间"`
	UpdatedAt          string            `json:"updated_at" dc:"更新时间"`
}

type PollReq struct {
	g.Meta       `path:"/poll" tags:"Poll" method:"put" summary:"投票"`
	Method       consts.PollMethod `json:"method" dc:"投票方式" v:"required|enums"`
	SentenceUUID string            `json:"sentence_uuid" dc:"句子 UUID" v:"required|length:36"`
	MarkID       uint              `json:"mark_id" dc:"标记" v:"required-unless:method,1"`
	Comment      string            `json:"comment" dc:"理由" v:"required-if:method,2,3|length:1,255"`
}

// PollRes 成功返回句子的投票记录
type PollRes struct {
	PollSchema
	PollData PollData `json:"poll_data" dc:"投票数据"`
}
