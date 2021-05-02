package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/molleer/delta/services"
)

var GAMMA_REDIRECT_URL = os.Getenv("GAMMA_REDIRECT_URL")
var COOKIE_DOMAIN = os.Getenv("COOKIE_DOMAIN")

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("token")
		if token == nil || !services.TokenIsValid(token.(string)) {
			session.Clear()
			c.String(401, services.GetLoginURL())
			c.Abort()
			return
		}
		c.Next()
	}
}

func RedisSession() gin.HandlerFunc {
	store, err := redis.NewStore(10, "tcp", 
		fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
		os.Getenv("REDIS_PASS"),
		[]byte(os.Getenv("SESSION_SECRET")))

	store.Options(sessions.Options{
		Path: "/",
		Domain: COOKIE_DOMAIN,
		MaxAge:	600,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	if err != nil {
		fmt.Println("Failed to connect to redis")
		panic(err)
	}
	
	return sessions.Sessions("delta", store)
}