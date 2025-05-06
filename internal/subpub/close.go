package subpub

import "context"

func (subPub *subpub) Close(ctx context.Context) error {
	subPub.mu.Lock()
	if subPub.closed {
		subPub.mu.Unlock()
		return nil
	}
	subPub.closed = true

	allSubs := make([]*subscriber, 0)
	for pub, subs := range subPub.pub2subs {
		allSubs = append(allSubs, subs...)
		delete(subPub.pub2subs, pub)
	}
	subPub.mu.Unlock()

	for _, sub := range allSubs {
		sub.Close()
	}

	done := make(chan interface{})
	go func() {
		subPub.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
