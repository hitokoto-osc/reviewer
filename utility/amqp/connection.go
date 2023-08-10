package amqp

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/rabbitmq/amqp091-go"
)

// IConnectionController 连接控制器接口
type IConnectionController interface {
	GetConnection() *amqp091.Connection           // 获取原始连接
	Reconnect() error                             // 重新连接
	IsClosed() bool                               // 是否已经关闭
	IsError() bool                                // 是否有错误
	GetError() error                              // 获取错误
	Close() error                                 // 关闭连接
	EstablishedTime() *gtime.Time                 // 获取连接建立时间
	IsIdle() bool                                 // 检查连接是否空闲
	SetIdleTimeout(timeout time.Duration)         // 设置空闲超时时间
	SetMaxRetries(retries int)                    // 设置重连最大尝试次数
	ResetIdleTime()                               // 重置空闲时间
	NotifyError(receiver chan error) <-chan error // 获取通知的管道
}

type ConnectionControllerInput struct {
	URI        string // AMQP URI
	AMQPConfig *amqp091.Config
	maxRetries int
	Ctx        context.Context
}

type ConnectionController struct {
	conn              *amqp091.Connection
	err               error
	noChan            bool
	errorReceivers    []chan error
	maxRetries        int
	initRetryWaitTime time.Duration
	maxRetryWaitTime  time.Duration
	retryWaitTime     time.Duration
	Ctx               context.Context
}

func NewConnectionController(in *ConnectionControllerInput) IConnectionController {
	var (
		conn *amqp091.Connection
		err  error
	)
	if in == nil {
		conn, err = nil, gerror.New("ConnectionControllerInput is nil")
	} else if in.AMQPConfig != nil {
		conn, err = amqp091.DialConfig(in.URI, *in.AMQPConfig)
	} else {
		conn, err = amqp091.Dial(in.URI)
	}
	connectionController := &ConnectionController{
		conn:       conn,
		err:        err,
		noChan:     false,
		maxRetries: in.maxRetries,
	}
	go connectionController.StartMonitor()
	return connectionController
}

func (c *ConnectionController) commitError(err error) {
	c.err = err
	if c.noChan || len(c.errorReceivers) == 0 {
		return
	}
	for _, receiver := range c.errorReceivers {
		receiver <- err
	}
}

// StartMonitor 启动连接的守护核心
func (c *ConnectionController) StartMonitor() {
	if c.conn == nil || c.err != nil {
		return
	}
	go c.keepAlive()                                          // 自动重连
	go c.monitorExceedIdleTimeoutAndMaxConnectionLiveliness() // 超时销毁，空闲销毁
}

func (c *ConnectionController) getRetryPolicy() time.Duration {
	if c.retryWaitTime >= c.maxRetryWaitTime {
		return c.maxRetryWaitTime
	}
	t, n := c.retryWaitTime, c.retryWaitTime*2 // double the wait time
	if n > c.maxRetryWaitTime {
		c.retryWaitTime = c.maxRetryWaitTime
	} else {
		c.retryWaitTime = n
	}
	return t
}

func (c *ConnectionController) resetRetryPolicy() {
	c.retryWaitTime = c.initRetryWaitTime
}

func (c *ConnectionController) keepAlive() {
	for err := range c.conn.NotifyClose(make(chan *amqp091.Error)) {
		g.Log().Errorf(c.Ctx, "[RabbitMQ] AMQP 连接丢失，错误信息: %s", err.Error())
		maxRetries := c.maxRetries
		for i := 1; i <= maxRetries; i++ { //nolint:staticcheck
			retryPolicy := c.getRetryPolicy()
			g.Log().Debugf(c.Ctx, "[RabbitMQ] AMQP 将在等待 %s 后，尝试第 %d 次重连……", retryPolicy.String(), i)
			time.Sleep(retryPolicy)
			e := c.Reconnect()
			if e == nil {
				g.Log().Debugf(c.Ctx, "[RabbitMQ] AMQP 重连成功")
				c.resetRetryPolicy()
				continue
			} else {
				g.Log().Errorf(c.Ctx, "[RabbitMQ] AMQP 重连失败，错误信息: %s", e.Error())
			}
			if i == c.maxRetries {
				g.Log().Errorf(c.Ctx, "[RabbitMQ] AMQP 重连失败，已达最大重连次数")
				c.commitError(e)
				break
			}
		}
	}
}

func (c *ConnectionController) monitorExceedIdleTimeoutAndMaxConnectionLiveliness() {}

func (c *ConnectionController) NotifyError(receiver chan error) <-chan error {
	// TODO implement me
	panic("implement me")
}

func (c *ConnectionController) Reconnect() error {
	// TODO implement me
	panic("implement me")
}

func (c *ConnectionController) IsClosed() bool {
	if c.conn == nil {
		return true
	}
	return c.conn.IsClosed()
}

func (c *ConnectionController) GetError() error {
	// TODO implement me
	panic("implement me")
}

func (c *ConnectionController) Close() error {
	// TODO implement me
	panic("implement me")
}

func (c *ConnectionController) EstablishedTime() *gtime.Time {
	// TODO implement me
	panic("implement me")
}

func (c *ConnectionController) IsIdle() bool {
	// TODO implement me
	panic("implement me")
}

func (c *ConnectionController) SetIdleTimeout(timeout time.Duration) {
	// TODO implement me
	panic("implement me")
}

func (c *ConnectionController) SetMaxRetries(retries int) {
	// TODO implement me
	panic("implement me")
}

func (c *ConnectionController) ResetIdleTime() {
	// TODO implement me
	panic("implement me")
}

func (c *ConnectionController) GetConnection() *amqp091.Connection {
	return c.conn
}

func (c *ConnectionController) IsError() bool {
	return c.err != nil
}
