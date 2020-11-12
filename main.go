package main

func main() {

	//CRUD

	el := NewEventList()
	el.Add(Event{0, "event1", "20/6/2020", "20/6/2020", "sarasa"})

	el.Print()
}
