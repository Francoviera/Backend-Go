package event

import (
	"backendo-go/internal/config"

	"github.com/jmoiron/sqlx"
)

// Event ...
type Event struct {
	ID          int64
	Name        string
	Start       string
	End         string
	Description string
}

// Service ...
type Service interface {
	AddEvent(Event) error
	FindByID(int) *Event
	FindAll() []*Event
}

// NewEventService ...
func NewEventService(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

// service ...
type service struct {
	db   *sqlx.DB
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
	// list = append(list, &Event{0, "event1", "20/6/2020", "20/6/2020", "sarasa"})
	if err := s.db.Select(&list, "SELECT * FROM events"); err != nil {
		panic(err)
	}
	return list
}
