package subpub

import "sync"

type subscriber struct {
	mu         sync.Mutex
	msgHandler MessageHandler
	queueMsg   []interface{}
}

func newSubscriber(msgHandler MessageHandler) *subscriber {
	return &subscriber{
		msgHandler: msgHandler,
	}
}
