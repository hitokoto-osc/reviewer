package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hitokoto-osc/reviewer/utility/time"
)

// GetApplyTokenReq 获得申请 Token，主要用于审查员的申请时的文档确认
type GetApplyTokenReq struct {
	g.Meta `path:"/user/apply/token" tags:"User" method:"get" summary:"获得申请 Token"`
}

type GetApplyTokenRes struct {
	VerificationToken string     `json:"verification_token" dc:"验证令牌"`
	ValidUntil        *time.Time `json:"valid_until" dc:"有效期"`
}

// ApplyReviewerReq 申请成为审查员，目前说白了就是自动激活
type ApplyReviewerReq struct {
	g.Meta            `path:"/user/apply/reviewer" tags:"User" method:"post" summary:"申请成为审查员"`
	VerificationToken string `json:"verification_token" dc:"验证令牌" v:"required|size:36#验证令牌必须为 36 位"`
}

type ApplyReviewerRes struct{}
