package auth

import (
	"time"
)

// Conf - auth config
type Conf struct {
	Auth     bool
	User     string
	Password string
	ExpStr   string
	Expire   time.Duration
}

var authConf Conf

// Session - one session
type Session struct {
	User    string
	Host    string
	Expire  time.Time
	TimeStr string
}

var allSessions = make(map[string]Session)

var cookieName = "forauth_session_token"
