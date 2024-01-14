package app

import (
	"swclabs/swiftcart/internal/config"
	"swclabs/swiftcart/internal/messaging"
	"swclabs/swiftcart/pkg/worker"
)

func init() {
	worker.SetBroker(config.RedisHost, config.RedisPort, config.RedisPassword)
}

type IWorker interface {
	Run() error
}

type Worker struct {
	engine *messaging.Messaging
}

func NewWorker(concurrency int) IWorker {
	return &Worker{
		engine: messaging.NewMessaging(),
	}
}

func (w *Worker) Run() error {
	return w.engine.Run(10)
}
