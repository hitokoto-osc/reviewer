package amqp

import (
	"sync"

	"github.com/gogf/gf/v2/errors/gerror"
)

// PoolContainerQueue 使用 channel 实现队列
type PoolContainerQueue struct {
	connections chan ConnectionController
	mu          sync.Mutex
}

func NewPoolContainerQueueWithCap(capacity int) *PoolContainerQueue {
	return &PoolContainerQueue{
		connections: make(chan ConnectionController, capacity),
	}
}

func NewPoolContainerQueue() *PoolContainerQueue {
	return &PoolContainerQueue{
		connections: make(chan ConnectionController),
	}
}

func (p *PoolContainerQueue) Pop() (ConnectionController, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.Len() == 0 {
		return nil, gerror.New("connection pool is empty")
	}
	return <-p.connections, nil
}

func (p *PoolContainerQueue) Push(c ConnectionController) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.connections <- c
	return nil
}

func (p *PoolContainerQueue) Len() int {
	return len(p.connections)
}

func (p *PoolContainerQueue) Remove(c ConnectionController) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	lens := p.Len()
	for i := 0; i < lens; i++ {
		conn, _ := p.Pop()
		if conn.(*connectionController).id == c.(*connectionController).id {
			return nil // 不需要再放回队列
		}
		_ = p.Push(conn) // 重新放回队列
	}
	return gerror.New("connection not found")
}

func (p *PoolContainerQueue) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	lens := p.Len()
	for i := 0; i < lens; i++ {
		conn, _ := p.Pop()
		_ = conn.Close()
	}
	close(p.connections)
	return nil
}

func (p *PoolContainerQueue) ReCap(newCap int) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	lens := p.Len()
	// 重新分配容量
	newChan := make(chan ConnectionController, newCap)
	for i := 0; i < lens; i++ {
		conn, _ := p.Pop()
		newChan <- conn
	}
	close(p.connections)
	p.connections = newChan
	return nil
}
