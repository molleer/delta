package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var GAMMA_REDIRECT_URL = os.Getenv("GAMMA_REDIRECT_URL")

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		log.Println(token)
		if err != nil || !TokenIsValid(token) {
			c.String(401, GetLoginURL())
			c.Abort()
			return
		}
		c.Next()
	}
}