package auth

import (
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

	if !exists || userSession.Before(time.Now()) {
		delete(allSessions, sessionToken)
		return false
	}

	return true
}
