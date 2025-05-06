package subpub

import (
	"context"
	"sync"
)

type MessageHandler func(msg interface{})

type Subscription interface {
	Unsubscribe()
}

type SubPub interface {
	Subscribe(subject string, cb MessageHandler) (Subscription, error)
	Publish(subject string, msg interface{}) error
	Close(ctx context.Context) error
}

type subpub struct {
	wg       *sync.WaitGroup
	mu       *sync.Mutex
	closed   bool
	pub2subs map[string][]*subscriber
}

func NewSubPub() SubPub {
	return &subpub{
		wg:       &sync.WaitGroup{},
		mu:       &sync.Mutex{},
		pub2subs: make(map[string][]*subscriber),
	}
}
