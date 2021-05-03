package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(RedisSession())

	authorized := r.Group("/api/admin")
	authorized.Use(AuthRequired())

	authorized.POST("/setPassword", HandleSetPassword)
	authorized.GET("/checkLogin", HandleCheckLogin)

	r.POST("/api/logout", HandleLogout)
	r.POST("/api/exchangeCode", HandleExchangeCode)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}