package subscription

func (subscription *Subscription) run() {
	go func() {
		for {
			select {
			case msg := <-subscription.messages:
				subscription.msgHandler(msg)
			case <-subscription.done:
				return
			}
		}
	}()
}
