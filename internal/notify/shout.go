package notify

import (
	"log"
	"os"

	"github.com/nicholas-fedor/shoutrrr"
)

// Shout - send message with shoutrrr
func Shout(message string, url string) {
	if url != "" {
		hostname, _ := os.Hostname()
		err := shoutrrr.Send(url, "ForAuth on '"+hostname+"':\n"+message)
		if err != nil {
			log.Println("ERROR: Notification failed (shoutrrr):", err)
		}
	}
}
