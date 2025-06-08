package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"task-api/internal/models"
	"task-api/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTasks(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.Error(c, http.StatusBadRequest, "error", "nil")
		return
	}

	fmt.Println(task, "*******************")

	task.ID = models.NextId
	task.CreatedAt = time.Now()
	models.NextId++
	models.Tasks = append(models.Tasks, task)

	utils.Success(c, http.StatusCreated, "task created", task)
	// c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "task created", "data": task})

}

func GetTasks(c *gin.Context) {
	utils.Success(c, http.StatusAccepted, "tasks list", models.Tasks)
}

func GetTask(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	for _, v := range models.Tasks {
		if v.ID == id {
			utils.Success(c, http.StatusOK, "task ", v)
			return
		}
	}
	utils.Success(c, http.StatusAccepted, "task", models.Tasks)
}

// func UpdateTask() {

// }
// func DeleteTask() {

// }
