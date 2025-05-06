package subpub

func (subPub *subpub) unsubscribe(subject string, sub *subscriber) {
	subPub.mu.Lock()
	defer subPub.mu.Unlock()

	subs := subPub.publisher2Subscribers[subject]
	for i, v := range subs {
		if v == sub {
			subPub.publisher2Subscribers[subject] = append(subs[:i], subs[i+1:]...)
			break
		}
	}
}
