package notification

import (
	"context"
	"encoding/json"

	"golang.org/x/sync/errgroup"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hitokoto-osc/reviewer/internal/service"
	"github.com/rabbitmq/amqp091-go"
)

// DoNotification 发送通知
// TODO: 合并成一条消息，让消费者自行处理
func DoNotification(ctx context.Context, exchange, routingKey string, data []any) error {
	controller, err := service.AMQP().GetConnectionController()
	if err != nil {
		return err
	}
	defer controller.ReleaseConnection()
	conn, err := controller.GetConnection()
	if err != nil {
		return err
	}
	channel, err := conn.Channel()
	if err != nil {
		return gerror.Wrap(err, "获取 AMQP 通道失败")
	}
	defer channel.Close()
	eg, egCtx := errgroup.WithContext(ctx)
	for _, v := range data {
		value := v
		eg.Go(func() error {
			buff, e := json.Marshal(value)
			if e != nil {
				return e
			}
			return channel.PublishWithContext(egCtx, exchange, routingKey, false, false, amqp091.Publishing{
				ContentType:     "application/json",
				ContentEncoding: "utf-8",
				DeliveryMode:    2,
				Timestamp:       gtime.Now().Time,
				Body:            buff,
			})
		})
	}
	err = eg.Wait()
	if err != nil {
		return gerror.Wrap(err, "发送消息失败")
	}
	return nil
}
