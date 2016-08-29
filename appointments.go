package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func createAppointment(c *gin.Context) {
	username := c.Params.ByName("username")
	if len(username) <= 0 {
		c.String(400, "No username specified")
		return
	}
	date := c.Query("date")
	if len(date) <= 0 {
		c.String(400, "No date specified")
		return
	}
	parsedDate, error := time.Parse("2006-01-02", date)
	if nil != error {
		c.String(400, "Invalid date")
		return
	}
	if parsedDate.Before(time.Now()) {
		c.String(400, "Date has already passed")
		return
	}
	timeOfDay := c.Query("time_of_day")
	if !timeIsValid(timeOfDay) {
		c.String(400, "Invalid time of day")
		return
	}
	if username == getUserAppointment(username).username {
		c.String(409, "User already has an appointment")
		return
	}
	createUserAppointment(Appointment{username, parsedDate, timeOfDay})
	c.JSON(201, gin.H{
		"username":    username,
		"date":        parsedDate,
		"time_of_day": timeOfDay,
	})
}

func deleteAppointment(c *gin.Context) {
	username := c.Params.ByName("username")
	if len(username) <= 0 {
		c.String(400, "No username specified")
		return
	}
	deleteUserAppointment(username)
	c.Status(204)
}

func getAppointment(c *gin.Context) {
	username := c.Params.ByName("username")
	if len(username) <= 0 {
		c.String(400, "No username specified")
		return
	}
	appointment := getUserAppointment(username)
	if username == appointment.username {
		c.JSON(200, gin.H{
			"username":    appointment.username,
			"date":        appointment.date,
			"time_of_day": appointment.timeOfDay,
		})
		return
	}
	c.Status(404)
}
