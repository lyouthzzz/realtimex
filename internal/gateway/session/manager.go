package session

import (
	"github.com/lyouthzzz/realtimex/internal/gateway/transport"
	"sync"
)

// Manager session管理器
type Manager struct {
	// sessions 读写锁
	mutex sync.RWMutex
	// 连接map 通过session id获取连接
	sessionIdMap map[string]transport.Session
}

func (m *Manager) Add(session transport.Session) {
	m.mutex.Lock()
	m.sessionIdMap[session.ID()] = session
	m.mutex.Unlock()
}

func (m *Manager) Del(session transport.Session) {
	m.mutex.Lock()
	delete(m.sessionIdMap, session.ID())
	m.mutex.Unlock()
}
