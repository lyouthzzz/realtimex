package middleware

import (
	"context"
	"github.com/lyouthzzz/realtimex/api/protocol"
	"google.golang.org/protobuf/proto"
)

type Handler func(ctx context.Context, proto *protocol.Proto) (proto.Message, error)

type Middleware func(Handler) Handler

func Chain(ms ...Middleware) Middleware {
	return func(next Handler) Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			next = ms[i](next)
		}
		return next
	}
}
