package auth

import (
	"encoding/json"
	"os"
	"time"

	"github.com/aceberg/ForAuth/internal/check"
)

// RestoreSessions - restore sessions from file
func RestoreSessions() {
	file, err := os.ReadFile(SessionsFilePath)
	check.IfError(err)

	mu.Lock()
	allSessions = make(map[string]Session)
	err = json.Unmarshal(file, &allSessions)
	mu.Unlock()

	check.IfError(err)
}

// SaveSessions - save sessions to file
func SaveSessions() {
	mu.RLock()
	jsonData, err := json.MarshalIndent(allSessions, "", "  ")
	mu.RUnlock()
	check.IfError(err)

	err = os.WriteFile(SessionsFilePath, jsonData, 0644)
	check.IfError(err)
}

// SessionWriter - save sessions every N seconds
func SessionWriter() {
	for {
		time.Sleep(5 * time.Second)

		if sessionDirty {

			SaveSessions()
			sessionDirty = false

			// log.Println("Writing to sessions.json")
		}
	}
}
