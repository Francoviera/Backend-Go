package main

// fmt ...
import "fmt"

// EventList ...
type EventList struct {
	events map[int]*Event
}

// Event ...
type Event struct {
	ID       int
	Name     string
	DateFrom string
	DateTo   string
	Desc     string
}

// NewEventList ...
func NewEventList() EventList {
	events := make(map[int]*Event)
	return EventList{
		events,
	}
}

// Add ...
func (el EventList) Add(e Event) {
	el.events[e.ID] = &e
}

// Print ...
func (el EventList) Print() {
	for _, e := range el.events {
		fmt.Print(e.ID, e.Name)
	}
}

//FindByID ...
func (el EventList) FindByID(ID int) *Event {
	return el.events[ID]
}
