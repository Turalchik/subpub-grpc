package subpub

import "github.com/Turalchik/subpub-grpc/internal/subscription"

func (subPub *subpub) Subscribe(subject string, msgHandler MessageHandler) (Subscription, error) {
	subPub.mu.Lock()
	defer subPub.mu.Unlock()

	sub := newSubscriber(msgHandler)
	subPub.publisher2Subscribers[subject] = append(subPub.publisher2Subscribers[subject], sub)

	unsubscribeFunc := func() {
		subPub.unsubscribe(subject, sub)
	}

	return &subscription.Subscription{
		UnsubscribeFunc: unsubscribeFunc,
	}, nil
}
