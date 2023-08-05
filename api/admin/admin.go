// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	v1 "github.com/hitokoto-osc/reviewer/api/admin/v1"
)

type IAdminV1 interface {
	GrantUserAuthority(ctx context.Context, req *v1.GrantUserAuthorityReq) (res *v1.GrantUserAuthorityRes, err error)
	ModifySentence(ctx context.Context, req *v1.ModifySentenceReq) (res *v1.ModifySentenceRes, err error)
}
