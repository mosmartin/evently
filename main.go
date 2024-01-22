package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mosmartin/evently/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events := models.GetEvents()

	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	event.ID = len(models.GetEvents()) + 1
	event.UserID = 1

	event.Save()

	c.JSON(http.StatusCreated, gin.H{"message": "Event successfully created", "event": event})
}
