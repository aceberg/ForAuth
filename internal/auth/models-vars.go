package auth

import (
	"sync"
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
	User     string    `json:"User"`
	Host     string    `json:"Host"`
	Expire   time.Time `json:"Expire"`
	TimeStr  string    `json:"TimeStr"`
	ClientIP string    `json:"ClientIP"`
	Target   string    `json:"Target"`
	Started  string    `json:"Started"`
	LastSeen string    `json:"-"`
}

var mu sync.RWMutex

var allSessions = make(map[string]Session)

var cookieName = "forauth_session_token"

// SessionsFilePath - path to sessions.json file
var SessionsFilePath string

var sessionDirty bool
