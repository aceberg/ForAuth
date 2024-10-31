package notify

import (
	"github.com/containrrr/shoutrrr"
	"log"
)

// Shout - send message with shoutrrr
func Shout(message string, url string) {
	if url != "" {
		err := shoutrrr.Send(url, message)
		if err != nil {
			log.Println("ERROR: Notification failed (shoutrrr):", err)
		}
	}
}
