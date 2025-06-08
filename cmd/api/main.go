package main

import (
	"fmt"
	"net/http"
	"task-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello")

	r := gin.Default()

	routes.TaskRoutes(r)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "server is running",
		})
	})

	r.Run()
}
