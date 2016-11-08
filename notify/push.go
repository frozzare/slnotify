package notify

import (
	"time"

	"github.com/frozzare/slnotify/config"
	"github.com/jdiez17/go-pushover"
)

// Push will push a title with message to Pushover.
func Push(title string, message string) error {
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

	if _, err := p.Notify(n); err != nil {
		return err
	}

	return nil
}
