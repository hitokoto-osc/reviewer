package amqp

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/rabbitmq/amqp091-go"
)

type connectionController struct {
	id       string // 连接唯一标识
	conn     *amqp091.Connection
	pool     Pool
	lastUsed *gtime.Time
	isUsed   bool
	mu       sync.Mutex
}

func NewConnectionController(conn *amqp091.Connection, pool Pool) ConnectionController {
	controller := &connectionController{
		id:       gtime.TimestampNanoStr(),
		conn:     conn,
		pool:     pool,
		lastUsed: gtime.Now(),
		isUsed:   false,
		mu:       sync.Mutex{},
	}
	go controller.registerCloseHandler()
	return controller
}

func (c *connectionController) GetConnection() (*amqp091.Connection, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.IsClosed() {
		return nil, gerror.New("connection is closed")
	}
	c.lastUsed = gtime.Now()
	c.isUsed = true
	return c.conn, nil
}

func (c *connectionController) ReleaseConnection() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.IsClosed() {
		return
	}
	c.isUsed = false
	c.lastUsed = gtime.Now()
	_ = c.pool.(*pool).releaseConnectionController(c)
}

func (c *connectionController) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.IsClosed() {
		return nil
	}
	err := c.conn.Close()
	c.conn = nil
	return err
}

func (c *connectionController) IsClosed() bool {
	return c.conn == nil || c.conn.IsClosed()
}

func (c *connectionController) registerCloseHandler() {
	ctx := context.Background()
	for err := range c.conn.NotifyClose(make(chan *amqp091.Error)) {
		if err != nil {
			g.Log().Debugf(ctx, "amqp connection[%s] closed: %s", c.id, err.Error())
			c.conn = nil
			_ = c.pool.(*pool).removeConnectionController(c)
			return
		}
	}
}

func dial(config *ConnectionConfig) (*amqp091.Connection, error) {
	if config.Config != nil {
		return amqp091.DialConfig(config.URL, *config.Config)
	} else {
		return amqp091.Dial(config.URL)
	}
}
