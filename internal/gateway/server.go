package gateway

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/lyouthzzz/realtimex/api/msg"
	"github.com/lyouthzzz/realtimex/api/protocol"
	"github.com/lyouthzzz/realtimex/internal/gateway/cmd"
	"github.com/lyouthzzz/realtimex/internal/gateway/middleware"
	"github.com/lyouthzzz/realtimex/internal/gateway/transport"
	"github.com/lyouthzzz/realtimex/internal/gateway/transport/websocket"
	"google.golang.org/protobuf/encoding/protojson"
)

type Server struct {
	srv  transport.Server
	msgc msg.MsgServiceClient
	ms   []middleware.Middleware
}

func NewServer() *Server {
	srv := &Server{}

	srv.srv = websocket.NewServer()
	return srv
}

func (srv *Server) Start(ctx context.Context) error {
	for conn := range srv.srv.Sessions() {
		srv.dispatchConn(conn)
	}

	if err := srv.srv.Start(ctx); err != nil {
		return err
	}
	return nil
}

func (srv *Server) MsgHandler(ctx context.Context, proto *protocol.Proto) error {
	sess := transport.SessionFromContext(ctx)

	ctx = metadata.AppendToClientContext(ctx,
		// 连接ID
		"X-Session-ID", sess.ID(),
		// 服务连接协议
		"X-Session-Protocol", sess.Kind().String(),
	)

	// todo 判断session是否合法

	if err := middleware.Chain(srv.ms...)(srv.handleMsg)(ctx, proto); err != nil {
		return err
	}

	return nil
}

func (srv *Server) handleMsg(ctx context.Context, proto *protocol.Proto) error {
	sess := transport.SessionFromContext(ctx)

	switch proto.Cmd {
	case protocol.Cmd_Connect:
		if _, err := srv.msgc.Connect(ctx, &msg.ConnectReq{}); err != nil {
			_ = sess.Push(ctx, &protocol.Proto{Cmd: cmd.Disconnect, Body: err.Error()})
			return err
		}

	case protocol.Cmd_PingReq:
		if _, err := srv.msgc.Ping(ctx, &msg.PingReq{}); err != nil {
			return err
		}
		_ = sess.Push(ctx, &protocol.Proto{Cmd: cmd.Pingresp})

	case protocol.Cmd_Publish:
		var pubReq msg.PublishReq
		_ = protojson.Unmarshal([]byte(proto.Body), &pubReq)

		if _, err := srv.msgc.Publish(ctx, &pubReq); err != nil {
			return err
		}

	case protocol.Cmd_Subscribe:
		var subReq msg.SubscribeReq
		_ = protojson.Unmarshal([]byte(proto.Body), &subReq)

		if _, err := srv.msgc.Subscribe(ctx, &subReq); err != nil {
			return err
		}

		sess.Push(ctx, &protocol.Proto{Cmd: protocol.Cmd_Subscribe})

	case protocol.Cmd_Unsubscribe:

	default:
		log.Infof("unknown cmd: %d", proto.Cmd)
	}

	return nil
}

func (srv *Server) Stop(ctx context.Context) error {
	return srv.srv.Stop(ctx)
}
