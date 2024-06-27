package websocket

import (
	"github.com/gorilla/websocket"
	"time"
)

const (
	defaultAddress          = ":8000"
	defaultPath             = "/gateway/ws/v1"
	defaultSessConnectedCap = 2048
)

var defaultWebsocketUpgrader = &websocket.Upgrader{
	HandshakeTimeout:  time.Second,
	ReadBufferSize:    10 * 1024 * 1024, // 10k
	WriteBufferSize:   64 * 1024 * 1024, // 64k
	WriteBufferPool:   nil,
	Subprotocols:      nil,
	Error:             nil,
	CheckOrigin:       nil,
	EnableCompression: true,
}

type Option func(*Server)

func Address(addr string) Option {
	return func(srv *Server) { srv.address = addr }
}

func Path(path string) Option {
	return func(srv *Server) { srv.path = path }
}

func Upgrader(upgrader *websocket.Upgrader) Option {
	return func(srv *Server) { srv.upgrader = upgrader }
}
