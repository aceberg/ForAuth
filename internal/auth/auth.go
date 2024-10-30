package auth

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Auth - main auth func
func Auth(c *gin.Context, conf *Conf) bool {

	sessionToken := getTokenFromCookie(c)

	userSession, exists := allSessions[sessionToken]

	if !exists || userSession.Before(time.Now()) {
		delete(allSessions, sessionToken)
		return false
	}

	userSession = time.Now().Add(conf.Expire)
	allSessions[sessionToken] = userSession

	return true
}
