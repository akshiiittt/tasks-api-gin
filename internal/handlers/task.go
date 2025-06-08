package handlers

import (
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
		utils.Error(c, http.StatusBadRequest, "Invalid payload", nil)
		return
	}

	// fmt.Println(task, "*******************")

	task.ID = models.NextId
	task.CreatedAt = time.Now()
	models.NextId++
	models.Tasks = append(models.Tasks, task)

	utils.Success(c, http.StatusCreated, "Task created", task)
	// c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "task created", "data": task})

}

func GetTasks(c *gin.Context) {
	utils.Success(c, http.StatusOK, "Tasks list", models.Tasks)
}

func GetTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid task id", nil)
		return
	}

	for _, v := range models.Tasks {
		if v.ID == id {
			utils.Success(c, http.StatusOK, "Task found ", v)
			return
		}
	}
	utils.Error(c, http.StatusNotFound, "Task not found", nil)
}

func UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid task id", nil)
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid payload", nil)
		return
	}

	for k, v := range models.Tasks {
		if v.ID == id {
			models.Tasks[k].Title = task.Title
			models.Tasks[k].Description = task.Description
			models.Tasks[k].Status = task.Status
			utils.Success(c, http.StatusOK, "Task updated ", models.Tasks[k])
			return
		}
	}

	utils.Error(c, http.StatusNotFound, "Task not found", nil)
}

func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid task id", nil)
		return
	}

	for k, v := range models.Tasks {
		if v.ID == id {
			models.Tasks = append(models.Tasks[:k], models.Tasks[k+1:]...)
			utils.Success(c, http.StatusOK, "Task deleted ", v)
			return
		}
	}

	utils.Error(c, http.StatusNotFound, "Task not found", nil)
}
