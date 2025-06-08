package utils

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, statusCode int, message string, data interface{}){
	c.JSON(statusCode, gin.H{
		"status": "success",
		"message": message,
		"data": data,
	})
}

func Error(c *gin.Context, statusCode int,  message string, data interface{}){
	c.JSON(statusCode, gin.H{
		"status": "error",
		"message": message,
		"data": data,
	})
}