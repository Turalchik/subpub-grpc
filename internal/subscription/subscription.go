package subscription

type Subscription struct {
	UnsubscribeFunc func()
}

func (subscription *Subscription) Unsubscribe() {
	subscription.UnsubscribeFunc()
}
