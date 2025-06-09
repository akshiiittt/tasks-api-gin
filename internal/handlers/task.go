package handlers

import (
	"fmt"
	"net/http"
	"task-api/internal/database"
	"task-api/internal/models"
	"task-api/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateTasks(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid payload", nil)
		return
	}

	fmt.Println(task, "*******************")
	task.ID = uuid.New()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	if err := database.DB.Create(&task).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Task not created", nil)
		return
	}

	utils.Success(c, http.StatusCreated, "Task created", task)

}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	if err := database.DB.Find(&tasks).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Could not able to retrive tasks", nil)
		return
	}

	utils.Success(c, http.StatusOK, "Tasks list", tasks)
}

func GetTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid UUID", nil)
		return
	}

	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.Error(c, http.StatusNotFound, "Task not found", nil)
		} else {
			utils.Error(c, http.StatusInternalServerError, "Error getting all tasks", nil)
		}
		return
	}

	utils.Success(c, http.StatusOK, "Tasks retrieved", task)
}

func UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid UUID", nil)
		return
	}
	var task models.Task

	if err := database.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.Error(c, http.StatusNotFound, "Task not found", nil)
		} else {
			utils.Error(c, http.StatusInternalServerError, "Error getting task", nil)
		}
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid payload", nil)
		return
	}

	task.UpdatedAt = time.Now()

	if err := database.DB.Save(&task).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Error saving task", nil)
		return
	}

	utils.Success(c, http.StatusOK, "Tasks retrived", task)
}

func DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, "Invalid UUID", nil)
		return
	}
	if err := database.DB.Delete(&models.Task{}, id).Error; err != nil {
		utils.Error(c, http.StatusInternalServerError, "Error deleting task", nil)
		return
	}

	utils.Success(c, http.StatusOK, "Task deleted ", nil)
}
