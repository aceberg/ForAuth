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

	sessionToken := uuid.NewString()

	mu.Lock()
	allSessions[sessionToken] = time.Now().Add(authConf.Expire)
	mu.Unlock()

	setTokenCookie(c, sessionToken)

	// c.Redirect(http.StatusFound, "/")
}

// LogOut - log out
func LogOut(c *gin.Context) {

	sessionToken := getTokenFromCookie(c)

	delete(allSessions, sessionToken)

	setTokenCookie(c, "")

	// c.Redirect(http.StatusFound, "/")
}
