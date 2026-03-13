package auth

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

// StartSession for new login
func StartSession(c *gin.Context, currentAuth Conf, clientIP, target string) {
	var ses Session

	sessionToken := uuid.NewString()

	ses.User = currentAuth.User
	ses.Host = c.Request.Host
	ses.Expire = time.Now().Add(currentAuth.Expire)
	ses.TimeStr = ses.Expire.Format("2006-01-02 15:04")
	ses.Target = target
	ses.Started = time.Now().Format("2006-01-02 15:04")
	ses.LastSeen = ses.Started

	if clientIP == "" {
		clientIP = "Enable IP Info to see"
	}
	ses.ClientIP = clientIP

	mu.Lock()
	allSessions[sessionToken] = ses
	mu.Unlock()
	sessionDirty = true

	setTokenCookie(c, sessionToken)
}

// LogOut - log out
func LogOut(c *gin.Context) {

	sessionToken := getTokenFromCookie(c)

	mu.Lock()
	delete(allSessions, sessionToken)
	mu.Unlock()
	sessionDirty = true

	setTokenCookie(c, "")
}

// LogOutByToken - log out
func LogOutByToken(token string) {

	mu.Lock()
	delete(allSessions, token)
	mu.Unlock()
	sessionDirty = true
}

// GetAllSessions - get current sessions
func GetAllSessions() map[string]Session {

	mu.RLock()
	ses := allSessions
	mu.RUnlock()

	return ses
}
