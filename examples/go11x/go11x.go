package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	log.Print(port)

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}