package worker

import (
	"api-service/internal/models"
	"api-service/pkg"
	"api-service/pkg/mq"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type Consumer struct {
	conn       *amqp.Connection
	channel    *amqp.Channel
	checkerSvc pkg.CheckerService
	wg         sync.WaitGroup
	done       chan struct{}
}

func NewConsumer(url string, checkerSvc pkg.CheckerService) (*Consumer, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &Consumer{
		conn:       conn,
		channel:    ch,
		checkerSvc: checkerSvc,
		done:       make(chan struct{}),
	}, nil
}

func (c *Consumer) Start(concurrency int) {
	_, err := c.channel.QueueDeclare(
		mq.SiteCheckQueue,
		true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	err = c.channel.Qos(concurrency, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS: %v", err)
	}

	msgs, err := c.channel.Consume(
		mq.SiteCheckQueue, "", false, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	for i := 0; i < concurrency; i++ {
		c.wg.Add(1)
		go func(workerID int) {
			defer c.wg.Done()
			log.Printf("Worker %d started", workerID)
			for {
				select {
				case msg, ok := <-msgs:
					if !ok {
						log.Printf("Worker %d: channel closed, exiting.", workerID)
						return
					}
					c.processMessage(msg)
				case <-c.done:
					log.Printf("Worker %d received shutdown signal, exiting.", workerID)
					return
				}
			}
		}(i)
	}
}

func (c *Consumer) processMessage(msg amqp.Delivery) {
	var site models.Site
	if err := json.Unmarshal(msg.Body, &site); err != nil {
		log.Printf("Error unmarshalling message: %v. Dropping message.", err)
		msg.Nack(false, false)
		return
	}

	err := c.checkerSvc.CheckSite(context.Background(), &site)
	if err != nil {
		log.Printf("Failed to process site %d (%s), requeueing. Error: %v", site.ID, site.Url, err)
		msg.Nack(false, true)
		return
	}

	msg.Ack(false)
}

func (c *Consumer) Shutdown() {
	close(c.done)
	c.wg.Wait()
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
}
