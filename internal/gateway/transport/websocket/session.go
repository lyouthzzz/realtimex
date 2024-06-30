package websocket

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/lyouthzzz/realtimex/api/protocol"
	"github.com/lyouthzzz/realtimex/internal/gateway/middleware"
	"github.com/lyouthzzz/realtimex/internal/gateway/operation"
	"github.com/lyouthzzz/realtimex/internal/gateway/transport"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

var _ transport.Session = (*Session)(nil)

type Session struct {
	Id         string
	server     *Server
	conn       *websocket.Conn
	msgHandler middleware.Handler
	ms         []middleware.Middleware
}

func (sess *Session) Read() (*protocol.Proto, error) {
	msgType, data, err := sess.conn.ReadMessage()
	if err != nil {
		// todo
		if websocket.IsUnexpectedCloseError(err,
			websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
		) {
			return nil, err
		}
	}

	p := &protocol.Proto{}
	switch msgType {
	case websocket.TextMessage:
		_ = protojson.Unmarshal(data, p)
	case websocket.PingMessage:
		p.Operation = operation.Pingreq
	case websocket.CloseMessage:
		p.Operation = operation.Disconnect
	default:
		// todo
	}

	return p, nil
}

func (sess *Session) Kind() transport.SessionKind {
	return transport.SessionKindWebsocket
}

func (sess *Session) ID() string { return sess.Id }

func (sess *Session) Push(ctx context.Context, proto *protocol.Proto) error {
	body, err := protojson.Marshal(proto)
	if err != nil {
		return err
	}
	// todo ping pong
	return sess.conn.WriteMessage(websocket.TextMessage, body)
}

func (sess *Session) Close(ctx context.Context) error {
	return sess.conn.Close()
}
