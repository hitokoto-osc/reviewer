// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hitokoto

import (
	"github.com/hitokoto-osc/reviewer/api/hitokoto"
)

type ControllerV1 struct{}

func NewV1() hitokoto.IHitokotoV1 {
	return &ControllerV1{}
}

type ControllerAdminV1 struct{}

func NewAdminV1() hitokoto.IHitokotoAdminV1 {
	return &ControllerAdminV1{}
}
