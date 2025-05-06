package subpub

import "errors"

func (subPub *subpub) Publish(subject string, msg interface{}) error {
	subPub.mu.Lock()
	defer subPub.mu.Unlock()

	if subPub.closed {
		return errors.New("subpub is closed")
	}

	for _, sub := range subPub.pub2subs[subject] {
		sub.AppendMsg(msg)
	}

	return nil
}
