package routes

import (
	"task-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	tasks := r.Group("/tasks")
	tasks.POST("/", handlers.CreateTasks)
	tasks.GET("/", handlers.GetTasks)
	tasks.GET("/:id", handlers.GetTask)
	tasks.PUT("/:id", handlers.UpdateTask)
	tasks.DELETE("/:id", handlers.DeleteTask)
}
