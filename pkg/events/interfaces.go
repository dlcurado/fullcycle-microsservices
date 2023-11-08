package events

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
	SetPayload(interface{})
}

type EventHandlerInterace interface {
	Handle(event EventInterface)
}

type EventDispacherInterface interface {
	Register(eventName string, handler EventHandlerInterace) error
	Remove(eventName string, handler EventHandlerInterace) error
	Has(eventName string, handler EventHandlerInterace) bool
	Clear()

	Dispatch(event EventInterface) error
}
