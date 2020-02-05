package events

import (
	"encoding/json"
	"net/http"
	"net/url"

	api "github.com/lusis/gochimp/mandrill/api/events"
)

const (
	// BlacklistEventType is the string for matching blacklist sync events
	BlacklistEventType = "blacklist"
	// WhitelistEventType is the string for matching whitelist sync events
	WhitelistEventType = "whitelist"
	// SendEventType is the string for matching send message events
	SendEventType = "send"
	// DeferralEventType is the string for matching deferral message events
	DeferralEventType = "deferral"
	// HardBounceEventType is the string for matching hard_bounce message events
	HardBounceEventType = "hard_bounce"
	// SoftBounceEventType is the string for matching soft_bounce message events
	SoftBounceEventType = "soft_bounce"
	// OpenEventType is the string for matching open message events
	OpenEventType = "open"
	// ClickEventType is the string for matching click message events
	ClickEventType = "click"
	// SpamEventType is the string for matching spam message events
	SpamEventType = "spam"
	// UnsubEventType is the string for matching unsub message events
	UnsubEventType = "unsub"
	// RejectEventType is the string for matching reject message events
	RejectEventType = "reject"
	// InboundMessageEventType is the string for matching inbound message events
	InboundMessageEventType = "inbound"
)

var jsonMessageEventMapping = map[string]interface{}{
	SendEventType:       api.SendMessageEvent{},
	DeferralEventType:   api.DeferralMessageEvent{},
	HardBounceEventType: api.HardBounceMessageEvent{},
	SoftBounceEventType: api.SoftBounceMessageEvent{},
	OpenEventType:       api.OpenMessageEvent{},
	ClickEventType:      api.ClickMessageEvent{},
	SpamEventType:       api.SpamMessageEvent{},
	UnsubEventType:      api.UnsubMessageEvent{},
	RejectEventType:     api.RejectMessageEvent{},
}

var messageEventMapping = map[string]interface{}{
	SendEventType:       SendEvent{},
	DeferralEventType:   DeferralEvent{},
	HardBounceEventType: HardBounceEvent{},
	SoftBounceEventType: SoftBounceEvent{},
	OpenEventType:       OpenEvent{},
	ClickEventType:      ClickEvent{},
	SpamEventType:       SpamEvent{},
	UnsubEventType:      UnsubEvent{},
	RejectEventType:     RejectEvent{},
}

var jsonSyncEventMapping = map[string]interface{}{
	WhitelistEventType: api.WhitelistMessageEvent{},
	BlacklistEventType: api.BlacklistMessageEvent{},
}

// ParseInnerEvent parses the inner event
// this is provided as an alternate mechanism
func ParseInnerEvent(wh WebhookEvent) (interface{}, error) {
	switch wh.Type {
	case MessageEventType:
		return parseMessageEvent(wh)
	case SyncEventType:
		return parseSyncEvent(wh)
	default:
		return nil, InvalidEventType{eventType: wh.InnerEventType}
	}
}

// ParseOuterEvent parses the event and returns basic information about it
func ParseOuterEvent(b []byte) (WebhookEvent, error) {
	wh, err := parseOuterEvent(json.RawMessage(b))
	return wh, err
}

// ParseEvent parses the event and returns the final event
// i.e. BlacklistEvent or SendEvent
func ParseEvent(b []byte) (interface{}, error) {
	wh, err := ParseOuterEvent(json.RawMessage(b))
	if err != nil {
		return nil, err
	}
	return ParseInnerEvent(wh)
}

// ParseEvents parses a collection of events
func ParseEvents(b []byte) ([]WebhookEvent, error) {
	evts := []json.RawMessage{}
	res := []WebhookEvent{}
	err := json.Unmarshal(b, &evts)
	if err != nil {
		return nil, UnmarshallError{data: string(b), msg: err.Error()}
	}
	for _, e := range evts {
		evt, err := ParseOuterEvent(e)
		if err != nil {
			return nil, UnmarshallError{data: string(e), msg: err.Error()}
		}
		res = append(res, evt)
	}
	return res, nil
}

// ParseRequest parses the http request directly and returns a slice of WebhookEvent for processing
func ParseRequest(r *http.Request) ([]WebhookEvent, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, &ErrParse{parseErr: err.Error()}
	}
	events := r.Form.Get("mandrill_events")
	if len(events) == 0 {
		return nil, &ErrMissingMandrillEvents{postForm: r.Form.Encode()}
	}
	m, err := url.QueryUnescape(events)
	if err != nil {
		return nil, &ErrParse{parseErr: err.Error()}
	}
	return ParseEvents([]byte(m))
}
