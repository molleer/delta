package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authorized := r.Group("/api/admin")
	authorized.Use(AuthRequired())

	authorized.POST("/setPassword", HandleSetPassword)
	authorized.GET("/checkLogin", func(c *gin.Context) {
		//TODO: Return user real name
		c.JSON(200, gin.H{"logged_in":true})
	})

	r.POST("/api/exchangeCode", HandleExchangeCode)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}