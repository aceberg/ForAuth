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

	user, exists := GetCurrentUser(c)

	if exists && (user == conf.User) {
		return true
	}

	return false
}

// GetCurrentUser - get current session user from cookie
func GetCurrentUser(c *gin.Context) (string, bool) {
	var userSession Session

	sessionToken := getTokenFromCookie(c)
	userSession, ok := allSessions[sessionToken]
	exp := userSession.Expire.Before(time.Now())

	if ok && exp {
		log.Println("INFO: session for user '" + userSession.User + "' logged in from " + c.Request.RemoteAddr + " expired.")

		mu.Lock()
		delete(allSessions, sessionToken)
		SaveSessions()
		mu.Unlock()

		ok = false
	}

	return userSession.User, ok
}
