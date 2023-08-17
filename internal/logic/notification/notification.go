package notification

import "github.com/hitokoto-osc/reviewer/internal/service"

type sNotification struct{}

func init() {
	service.RegisterNotification(New())
}

func New() service.INotification {
	return &sNotification{}
}
