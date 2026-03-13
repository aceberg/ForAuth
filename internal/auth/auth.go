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
	mu.RLock()
	userSession, ok := allSessions[sessionToken]
	mu.RUnlock()
	exp := userSession.Expire.Before(time.Now())

	if ok && exp {
		log.Println("INFO: session for user '" + userSession.User + "' logged in from " + c.Request.RemoteAddr + " expired.")

		mu.Lock()
		delete(allSessions, sessionToken)
		mu.Unlock()

		sessionDirty = true

		ok = false
	}

	if ok && !exp {
		now := time.Now().Format("2006-01-02 15:04")

		if now != userSession.LastSeen {
			userSession.LastSeen = now

			// log.Println("NOW", now)

			mu.Lock()
			allSessions[sessionToken] = userSession
			mu.Unlock()
		}
	}

	return userSession.User, ok
}
