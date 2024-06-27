package gateway

import (
	"bufio"
	"container/list"
	"sync"
)

type Channel struct {
	Writer bufio.Writer
	Reader bufio.Reader

	Next *Channel
	Prev *Channel

	Uid   string
	Subs  *list.List // 订阅主题通道链表
	mutex sync.RWMutex
}

func NewChannel() *Channel {
	c := new(Channel)
	c.Subs = list.New()
	return c
}

func (c *Channel) Subscribe(topic string) {
	c.mutex.Lock()
	c.Subs.PushFront(topic)
	c.mutex.Unlock()
}

func (c *Channel) Unsubscribe(topic string) {
	c.mutex.Lock()
	c.Subs.Remove(&list.Element{Value: topic})
	c.mutex.Unlock()
}
