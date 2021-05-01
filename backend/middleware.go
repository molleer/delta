package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var GAMMA_REDIRECT_URL = os.Getenv("GAMMA_REDIRECT_URL")

func TokenIsValid(token string) bool{
	//TODO: Send request to Gamma and check token
	return token != ""
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil || !TokenIsValid(token) {
			c.String(401, GAMMA_REDIRECT_URL)
			c.Abort()
			return
		}
		c.Next()
	}
}