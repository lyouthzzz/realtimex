package websocket

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"github.com/lyouthzzz/realtimex/api/protocol"
	"github.com/lyouthzzz/realtimex/internal/gateway/cmd"
	"github.com/lyouthzzz/realtimex/internal/gateway/middleware"
	"github.com/lyouthzzz/realtimex/internal/gateway/transport"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

type sessionKey struct{}

var _ transport.Session = (*Session)(nil)

type Session struct {
	Id         string
	server     *Server
	conn       *websocket.Conn
	msgHandler middleware.Handler
	ms         []middleware.Middleware
}

func (sess *Session) Kind() transport.SessionKind {
	return transport.SessionKindWebsocket
}

func (sess *Session) ID() string {
	return sess.Id
}

func (sess *Session) Start(ctx context.Context) error {
	go sess.read()
	return nil
}

func (sess *Session) read() {
	for {
		var (
			msgType int
			data    []byte
			err     error
		)

		if msgType, data, err = sess.conn.ReadMessage(); err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure,
			) {
				log.Context(context.Background()).Errorf("session read message err: %s", err)
				break
			}
		}

		p := &protocol.Proto{}
		switch msgType {
		case websocket.TextMessage:
			if err = protojson.Unmarshal(data, p); err != nil {
				continue
			}
		case websocket.PingMessage:
			p.Cmd = cmd.Pingreq
		case websocket.CloseMessage:
			p.Cmd = cmd.Disconnect
		default:
			continue
		}

		ctx := transport.ContextWithSession(context.Background(), sess)
		if err = middleware.Chain(sess.ms...)(sess.msgHandler)(ctx, p); err != nil {
			continue
		}
	}
}

func (sess *Session) Push(ctx context.Context, proto *protocol.Proto) error {
	body, err := protojson.Marshal(proto)
	if err != nil {
		return err
	}
	return sess.conn.WriteMessage(websocket.TextMessage, body)
}

func (sess *Session) Stop(ctx context.Context) error {
	return sess.conn.Close()
}
