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

	authorized.GET("/setPassword", HandleSetPassword)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}