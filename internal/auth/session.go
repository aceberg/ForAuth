package auth

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

var mu sync.Mutex

// StartSession for new login
func StartSession(c *gin.Context) {
	var ses Session

	sessionToken := uuid.NewString()

	ses.User = authConf.User
	ses.Host = c.Request.Host
	ses.Expire = time.Now().Add(authConf.Expire)
	ses.TimeStr = ses.Expire.Format("2006-01-02 15:04:05")

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
