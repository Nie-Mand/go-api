package pubsub

import "go.uber.org/zap"

type PubSubClient struct {
	PubSubConfig
}

func NewPubSubClient(cfgs ...PubSubConfig) *PubSubClient {
	var cfg PubSubConfig
	if len(cfgs) > 0 {
		cfg = cfgs[0]
	}

	return &PubSubClient{
		PubSubConfig: cfg,
	}
}

func (c *PubSubClient) Pub(message string) error {
	zap.L().Info("Publishing message", zap.String("message", message))
	return nil
}
