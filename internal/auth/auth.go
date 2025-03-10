package auth

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Auth - main auth func
func Auth(c *gin.Context, conf *Conf) bool {

	if !conf.Auth || conf.User == "" || conf.Password == "" {
		return true
	}

	authConf = *conf
	sessionToken := getTokenFromCookie(c)

	userSession, exists := allSessions[sessionToken]
	exp := userSession.Before(time.Now())

	if exists && !exp {
		return true
	}

	if exists && exp {
		log.Println("INFO: session for user '" + authConf.User + "' logged in from " + c.Request.RemoteAddr + " expired.")
		delete(allSessions, sessionToken)
	}

	return false
}
