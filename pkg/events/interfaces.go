package events

import (
	"sync"
	"time"
)

type IEvent interface {
	GetName() string
	GetDateTime() time.Time
	GetPayLoad() any
}

type IEventHandler interface {
	Handle(event IEvent, wg *sync.WaitGroup)
}

type IEventDispacher interface {
	Register(eventName string, handler IEventHandler) error
	Dispatch(event IEvent) error
	Remove(eventName string, handler IEventHandler) error
	Has(eventName string, handler IEventHandler) bool
	Clear() error
}
