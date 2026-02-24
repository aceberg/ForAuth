package auth

import (
	"encoding/json"
	"os"

	"github.com/aceberg/ForAuth/internal/check"
)

// RestoreSessions - restore sessions from file
func RestoreSessions() {
	file, err := os.ReadFile(SessionsFilePath)
	check.IfError(err)

	allSessions = make(map[string]Session)
	err = json.Unmarshal(file, &allSessions)
	check.IfError(err)
}

// SaveSessions - save sessions to file
func SaveSessions() {
	jsonData, err := json.MarshalIndent(allSessions, "", "  ")
	check.IfError(err)

	err = os.WriteFile(SessionsFilePath, jsonData, 0644)
	check.IfError(err)
}
