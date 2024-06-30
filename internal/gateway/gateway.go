package gateway

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/lyouthzzz/realtimex/api/msg"
	"github.com/lyouthzzz/realtimex/api/protocol"
	"github.com/lyouthzzz/realtimex/internal/gateway/middleware"
	"github.com/lyouthzzz/realtimex/internal/gateway/operation"
	"github.com/lyouthzzz/realtimex/internal/gateway/session"
	"github.com/lyouthzzz/realtimex/internal/gateway/transport"
	"google.golang.org/protobuf/encoding/protojson"
)

type Gateway struct {
	msgc           msg.ReceiveServiceClient
	ms             []middleware.Middleware
	sessionManager session.Manager
	servers        []transport.Server
}

func NewServer() *Gateway {
	srv := &Gateway{}

	return srv
}

func (gateway *Gateway) Start(ctx context.Context) error {

	for _, srv := range gateway.servers {
		go gateway.listenSession(ctx, srv)
	}

	return nil
}

// 监听连接
func (gateway *Gateway) listenSession(_ context.Context, server transport.Server) {

	for sess := range server.Sessions() {
		// 添加连接
		gateway.sessionManager.Add(sess)

		go gateway.handleSession(sess)
	}

}

func (gateway *Gateway) handleSession(sess transport.Session) {
	defer func() {
		// 关闭连接
		_ = sess.Close(context.Background())
	}()

	var (
		proto *protocol.Proto
		err   error
	)

	// 第一条消息必须是连接消息
	if proto, err = sess.Read(); err != nil {
		return
	}

	if proto.Operation != protocol.Operation_Connect {
		// todo 鉴权逻辑
	}

	for {
		if proto, err = sess.Read(); err != nil {
			break
		}

		if err = middleware.Chain(gateway.ms...)(gateway.handleMsg)(context.Background(), proto); err != nil {
			// todo log
			continue
		}
	}
}

func (gateway *Gateway) MsgHandler(ctx context.Context, proto *protocol.Proto) error {
	sess := transport.SessionFromContext(ctx)

	ctx = metadata.AppendToClientContext(
		ctx,
		// 连接ID
		"X-Session-ID", sess.ID(),
		// 服务连接协议
		"X-Session-Protocol", sess.Kind().String(),
	)

	if err := middleware.Chain(gateway.ms...)(gateway.handleMsg)(ctx, proto); err != nil {
		return err
	}

	return nil
}

func (gateway *Gateway) handleMsg(ctx context.Context, proto *protocol.Proto) error {
	sess := transport.SessionFromContext(ctx)

	switch proto.Operation {
	case protocol.Operation_PingReq:
		if _, err := gateway.msgc.Ping(ctx, &protocol.PingReqPacket{}); err != nil {
			return err
		}
		_ = sess.Push(ctx, &protocol.Proto{Operation: operation.Pingresp})

	case protocol.Operation_Publish:
		var pubReq protocol.PublishPacket
		_ = protojson.Unmarshal([]byte(proto.Body), &pubReq)

		if _, err := gateway.msgc.Publish(ctx, &pubReq); err != nil {
			return err
		}

	case protocol.Operation_Subscribe:
		var subReq protocol.SubscribePacket
		_ = protojson.Unmarshal([]byte(proto.Body), &subReq)

		if _, err := gateway.msgc.Subscribe(ctx, &subReq); err != nil {
			return err
		}

		_ = sess.Push(ctx, &protocol.Proto{Operation: protocol.Operation_SubscribeAck})

	case protocol.Operation_Unsubscribe:

	case protocol.Operation_Disconnect:
		gateway.sessionManager.Del(sess)
	default:
		log.Infof("unknown operation: %d", proto.Operation)
	}

	return nil
}

func (gateway *Gateway) Stop(ctx context.Context) error {
	for _, srv := range gateway.servers {
		_ = srv.Stop(ctx)
	}
	return nil
}
