package events

import "time"

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}
