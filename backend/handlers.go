package main

import "github.com/gin-gonic/gin"

func HandleSetPassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":"ok",
	})
}