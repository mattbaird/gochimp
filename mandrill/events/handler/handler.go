package handler

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/lusis/gochimp/mandrill/events"
)

// MessageEventFunc is a function that can process an events.MessageEvent
type MessageEventFunc func(context.Context, events.MessageEvent)

// SyncEventFunc is a function that can process an events.SyncEvent
type SyncEventFunc func(context.Context, events.SyncEvent)

// UnknownEventFunc is a function that can process any events that don't match
type UnknownEventFunc func(context.Context, events.WebhookEvent)

// Option is a functional type for setting options
type Option func(*EventHandler) error

// WithValidation enabled validation on the handler
// requires the exact url use for the webhook entry
// and the key
func WithValidation(url, key string) Option {
	return func(e *EventHandler) error {
		e.url = url
		e.key = key
		e.validate = true
		return nil
	}
}

// WithLogger sets the logger to use
func WithLogger(l *log.Logger) Option {
	return func(e *EventHandler) error {
		e.logger = l
		return nil
	}
}

// WithDebug enables debugging
func WithDebug() Option {
	return func(e *EventHandler) error {
		e.debug = true
		return nil
	}
}

// EventHandler is a http handler for handling Mandrill webhooks
type EventHandler struct {
	url             string
	key             string
	syncHandlers    []SyncEventFunc
	messageHandlers []MessageEventFunc
	unknownHandlers []UnknownEventFunc
	validate        bool
	lock            sync.RWMutex
	logger          *log.Logger
	debug           bool
}

// NewEventHandler returns a new event handler
func NewEventHandler(opts ...Option) (*EventHandler, error) {
	e := &EventHandler{}
	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}
	if len(e.key) == 0 || len(e.url) == 0 {
		e.validate = false
	}
	if e.logger == nil {
		e.logger = log.New(os.Stdout, "mandrill_events", log.LstdFlags)
	}
	e.messageHandlers = []MessageEventFunc{}
	e.unknownHandlers = []UnknownEventFunc{}
	e.syncHandlers = []SyncEventFunc{}
	return e, nil
}

// Handle handles http requests from mandrill webhooks
func (e *EventHandler) Handle(w http.ResponseWriter, r *http.Request) {
	evts, err := events.ParseRequest(r)
	if err != nil {
		e.logger.Printf("error parsing request: %s", err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if e.validate {
		if err := events.ValidateWebhookSignature(e.url, e.key, r); err != nil {
			e.logger.Printf("error validating signature: %s", err.Error())
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	failedParses := []events.WebhookEvent{}
	messageParses := []events.MessageEvent{}
	syncParses := []events.SyncEvent{}

	// Parse all events into buckets
	for _, evt := range evts {
		switch evt.Type {
		case events.MessageEventType:
			me, err := events.ParseInnerEvent(evt)
			if err != nil {
				failedParses = append(failedParses, evt)
			} else {
				messageParses = append(messageParses, me.(events.MessageEvent))
			}
		case events.SyncEventType:
			se, err := events.ParseInnerEvent(evt)
			if err != nil {
				failedParses = append(failedParses, evt)
			} else {
				syncParses = append(syncParses, se.(events.SyncEvent))
			}
		default:
			failedParses = append(failedParses, evt)
		}
	}

	for _, p := range failedParses {
		for _, f := range e.unknownHandlers {
			e.lock.RLock()
			defer e.lock.RUnlock()
			go f(r.Context(), p)
		}
	}
	for _, p := range messageParses {
		for _, f := range e.messageHandlers {
			e.lock.RLock()
			defer e.lock.RUnlock()
			go f(r.Context(), p)
		}
	}
	for _, p := range syncParses {
		for _, f := range e.syncHandlers {
			e.lock.RLock()
			defer e.lock.RUnlock()
			go f(r.Context(), p)

		}
	}
	w.WriteHeader(http.StatusOK)
}

// AddMessageHandler adds a function to process events.MessageEvent
func (e *EventHandler) AddMessageHandler(f ...MessageEventFunc) {
	e.lock.Lock()
	e.messageHandlers = append(e.messageHandlers, f...)
	e.lock.Unlock()
}

// AddSyncHandler adds a function to process events.SyncEvent
func (e *EventHandler) AddSyncHandler(f ...SyncEventFunc) {
	e.lock.Lock()
	e.syncHandlers = append(e.syncHandlers, f...)
	e.lock.Unlock()
}

// AddUnknownHandler adds a function to process unknown events
func (e *EventHandler) AddUnknownHandler(f ...UnknownEventFunc) {
	e.lock.Lock()
	e.unknownHandlers = append(e.unknownHandlers, f...)
	e.lock.Unlock()
}
