package auth

import (
	"time"
)

// Conf - auth config
type Conf struct {
	Auth     bool          `yaml:"enabled"`
	User     string        `yaml:"username"`
	Password string        `yaml:"password"`
	ExpStr   string        `yaml:"expire"`
	Expire   time.Duration `yaml:"-"`
}

// Session - one session
type Session struct {
	User     string
	Host     string
	Expire   time.Time
	TimeStr  string
	ClientIP string
}

var allSessions = make(map[string]Session)

var cookieName = "forauth_session_token"
