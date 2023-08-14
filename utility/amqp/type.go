package amqp

import (
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type PoolStrategy int

const (
	PoolStrategyLRU     PoolStrategy = iota // 最近最少使用
	PoolStrategyDefault              = PoolStrategyLRU
)

// ------------------------------------------------------------
//
//	Pool Definition
//
// ------------------------------------------------------------

type Pool interface {
	GetConnectionController() (ConnectionController, error)
	SetMaxOpenConnections(maxOpenConnections int) // 设置最大连接数
	SetMaxLifetime(maxLifetime time.Duration)     // 设置连接最大生命周期
	SetMaxIdleConnections(maxIdleConnections int) // 设置最大空闲连接数
	SetMinIdleConnections(minIdleConnections int) // 设置最小空闲连接数
	SetIdleTimeout(idleTimeout time.Duration)     // 设置空闲连接超时时间
	Len() int                                     // 获取连接池大小
	Close() error                                 // 关闭连接池
	IsExisted() bool                              // 判断连接池是否已经关闭
}

type ConnectionContainer interface {
	Pop() (ConnectionController, error) // 弹出一个连接
	Push(ConnectionController) error    // 推入一个连接
	Len() int                           // 获取连接池长度
	Remove(ConnectionController) error  // 移除一个连接
	Close() error                       // 销毁整个容器
}

// ------------------------------------------------------------
//
//	Connections Definition
//
// ------------------------------------------------------------

type ConnectionController interface {
	GetConnection() (*amqp091.Connection, error)
	ReleaseConnection()
	Close() error
	IsClosed() bool
}

type ConnectionConfig struct {
	Config *amqp091.Config
	URL    string
}
