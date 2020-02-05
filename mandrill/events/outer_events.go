package events

import (
	"encoding/json"
	"time"

	api "github.com/lusis/gochimp/mandrill/api/events"
)

const (
	// SyncEventType is the outer wrapping type for sync events
	SyncEventType = "sync"
	// MessageEventType is the outer wrapping for message events
	MessageEventType = "message"
)

// WebhookEvent represents the outer wrapper of a Mandrill event
type WebhookEvent struct {
	Type           string // i.e. message or sync
	InnerEventType string // i.e. send or whitelist
	Timestamp      time.Time
	raw            json.RawMessage
}

// parseOuterEvent parses the minimum amount of data from the
// event to determine if it's a Sync event or a Message event
// and the type of each
func parseOuterEvent(j json.RawMessage) (WebhookEvent, error) {
	res := api.MandrillEventJSON{}
	e := WebhookEvent{}
	err := json.Unmarshal(j, &res)
	if err != nil {
		return e, UnmarshallError{msg: err.Error(), data: string(j)}
	}
	e.raw = j
	e.Timestamp = res.TS.Time
	if len(res.Event) == 0 {
		// Event is empty so this is a sync event
		_, ok := jsonSyncEventMapping[res.Type]
		if !ok {
			return e, InvalidEventType{eventType: res.Type}
		}
		e.Type = SyncEventType
		e.InnerEventType = res.Type
	} else {
		_, ok := jsonMessageEventMapping[res.Event]
		if !ok {
			return e, InvalidEventType{eventType: res.Event}
		}
		e.Type = MessageEventType
		e.InnerEventType = res.Event
	}
	return e, nil
}
