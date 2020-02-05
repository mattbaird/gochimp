package events

import (
	"encoding/json"
	"strings"
	"time"

	api "github.com/lusis/gochimp/mandrill/api/events"
)

// SyncEvent represents a basic sync event
type SyncEvent struct {
	Data interface{}
}

func parseSyncEvent(e WebhookEvent) (SyncEvent, error) {
	se := SyncEvent{}
	if e.Type != SyncEventType {
		return se, InvalidEventType{eventType: e.Type}
	}
	decoder := json.NewDecoder(strings.NewReader(string(e.raw)))
	decoder.DisallowUnknownFields()
	var evt interface{}
	switch e.InnerEventType {
	case WhitelistEventType:
		apiEvt := api.WhitelistMessageEvent{}
		if err := decoder.Decode(&apiEvt); err != nil {
			return se, UnmarshallError{data: string(e.raw), msg: err.Error()}
		}
		evt = WhitelistEvent{
			Action:    apiEvt.Action,
			Timestamp: apiEvt.TS.Time,
			Entry: WhitelistDetails{
				Email:     apiEvt.Entry.Email,
				Detail:    apiEvt.Entry.Detail,
				CreatedAt: apiEvt.Entry.CreatedAt.Time,
			},
		}
	case BlacklistEventType:
		apiEvt := api.BlacklistMessageEvent{}
		if err := decoder.Decode(&apiEvt); err != nil {
			return se, UnmarshallError{data: string(e.raw), msg: err.Error()}
		}
		evt = BlacklistEvent{
			Action:    apiEvt.Action,
			Timestamp: apiEvt.TS.Time,
			Reject: BlacklistDetails{
				Reason:      apiEvt.Reject.Reason,
				Detail:      apiEvt.Reject.Detail,
				LastEventAt: apiEvt.Reject.LastEventAt.Time,
				Email:       apiEvt.Reject.Email,
				CreatedAt:   apiEvt.Reject.CreatedAt.Time,
				ExpiresAt:   apiEvt.Reject.ExpiresAt.Time,
				Expired:     apiEvt.Reject.Expired,
				SubAccount:  apiEvt.Reject.SubAccount,
				Sender:      apiEvt.Reject.Sender,
			},
		}
	default:
		return se, InvalidEventType{eventType: e.InnerEventType}
	}
	se.Data = evt
	return se, nil
}

// WhitelistEvent is a sync event for whitelists
// https://mandrill.zendesk.com/hc/en-us/articles/205583297-Sync-Event-Webhook-format
type WhitelistEvent struct {
	Timestamp time.Time
	Action    string
	Entry     WhitelistDetails
}

// WhitelistDetails is the Entry section of a whitelist event
type WhitelistDetails struct {
	Email     string
	Detail    string
	CreatedAt time.Time
}

// BlacklistEvent is a sync event for blacklists
// https://mandrill.zendesk.com/hc/en-us/articles/205583297-Sync-Event-Webhook-format
type BlacklistEvent struct {
	Timestamp time.Time
	Action    string
	Reject    BlacklistDetails
}

// BlacklistDetails is the Reject section of a blacklist event
type BlacklistDetails struct {
	Reason      string
	Detail      string
	LastEventAt time.Time
	Email       string
	CreatedAt   time.Time
	ExpiresAt   time.Time
	Expired     bool
	SubAccount  string
	Sender      string
}
