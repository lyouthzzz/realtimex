package websocket

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/lyouthzzz/realtimex/internal/gateway/middleware"
	"github.com/lyouthzzz/realtimex/internal/gateway/transport"
	"github.com/pkg/errors"
	"net/http"
)

var _ transport.Server = (*Server)(nil)

type Server struct {
	*http.Server
	// websocket服务地址
	address string
	// ws开放路径
	path string
	// http路由
	router *mux.Router
	// websocket upgrader
	upgrader *websocket.Upgrader
	// sessions 连接池
	sessions chan transport.Session
	// 消息处理器
	msgHandler middleware.Handler
	// 消息处理中间件
	msgMiddlewares []middleware.Middleware
}

func NewServer(opts ...Option) *Server {
	srv := &Server{
		address:  defaultAddress,
		path:     defaultPath,
		router:   mux.NewRouter(),
		upgrader: defaultWebsocketUpgrader,
		sessions: make(chan transport.Session, defaultSessConnectedCap),
	}

	for _, opt := range opts {
		opt(srv)
	}

	srv.Server = &http.Server{
		Addr:    srv.address,
		Handler: srv.router,
	}

	return srv
}

func (srv *Server) Start(_ context.Context) error {
	srv.router.Handle(srv.path, srv.websocketHandler())

	if err := srv.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (srv *Server) websocketHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var (
			conn *websocket.Conn
			err  error
		)

		if conn, err = srv.upgrader.Upgrade(w, req, nil); err != nil {
			return
		}

		session := &Session{
			Id:         uuid.New().String(),
			server:     srv,
			conn:       conn,
			msgHandler: srv.msgHandler,
			ms:         srv.msgMiddlewares,
		}

		srv.sessions <- session
	})
}

func (srv *Server) Sessions() chan transport.Session {
	return srv.sessions
}

func (srv *Server) Stop(ctx context.Context) error {
	_ = srv.Server.Shutdown(ctx)
	close(srv.sessions)
	return nil
}
