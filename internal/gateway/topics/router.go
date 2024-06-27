package topics

import (
	"container/list"
	"github.com/lyouthzzz/realtimex/internal/gateway/middleware"
	"sync"
)

type Node struct {
	// 主题名称
	Topic string
	// 是否是通配符
	wildChild bool
	Handler   middleware.Middleware
	Children  list.List
}

type Router struct {
	topics sync.Map
	regex  bool
}
