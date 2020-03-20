package eventbus

import (
	"sync"
)

type WorkerPool interface {
}

type Eventbus struct {
	consumers map[string][]Consumer
	producers []Producer

	messageQueue chan Messsage
	pool         WorkerPool

	shotdownChan chan interface{}

	mtx sync.RWMutex
}

func NewEventbus() *Eventbus {
	return &Eventbus{
		consumers:    make(map[string][]Consumer),
		producers:    []Producer{},
		messageQueue: make(chan Messsage, 256),
		pool:         nil,
		mtx:          sync.RWMutex{},
	}
}

func (eb *Eventbus) Run() {
	for {
		select {
		case <-eb.shotdownChan:
			return
		case message := <-eb.messageQueue:
			if err := processMessage(message); err != nil {

			}

		}

	}
}

func (eb *Eventbus) processMessage(messsage Messsage) error {

	return nil
}

func (eb *Eventbus) RegisterConsumer(topic []string, consumer ...Consumer) error {
	eb.mtx.Lock()
	defer eb.mtx.Unlock()

	return nil
}

func (eb *Eventbus) RegisterProducer(producer ...Producer) error {
	eb.producers = append(eb.producers, producer...)
	return nil
}

type Consumer interface {
	Consume(message Messsage)
}

type Producer interface {
	RegisterListenChannel(chan Messsage)
}
