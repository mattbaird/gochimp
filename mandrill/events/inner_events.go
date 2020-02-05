package events

import (
	"time"
)

// InnerEvent is the inner event type
type InnerEvent struct {
	Type      string
	Timestamp time.Time
}
