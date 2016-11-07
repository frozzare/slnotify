package notify

import (
	"fmt"
	"time"

	"github.com/frozzare/sl/config"
	"github.com/jdiez17/go-pushover"
)

// Push will push a title with message to Pushover.
func Push(title string, message string) {
	c := config.Get()

	p := pushover.Pushover{
		UserKey: c.Pushover.UserKey,
		AppKey:  c.Pushover.AppKey,
	}

	n := pushover.Notification{
		Title:     title,
		Message:   message,
		Timestamp: time.Now(),
		Priority:  0,
		Retry:     0,
		Expire:    90,
	}

	response, err := p.Notify(n)

	if err != nil {
		if err != pushover.PushoverError {
			panic(err)
		} else {
			fmt.Println(err)
			fmt.Println(response)
		}
	}
}
