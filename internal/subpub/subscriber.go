package subpub

import "sync"

type subscriber struct {
	msgHandler MessageHandler
	mu         *sync.Mutex
	cond       *sync.Cond
	closed     bool
	done       chan interface{}
	queueMsg   []interface{}
}

func newSubscriber(msgHandler MessageHandler, wg *sync.WaitGroup) *subscriber {
	sub := &subscriber{
		msgHandler: msgHandler,
		mu:         &sync.Mutex{},
	}
	sub.cond = sync.NewCond(sub.mu)
	wg.Add(1)
	go sub.Process(wg)
	return sub
}

func (sub *subscriber) AppendMsg(msg interface{}) {
	sub.mu.Lock()
	defer sub.mu.Unlock()

	if sub.closed {
		return
	}

	sub.queueMsg = append(sub.queueMsg, msg)
	sub.cond.Signal()
}

func (sub *subscriber) Process(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		sub.mu.Lock()

		if len(sub.queueMsg) == 0 && !sub.closed {
			sub.cond.Wait()
		}

		if sub.closed {
			sub.mu.Unlock()
			return
		}

		msg := sub.queueMsg[0]
		sub.queueMsg = sub.queueMsg[1:]
		sub.mu.Unlock()

		sub.msgHandler(msg)
	}
}

func (sub *subscriber) HandleMsg(msg interface{}) <-chan interface{} {
	sub.msgHandler(msg)
	ch := make(chan interface{})
	close(ch)
	return ch
}

func (sub *subscriber) Close() {
	sub.mu.Lock()
	defer sub.mu.Unlock()

	if sub.closed {
		return
	}

	sub.closed = true
	sub.cond.Signal()
}
