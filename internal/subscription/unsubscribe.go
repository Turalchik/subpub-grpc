package subscription

func (subscription *Subscription) Unsubscribe() {
	close(subscription.done)
}
