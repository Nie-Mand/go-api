package pubsub

type PubSubConfig struct {
}

func NewPubSubConfig() PubSubConfig {
	return PubSubConfig{}
}

func NewDefaultPubSubConfig() PubSubConfig {
	return NewPubSubConfig()
}
