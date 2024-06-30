package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lyouthzzz/realtimex/api/protocol"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

type Handler func(ctx context.Context, proto *protocol.Proto) error

type Middleware func(Handler) Handler

func Chain(ms ...Middleware) Middleware {
	return func(h Handler) Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			h = ms[i](h)
		}
		return h
	}
}

func Logger() Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, p *protocol.Proto) error {
			var (
				startTime = time.Now()
			)

			err := next(ctx, p)

			body, _ := protojson.Marshal(p)
			kvs := make([]any, 0)
			kvs = append(kvs, "msg", string(body))
			kvs = append(kvs, "operation", p.Cmd)
			kvs = append(kvs, "ingress", len(body))
			kvs = append(kvs, "latency", time.Since(startTime).Seconds())
			kvs = append(kvs, "msg", string(body))

			if err != nil {
				kvs = append(kvs, "error", err.Error())
				log.Context(ctx).Errorw(kvs...)
			} else {
				log.Context(ctx).Infow(kvs...)
			}

			return err
		}
	}
}
