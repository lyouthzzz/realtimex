package transport

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/lyouthzzz/realtimex/api/protocol"
)

type SessionKind string

var (
	SessionKindWebsocket SessionKind = "websocket"
)

func (kind SessionKind) String() string {
	return string(kind)
}

type Session interface {
	// ID 唯一ID
	ID() string
	// Start 开始连接工作
	Start(ctx context.Context) error
	// Stop 停止连接工作
	Stop(ctx context.Context) error
	// Push 推送消息
	Push(ctx context.Context, proto *protocol.Proto) error
	// Kind 连接类型 websocket/mqtt
	Kind() SessionKind
}

type Server interface {
	transport.Server
	// Sessions 连接池
	Sessions() chan Session
}

type connKey struct{}

func ContextWithSession(ctx context.Context, conn Session) context.Context {
	return context.WithValue(ctx, connKey{}, conn)
}

func SessionFromContext(ctx context.Context) Session {
	v := ctx.Value(connKey{})
	if v == nil {
		return nil
	}
	return v.(Session)
}
