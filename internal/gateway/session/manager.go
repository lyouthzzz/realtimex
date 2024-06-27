package session

import (
	"container/list"
	"github.com/lyouthzzz/realtimex/internal/gateway/transport"
	"sync"
)

// Manager session管理器
type Manager struct {
	// sessions 读写锁
	mutex sync.RWMutex
	// 连接列表 transport.Session 对象
	sessions list.List
	// 连接map 通过session id获取连接
	sessionIdMap map[string]transport.Session
}

func (m *Manager) Add(session transport.Session) {
	m.mutex.Lock()
	// 放在队头
	m.sessions.PushFront(session)
	m.mutex.Unlock()
}

func (m *Manager) Del(session transport.Session) {
	// todo 删除
}
