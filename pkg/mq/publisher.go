package mq

import "context"

const SiteCheckQueue = "site_check_queue"

type Publisher interface {
	Publish(ctx context.Context, queueName string, body []byte) error
	Close()
}
