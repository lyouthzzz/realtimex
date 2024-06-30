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
	// Close 关闭连接
	Close(ctx context.Context) error
	// Read 读取消息
	Read() (*protocol.Proto, error)
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

type sessionKey struct{}

func ContextWithSession(ctx context.Context, conn Session) context.Context {
	return context.WithValue(ctx, sessionKey{}, conn)
}

func SessionFromContext(ctx context.Context) Session {
	v := ctx.Value(sessionKey{})
	if v == nil {
		return nil
	}
	return v.(Session)
}
