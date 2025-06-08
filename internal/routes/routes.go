package routes

import (
	"task-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	tasks := r.Group("/tasks")
	tasks.POST("/create", handlers.CreateTasks)
	tasks.GET("/tasks", handlers.GetTasks)
	tasks.GET("/task/:id", handlers.GetTask)
	tasks.PUT("/task/:id", handlers.UpdateTask)
	tasks.DELETE("/task/:id", handlers.DeleteTask)
}
