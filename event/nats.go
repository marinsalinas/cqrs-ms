package event

import (
	"bytes"
	"encoding/gob"

	"github.com/marinsalinas/cqrs-ms/schema"
	"github.com/nats-io/go-nats"
)

//NatsEventStore ...
type NatsEventStore struct {
	nc                      *nats.Conn
	meowCreatedSubscription *nats.Subscription
	meowCreatedChan         chan MeowCreatedMessage
}

//NewNats ...
func NewNats(url string) (*NatsEventStore, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &NatsEventStore{nc: nc}, nil
}

//Close ...
func (e *NatsEventStore) Close() {
	if e.nc != nil {
		e.nc.Close()
	}
	if e.meowCreatedSubscription != nil {
		e.meowCreatedSubscription.Unsubscribe()
	}
	close(e.meowCreatedChan)
}

//PublishMeowCreated ...
func (e *NatsEventStore) PublishMeowCreated(meow schema.Meow) error {
	m := MeowCreatedMessage{meow.ID, meow.Body, meow.CreatedAt}
	data, err := e.writeMessage(&m)
	if err != nil {
		return err
	}
	return e.nc.Publish(m.Key(), data)
}

//OnMeowCreated implements the subscription function with callback function
func (e *NatsEventStore) OnMeowCreated(f func(MeowCreatedMessage)) (err error) {
	m := MeowCreatedMessage{}
	e.meowCreatedSubscription, err = e.nc.Subscribe(m.Key(), func(msg *nats.Msg) {
		e.readMessage(msg.Data, &m)
		f(m)
	})
	return
}

//SubscribeMeowCreated implements the subscription function with a channel
func (e *NatsEventStore) SubscribeMeowCreated() (<-chan MeowCreatedMessage, error) {
	m := MeowCreatedMessage{}
	e.meowCreatedChan = make(chan MeowCreatedMessage, 64)
	ch := make(chan *nats.Msg, 64)
	var err error
	e.meowCreatedSubscription, err = e.nc.ChanSubscribe(m.Key(), ch)
	if err != nil {
		return nil, err
	}
	// Decode message
	go func() {
		for {
			select {
			case msg := <-ch:
				e.readMessage(msg.Data, &m)
				e.meowCreatedChan <- m
			}
		}
	}()
	return (<-chan MeowCreatedMessage)(e.meowCreatedChan), nil
}

func (e *NatsEventStore) readMessage(data []byte, m interface{}) error {
	b := bytes.Buffer{}
	b.Write(data)
	return gob.NewDecoder(&b).Decode(m)
}

func (e *NatsEventStore) writeMessage(m Message) ([]byte, error) {
	b := bytes.Buffer{}
	err := gob.NewEncoder(&b).Encode(m)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
