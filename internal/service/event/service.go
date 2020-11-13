package event

import (
	"backendo-go/internal/config"
)

// Event ...
type Event struct {
	ID       int64
	Name     string
	DateFrom string
	DateTo   string
	Desc     string
}

// EventService
type EventService interface {
	AddEvent(Event) error
	FindByID(int) *Event
	FindAll() []*Event
}

// NewEventService ...
func NewEventService(c *config.Config) (EventService, error) {
	return service{c}, nil
}

type service struct {
	conf *config.Config
}

func (s service) AddEvent(e Event) error {
	return nil
}

func (s service) FindByID(ID int) *Event {
	return nil
}

func (s service) FindAll() []*Event {
	var list []*Event
	list = append(list, &Event{0, "event1", "20/6/2020", "20/6/2020", "sarasa"})
	return list
}
