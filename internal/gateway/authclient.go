package gateway

import (
	"context"
	"github.com/lyouthzzz/realtimex/api/gateway"
	"google.golang.org/grpc"
)

type authClient struct {
	gateway.UnimplementedAuthnServiceServer
}

func (auth *authClient) DoAuthn(ctx context.Context, in *gateway.DoAuthNRequest, opts ...grpc.CallOption) (*gateway.User, error) {
	return &gateway.User{Uid: "1"}, nil
}
