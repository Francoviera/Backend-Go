package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.GET("/events", getEvents)
	r.GET("/events/:eventName", getEvent)
	r.Run()

	//CRUD
	// el := NewEventList()
	// el.Add(Event{0, "event1", "20/6/2020", "20/6/2020", "sarasa"})

	// el.Print()
}

func getEvents(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
func getEvent(c *gin.Context) {
	eventName := c.Param("eventName")
	c.JSON(200, gin.H{
		"status": "ok",
		"name":   eventName,
	})
}
