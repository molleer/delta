package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/molleer/delta/services"
)

var cookie_domain = os.Getenv("COOKIE_DOMAIN")

func HandleSetPassword(c *gin.Context) {
	service, err := services.NewLDAPService()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	session := sessions.Default(c)
	cid := session.Get("cid").(string)

	service.AdminLogin()
	if !service.UserExist(cid) {
		c.String(http.StatusBadRequest,
			fmt.Sprintf("The user '%s' does not exist in the legacy " +
			"system, contact digIT for more support", cid))
		return
	}
	var body struct {
		Password string `json:"password"`
	}

	text, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(text, &body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if body.Password == "" {
		c.String(http.StatusBadRequest, "The password cannot be empty")
		return
	}

	err = service.SetPassword(cid, body.Password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Status(http.StatusOK)
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

	var body struct {
		Cid string `json:"cid"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
	}

	err = services.GammaGet("/api/users/me", token.AccessToken, &body)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	session := sessions.Default(c)
	session.Set("token", token.AccessToken)
	session.Set("name", fmt.Sprintf("%s %s", body.FirstName, body.LastName))
	session.Set("cid", body.Cid)
	session.Save()

	c.JSON(200, gin.H{
		"logged_in":true, 
		"name": fmt.Sprintf("%s %s", body.FirstName, body.LastName),
		"cid": body.Cid,
	})
}

func HandleCheckLogin(c *gin.Context) {
	session := sessions.Default(c)
	c.JSON(200, gin.H{
		"logged_in":true, 
		"name": session.Get("name"),
		"cid": session.Get("cid"),
	})
}

func HandleLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}