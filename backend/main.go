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
	authorized.GET("/checkLogin", func(c *gin.Context) {
		//TODO: Return user real name
		c.JSON(200, gin.H{"logged_in":true})
	})

	r.POST("/api/exchangeCode", HandleExchangeCode)
	/*r.GET("/api/ping", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("ping", fmt.Sprintf("%s ping", session.Get("ping")))
		session.Save()
		c.JSON(200, gin.H{
			"ping": true,
		})
	})
	r.GET("/api/pong", func(c *gin.Context) {
		session := sessions.Default(c)
		ping := session.Get("ping")
		log.Printf("Ping -> %s\n", ping)
		c.JSON(200, gin.H{
			"pong": true,
			"res": ping,
		})
	})*/

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}