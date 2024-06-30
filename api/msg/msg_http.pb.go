// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.21.9
// source: api/msg/msg.proto

package msg

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	protocol "github.com/lyouthzzz/realtimex/api/protocol"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationReceiveServiceConnect = "/realtimex.msg.ReceiveService/Connect"
const OperationReceiveServiceDisconnect = "/realtimex.msg.ReceiveService/Disconnect"
const OperationReceiveServicePing = "/realtimex.msg.ReceiveService/Ping"
const OperationReceiveServicePublish = "/realtimex.msg.ReceiveService/Publish"
const OperationReceiveServiceSubscribe = "/realtimex.msg.ReceiveService/Subscribe"
const OperationReceiveServiceUnsubscribe = "/realtimex.msg.ReceiveService/Unsubscribe"

type ReceiveServiceHTTPServer interface {
	Connect(context.Context, *protocol.ConnectPacket) (*protocol.ConnectAckPacket, error)
	Disconnect(context.Context, *protocol.DisconnectPacket) (*protocol.DisconnectAckPacket, error)
	Ping(context.Context, *protocol.PingReqPacket) (*protocol.PingRespPacket, error)
	Publish(context.Context, *protocol.PublishPacket) (*protocol.PublishAckPacket, error)
	Subscribe(context.Context, *protocol.SubscribePacket) (*protocol.SubscribeAckPacket, error)
	Unsubscribe(context.Context, *protocol.UnsubscribePacket) (*protocol.UnsubscribeAckPacket, error)
}

func RegisterReceiveServiceHTTPServer(s *http.Server, srv ReceiveServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/realtimex/msg/receive/connect", _ReceiveService_Connect0_HTTP_Handler(srv))
	r.POST("/realtimex/msg/receive/subscribe", _ReceiveService_Subscribe0_HTTP_Handler(srv))
	r.POST("/realtimex/msg/receive/unsubscribe", _ReceiveService_Unsubscribe0_HTTP_Handler(srv))
	r.POST("/realtimex/msg/receive/publish", _ReceiveService_Publish0_HTTP_Handler(srv))
	r.POST("/realtimex/msg/receive/ping", _ReceiveService_Ping0_HTTP_Handler(srv))
	r.POST("/realtimex/msg/receive/disconnect", _ReceiveService_Disconnect0_HTTP_Handler(srv))
}

func _ReceiveService_Connect0_HTTP_Handler(srv ReceiveServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in protocol.ConnectPacket
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReceiveServiceConnect)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Connect(ctx, req.(*protocol.ConnectPacket))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*protocol.ConnectAckPacket)
		return ctx.Result(200, reply)
	}
}

func _ReceiveService_Subscribe0_HTTP_Handler(srv ReceiveServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in protocol.SubscribePacket
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReceiveServiceSubscribe)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Subscribe(ctx, req.(*protocol.SubscribePacket))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*protocol.SubscribeAckPacket)
		return ctx.Result(200, reply)
	}
}

func _ReceiveService_Unsubscribe0_HTTP_Handler(srv ReceiveServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in protocol.UnsubscribePacket
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReceiveServiceUnsubscribe)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Unsubscribe(ctx, req.(*protocol.UnsubscribePacket))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*protocol.UnsubscribeAckPacket)
		return ctx.Result(200, reply)
	}
}

func _ReceiveService_Publish0_HTTP_Handler(srv ReceiveServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in protocol.PublishPacket
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReceiveServicePublish)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Publish(ctx, req.(*protocol.PublishPacket))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*protocol.PublishAckPacket)
		return ctx.Result(200, reply)
	}
}

func _ReceiveService_Ping0_HTTP_Handler(srv ReceiveServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in protocol.PingReqPacket
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReceiveServicePing)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Ping(ctx, req.(*protocol.PingReqPacket))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*protocol.PingRespPacket)
		return ctx.Result(200, reply)
	}
}

func _ReceiveService_Disconnect0_HTTP_Handler(srv ReceiveServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in protocol.DisconnectPacket
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationReceiveServiceDisconnect)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Disconnect(ctx, req.(*protocol.DisconnectPacket))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*protocol.DisconnectAckPacket)
		return ctx.Result(200, reply)
	}
}

type ReceiveServiceHTTPClient interface {
	Connect(ctx context.Context, req *protocol.ConnectPacket, opts ...http.CallOption) (rsp *protocol.ConnectAckPacket, err error)
	Disconnect(ctx context.Context, req *protocol.DisconnectPacket, opts ...http.CallOption) (rsp *protocol.DisconnectAckPacket, err error)
	Ping(ctx context.Context, req *protocol.PingReqPacket, opts ...http.CallOption) (rsp *protocol.PingRespPacket, err error)
	Publish(ctx context.Context, req *protocol.PublishPacket, opts ...http.CallOption) (rsp *protocol.PublishAckPacket, err error)
	Subscribe(ctx context.Context, req *protocol.SubscribePacket, opts ...http.CallOption) (rsp *protocol.SubscribeAckPacket, err error)
	Unsubscribe(ctx context.Context, req *protocol.UnsubscribePacket, opts ...http.CallOption) (rsp *protocol.UnsubscribeAckPacket, err error)
}

type ReceiveServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewReceiveServiceHTTPClient(client *http.Client) ReceiveServiceHTTPClient {
	return &ReceiveServiceHTTPClientImpl{client}
}

func (c *ReceiveServiceHTTPClientImpl) Connect(ctx context.Context, in *protocol.ConnectPacket, opts ...http.CallOption) (*protocol.ConnectAckPacket, error) {
	var out protocol.ConnectAckPacket
	pattern := "/realtimex/msg/receive/connect"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationReceiveServiceConnect))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ReceiveServiceHTTPClientImpl) Disconnect(ctx context.Context, in *protocol.DisconnectPacket, opts ...http.CallOption) (*protocol.DisconnectAckPacket, error) {
	var out protocol.DisconnectAckPacket
	pattern := "/realtimex/msg/receive/disconnect"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationReceiveServiceDisconnect))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ReceiveServiceHTTPClientImpl) Ping(ctx context.Context, in *protocol.PingReqPacket, opts ...http.CallOption) (*protocol.PingRespPacket, error) {
	var out protocol.PingRespPacket
	pattern := "/realtimex/msg/receive/ping"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationReceiveServicePing))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ReceiveServiceHTTPClientImpl) Publish(ctx context.Context, in *protocol.PublishPacket, opts ...http.CallOption) (*protocol.PublishAckPacket, error) {
	var out protocol.PublishAckPacket
	pattern := "/realtimex/msg/receive/publish"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationReceiveServicePublish))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ReceiveServiceHTTPClientImpl) Subscribe(ctx context.Context, in *protocol.SubscribePacket, opts ...http.CallOption) (*protocol.SubscribeAckPacket, error) {
	var out protocol.SubscribeAckPacket
	pattern := "/realtimex/msg/receive/subscribe"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationReceiveServiceSubscribe))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ReceiveServiceHTTPClientImpl) Unsubscribe(ctx context.Context, in *protocol.UnsubscribePacket, opts ...http.CallOption) (*protocol.UnsubscribeAckPacket, error) {
	var out protocol.UnsubscribeAckPacket
	pattern := "/realtimex/msg/receive/unsubscribe"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationReceiveServiceUnsubscribe))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationPushServiceBroadcastTopicMsg = "/realtimex.msg.PushService/BroadcastTopicMsg"
const OperationPushServicePushUserMsg = "/realtimex.msg.PushService/PushUserMsg"

type PushServiceHTTPServer interface {
	BroadcastTopicMsg(context.Context, *protocol.PublishPacket) (*protocol.PublishAckPacket, error)
	PushUserMsg(context.Context, *PushUserMsgReq) (*PushUserMsgResp, error)
}

func RegisterPushServiceHTTPServer(s *http.Server, srv PushServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/realtimex/msg/push/broadcastTopicMsg", _PushService_BroadcastTopicMsg0_HTTP_Handler(srv))
	r.POST("/realtimex/msg/push/pushUserMsg", _PushService_PushUserMsg0_HTTP_Handler(srv))
}

func _PushService_BroadcastTopicMsg0_HTTP_Handler(srv PushServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in protocol.PublishPacket
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPushServiceBroadcastTopicMsg)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BroadcastTopicMsg(ctx, req.(*protocol.PublishPacket))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*protocol.PublishAckPacket)
		return ctx.Result(200, reply)
	}
}

func _PushService_PushUserMsg0_HTTP_Handler(srv PushServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PushUserMsgReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPushServicePushUserMsg)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PushUserMsg(ctx, req.(*PushUserMsgReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PushUserMsgResp)
		return ctx.Result(200, reply)
	}
}

type PushServiceHTTPClient interface {
	BroadcastTopicMsg(ctx context.Context, req *protocol.PublishPacket, opts ...http.CallOption) (rsp *protocol.PublishAckPacket, err error)
	PushUserMsg(ctx context.Context, req *PushUserMsgReq, opts ...http.CallOption) (rsp *PushUserMsgResp, err error)
}

type PushServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPushServiceHTTPClient(client *http.Client) PushServiceHTTPClient {
	return &PushServiceHTTPClientImpl{client}
}

func (c *PushServiceHTTPClientImpl) BroadcastTopicMsg(ctx context.Context, in *protocol.PublishPacket, opts ...http.CallOption) (*protocol.PublishAckPacket, error) {
	var out protocol.PublishAckPacket
	pattern := "/realtimex/msg/push/broadcastTopicMsg"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPushServiceBroadcastTopicMsg))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PushServiceHTTPClientImpl) PushUserMsg(ctx context.Context, in *PushUserMsgReq, opts ...http.CallOption) (*PushUserMsgResp, error) {
	var out PushUserMsgResp
	pattern := "/realtimex/msg/push/pushUserMsg"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPushServicePushUserMsg))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
