package subpub

func (subPub *subpub) Unsubscribe(subject string, sub *subscriber) {
	subPub.mu.Lock()
	defer subPub.mu.Unlock()

	for i, ptrSub := range subPub.pub2subs[subject] {
		if ptrSub == sub {
			subPub.pub2subs[subject] = append(subPub.pub2subs[subject][:i], subPub.pub2subs[subject][i+1:]...)
			break
		}
	}
}
