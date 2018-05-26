package event

import "time"

//Message ...
type Message interface {
	Key() string
}

//MeowCreatedMessage ...
type MeowCreatedMessage struct {
	ID        string
	Body      string
	CreatedAt time.Time
}

//Key ...
func (m *MeowCreatedMessage) Key() string {
	return "meow.created"
}
