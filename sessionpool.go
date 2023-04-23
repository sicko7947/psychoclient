package psychoclient

import "github.com/gogf/gf/v2/container/gqueue"

// SessionPool interface allows operating a psycho client session pool
type SessionPool interface {
	Pop() Session
	Push(Session)
	Size() int
	Close()
}

type sessionPool struct {
	pool *gqueue.Queue
}

func NewPool(limit int) SessionPool {
	pool := gqueue.New(limit)
	return &sessionPool{pool: pool}
}

func (p *sessionPool) Size() int {
	return int(p.pool.Len())
}

func (p *sessionPool) Close() {
	p.pool.Close()
}
func (p *sessionPool) Pop() Session {
	if s := p.pool.Pop(); s != nil {
		return s.(Session)
	}
	return nil
}

func (p *sessionPool) Push(s Session) {
	defer recover()
	p.pool.Push(s)
}
