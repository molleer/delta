package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var cookie_domain = os.Getenv("COOKIE_DOMAIN")

func HandleSetPassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":"ok",
	})
}

func HandleExchangeCode(c *gin.Context) {
	grant, ok := c.GetQuery("grant")
	if !ok {
		c.AbortWithError(http.StatusBadRequest, errors.New("No grant code provided"))
	}
	token, err := GetToken(grant)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	c.SetCookie("token", token.AccessToken, 600, "/", cookie_domain,true,true )
}