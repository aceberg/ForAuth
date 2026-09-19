package auth

import (
	"encoding/json"
	"maps"
	"os"

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
	saveMu.Lock()
	defer saveMu.Unlock()

	mu.RLock()
	data := maps.Clone(allSessions)
	mu.RUnlock()

	tmp := SessionsFilePath + ".tmp"

	file, err := os.Create(tmp)
	if check.IfError(err) {
		return
	}

	err = json.NewEncoder(file).Encode(data)
	check.IfError(err)
	err = file.Close()
	check.IfError(err)

	err = os.Rename(tmp, SessionsFilePath)
	check.IfError(err)
}
