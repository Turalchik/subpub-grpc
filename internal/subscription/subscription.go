package subscription

type Subscription struct {
	publisherID string
	msgHandler  func(msg interface{})
	messages    chan interface{}
	done        chan interface{}
}

type Config struct {
	messages    chan interface{}
	msgHandler  func(msg interface{})
	publisherID string
}

func NewSubscription(cfg *Config) *Subscription {
	subscription := &Subscription{
		publisherID: cfg.publisherID,
		msgHandler:  cfg.msgHandler,
		messages:    cfg.messages,
		done:        make(chan interface{}),
	}
	subscription.run()
	return subscription
}
