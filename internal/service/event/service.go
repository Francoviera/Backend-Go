package event

import (
	"backendo-go/internal/config"
	"database/sql"
	"errors"

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
	AddEvent(Event) (int64, error)
	FindByID(int) (*Event, error)
	FindAll() []*Event
	Delete(int) error
	Put(Event, int) error
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

func (s service) AddEvent(e Event) (int64, error) {

	insertEvent := `INSERT INTO events (name, start, end, description) VALUES (?,?,?,?);`
	id, err := s.db.MustExec(insertEvent, e.Name, e.Start, e.End, e.Description).LastInsertId()

	if err != nil {
		return -1, errors.New("Hubo error al agregar")
	}

	return id, err
}

func (s service) Put(e Event, ID int) error {

	err := s.db.QueryRow("UPDATE events SET name=?, start=?, end=?, description=? WHERE id=?", e.Name, e.Start, e.End, e.Description, ID).Scan()

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (s service) FindByID(ID int) (*Event, error) {
	var event Event

	err := s.db.QueryRowx("SELECT * FROM events WHERE id=?", ID).StructScan(&event)
	if err != nil {
		return nil, err
	}

	return &event, nil

}

func (s service) Delete(ID int) error {
	_, err := s.db.Exec("DELETE FROM events WHERE id=?", ID)
	if err != nil {
		return err
	}

	return nil

}

func (s service) FindAll() []*Event {
	var list []*Event
	// list = append(list, &Event{0, "event1", "20/6/2020", "20/6/2020", "sarasa"})
	if err := s.db.Select(&list, "SELECT * FROM events"); err != nil {
		return nil
	}
	return list
}
