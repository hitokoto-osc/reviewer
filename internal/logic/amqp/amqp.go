package amqp

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/hitokoto-osc/reviewer/internal/service"
	"github.com/hitokoto-osc/reviewer/utility/amqp"
)

type sAMQP struct {
	pool amqp.Pool
}

func init() {
	ctx := gctx.New()
	data := g.Cfg().MustData(ctx)
	amqpConfig, ok := data["amqp"].(g.Map)
	if !ok {
		panic("AMQP 配置不存在！")
	}
	url, ok := amqpConfig["url"].(string)
	if !ok {
		panic("AMQP 配置不存在！")
	}
	connectionConfig := &amqp.ConnectionConfig{
		URL: url,
	}
	amqpService := New(connectionConfig)
	service.RegisterAMQP(amqpService)
}

func New(connectionConfig *amqp.ConnectionConfig) service.IAMQP {
	defaultConfig := amqp.DefaultPoolConfig()
	pool, err := amqp.NewPool(defaultConfig, connectionConfig)
	if err != nil {
		panic(err)
	}
	return &sAMQP{
		pool: pool,
	}
}

func (s *sAMQP) GetConnectionController() (amqp.ConnectionController, error) {
	return s.pool.GetConnectionController()
}
