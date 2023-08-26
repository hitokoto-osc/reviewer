package amqp

import (
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/os/gtime"
)

type PoolConfig struct {
	MaxOpenConnections int           // 最大连接数
	MaxLifetime        time.Duration // 连接最大生命周期
	IdleTimeout        time.Duration // 空闲连接超时时间
	MaxIdleConnections int           // 最大空闲连接数
	MinIdleConnections int           // 最小空闲连接数
	Strategy           PoolStrategy  // 连接池返回策略
	CheckInterval      time.Duration // 健康检查间隔
	WaitInterval       time.Duration // 等待间隔
	WaitingTimeout     time.Duration // 最大等待时间
}

type pool struct {
	config           *PoolConfig
	connectionConfig *ConnectionConfig
	container        ConnectionContainer
	size             int // 当前连接池大小
	exited           bool
	mu               sync.Mutex
}

func DefaultPoolConfig() *PoolConfig {
	return &PoolConfig{
		MaxOpenConnections: 100,
		MaxLifetime:        time.Minute * 5,
		IdleTimeout:        time.Second * 30,
		MaxIdleConnections: 10,
		MinIdleConnections: 5,
		Strategy:           PoolStrategyDefault,
		CheckInterval:      time.Second * 5, // 5s 检查一次
		WaitInterval:       20 * time.Millisecond,
		WaitingTimeout:     200 * time.Millisecond,
	}
}

func NewPool(config *PoolConfig, connectionConfig *ConnectionConfig) (Pool, error) {
	if config == nil || connectionConfig == nil {
		return nil, gerror.New("Config or connection Config is nil")
	}
	p := &pool{
		config:           config,
		connectionConfig: connectionConfig,
		container:        NewPoolContainerQueueWithCap(config.MaxOpenConnections),
		size:             0,
		exited:           false,
		mu:               sync.Mutex{},
	}
	var err error
	for i := 0; i < config.MinIdleConnections; i++ {
		var controller ConnectionController
		controller, err = p.newConnectionController()
		if err != nil {
			break
		}
		_ = p.container.Push(controller)
	}
	if err != nil {
		return nil, err
	}
	go p.healthCheck()
	return p, nil
}

func (p *pool) GetConnectionController() (ConnectionController, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.container.Len() == 0 {
		if p.size < p.config.MaxOpenConnections {
			return p.newConnectionController()
		}
		waitStart := gtime.Now()
		for p.container.Len() == 0 {
			if p.config.WaitingTimeout > 0 && waitStart.Add(p.config.WaitingTimeout).Before(gtime.Now()) {
				return nil, gerror.New("failed to acquire connection from pool: waiting timeout")
			}
			time.Sleep(p.config.WaitInterval)
		}
	}
	conn, err := p.container.Pop()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (p *pool) SetMaxOpenConnections(maxOpenConnections int) {
	p.config.MaxOpenConnections = maxOpenConnections
}

func (p *pool) SetMaxLifetime(maxLifetime time.Duration) {
	p.config.MaxLifetime = maxLifetime
}

func (p *pool) SetMaxIdleConnections(maxIdleConnections int) {
	p.config.MaxIdleConnections = maxIdleConnections
}

func (p *pool) SetMinIdleConnections(minIdleConnections int) {
	p.config.MinIdleConnections = minIdleConnections
}

func (p *pool) SetIdleTimeout(idleTimeout time.Duration) {
	p.config.IdleTimeout = idleTimeout
}

func (p *pool) Close() error {
	err := p.container.Close()
	p.exited = true
	return err
}

func (p *pool) IsExisted() bool {
	return p.exited
}

// shouldDoRemove 判断是否需要移除连接
func (p *pool) shouldDoRemove(cc ConnectionController) bool {
	if cc.IsClosed() {
		return true
	}
	if p.config.MaxLifetime > 0 && cc.(*connectionController).lastUsed.Add(p.config.MaxLifetime).Before(gtime.Now()) {
		return true
	}
	if p.config.MaxIdleConnections > 0 && p.container.Len() > p.config.MaxIdleConnections {
		return true
	}
	if p.config.IdleTimeout > 0 && cc.(*connectionController).lastUsed.Add(p.config.IdleTimeout).Before(gtime.Now()) {
		return true
	}
	return false
}

func (p *pool) releaseConnectionController(controller ConnectionController) error {
	cc := controller.(*connectionController)
	if !cc.IsClosed() {
		if p.shouldDoRemove(cc) {
			err := cc.Close()
			if err != nil {
				return err
			}
			p.size--
			return nil
		}
		return p.container.Push(cc)
	}
	return nil
}

func (p *pool) Len() int {
	return p.size
}

func (p *pool) newConnectionController() (ConnectionController, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.size >= p.config.MaxOpenConnections {
		return nil, gerror.New("connection pool is full")
	}
	conn, err := dial(p.connectionConfig)
	if err != nil {
		return nil, err
	}
	p.size++
	return NewConnectionController(conn, p), nil
}

// healthCheck 健康检查
func (p *pool) healthCheck() {
	for {
		p.mu.Lock()
		if p.IsExisted() {
			return
		}
		if p.size < p.config.MinIdleConnections { // 维护最低连接数
			for i := 0; i < p.config.MinIdleConnections-p.size; i++ {
				conn, err := p.newConnectionController()
				if err != nil {
					break
				}
				_ = p.container.Push(conn)
			}
		}
		if p.config.MaxIdleConnections > 0 && p.container.Len() > p.config.MaxIdleConnections { // 维护最大空闲连接数
			for i := 0; i < p.container.Len()-p.config.MaxIdleConnections; i++ {
				conn, err := p.container.Pop()
				if err != nil {
					break
				}
				_ = conn.Close()
				p.size--
			}
		}
		p.mu.Unlock()
		time.Sleep(p.config.CheckInterval)
	}
}

func (p *pool) removeConnectionController(cc ConnectionController) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	err := p.container.Remove(cc)
	if err != nil {
		return err
	}
	p.size--
	return nil
}
