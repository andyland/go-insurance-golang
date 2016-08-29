package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("GO_INSURANCE_TEST_PORT")

	if port == "" {
		log.Fatal("$GO_INSURANCE_TEST_PORT must be set")
	}

	setupDB()

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", index)
	router.GET("/users/:username/appointment", getAppointment)
	router.POST("/users/:username/appointment", createAppointment)
	router.DELETE("/users/:username/appointment", deleteAppointment)

	router.Run(":" + port)
}
