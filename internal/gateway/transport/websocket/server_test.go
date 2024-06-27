package websocket

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lyouthzzz/realtimex/api/gateway"
	"github.com/lyouthzzz/realtimex/api/protocol"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"testing"
)

type authClient struct {
	gateway.UnimplementedAuthnServiceServer
}

func (auth *authClient) DoAuthn(ctx context.Context, in *gateway.DoAuthNRequest, opts ...grpc.CallOption) (*gateway.User, error) {
	return &gateway.User{Uid: "1"}, nil
}

func TestWebsocketServer(t *testing.T) {
	srv := NewServer(AuthnClient(&authClient{}))

	err := srv.Start(context.Background())
	require.NoError(t, err)
}

func TestWebsocketClient(t *testing.T) {
	uri := "ws://localhost:8000/gateway/ws/v1"
	conn, resp, err := websocket.DefaultDialer.Dial(uri, nil)
	require.NoError(t, err)
	defer conn.Close()

	fmt.Println(resp)

	t.Run("ping", func(t *testing.T) {
		err = conn.WriteMessage(websocket.PingMessage, []byte("ping"))
		require.NoError(t, err)
	})

}

func TestProto(t *testing.T) {

	proto := &protocol.Proto{
		Operation: 1,
	}

	body, err := protojson.Marshal(proto)
	require.NoError(t, err)

	fmt.Println(string(body))
}
