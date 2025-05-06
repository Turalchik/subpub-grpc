package subpub

import (
	"errors"
	"github.com/Turalchik/subpub-grpc/internal/subscription"
)

func (subPub *subpub) Subscribe(subject string, cb MessageHandler) (Subscription, error) {
	subPub.mu.Lock()
	defer subPub.mu.Unlock()

	if subPub.closed {
		return nil, errors.New("subpub is closed")
	}

	sub := newSubscriber(cb, subPub.wg)
	subPub.pub2subs[subject] = append(subPub.pub2subs[subject], sub)

	unsubscribeFunc := func() {
		subPub.Unsubscribe(subject, sub)
	}

	return &subscription.Subscription{UnsubscribeFunc: unsubscribeFunc}, nil
}
