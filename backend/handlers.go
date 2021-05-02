package main

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/molleer/delta/services"
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
	token, err := services.GetToken(grant)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	session := sessions.Default(c)
	session.Set("token", token.AccessToken)
	session.Save()
}