package auth

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

var mu sync.Mutex

// StartSession for new login
func StartSession(c *gin.Context, currentAuth Conf, clientIP string) {
	var ses Session

	sessionToken := uuid.NewString()

	ses.User = currentAuth.User
	ses.Host = c.Request.Host
	ses.Expire = time.Now().Add(currentAuth.Expire)
	ses.TimeStr = ses.Expire.Format("2006-01-02 15:04:05")

	if clientIP == "" {
		clientIP = "Enable IP Info to see"
	}
	ses.ClientIP = clientIP

	mu.Lock()
	allSessions[sessionToken] = ses
	mu.Unlock()

	setTokenCookie(c, sessionToken)
}

// LogOut - log out
func LogOut(c *gin.Context) {

	sessionToken := getTokenFromCookie(c)

	delete(allSessions, sessionToken)

	setTokenCookie(c, "")
}

// LogOutByToken - log out
func LogOutByToken(token string) {

	delete(allSessions, token)
}

// GetAllSessions - get current sessions
func GetAllSessions() map[string]Session {
	return allSessions
}
