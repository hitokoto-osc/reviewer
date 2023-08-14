// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/hitokoto-osc/reviewer/utility/amqp"
)

type (
	IAMQP interface {
		GetConnectionController() (amqp.ConnectionController, error)
	}
)

var (
	localAMQP IAMQP
)

func AMQP() IAMQP {
	if localAMQP == nil {
		panic("implement not found for interface IAMQP, forgot register?")
	}
	return localAMQP
}

func RegisterAMQP(i IAMQP) {
	localAMQP = i
}
