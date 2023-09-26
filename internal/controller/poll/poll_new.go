// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package poll

import (
	"github.com/hitokoto-osc/reviewer/api/poll"
)

type ControllerV1 struct{}

func NewV1() poll.IPollV1 {
	return &ControllerV1{}
}

type ControllerAdminV1 struct{}

func NewAdminV1() poll.IPollAdminV1 {
	return &ControllerAdminV1{}
}
