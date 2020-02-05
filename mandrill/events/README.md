# Mandrill Webhook Events

_This package is in a bit of flux as it's being developed_

This package exclusively deals with handling Mandrill's webhooks.

## How to use

There are a couple of ways to use this package:

### Using the custom handler

This is the easiest way as it cuts down on a lot of boiler plate code:

```go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lusis/gochimp/mandrill/events"
	"github.com/lusis/gochimp/mandrill/events/handler"
)

func messageHandler(ctx context.Context, evt events.MessageEvent) {
	log.Printf("got message event: %v", evt)
}

func syncHandler(ctx context.Context, evt events.SyncEvent) {
	log.Printf("got sync event: %v", evt)
}

func unknownHandler(ctx context.Context, evt events.WebhookEvent) {
	log.Printf("got an unknown event: %v", evt)
}

func sendHandler(ctx context.Context, evt events.MessageEvent) {
	se, ok := evt.Data.(events.SendEvent)
	if ok {
		log.Printf("got a send event from %s to %s with subject %s and id %s",
			se.Msg.Sender, se.Msg.Email, se.Msg.Subject, se.ID,
		)
	}
}

func blacklistHandler(ctx context.Context, evt events.SyncEvent) {
	be, ok := evt.Data.(events.BlacklistEvent)
	if ok {
		log.Printf("got a blacklist event with action %s", be.Action)
	}
}

func headFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	eventHandler, err := handler.NewEventHandler()
	if err != nil {
		log.Fatal(err.Error())
	}
	eventHandler.AddMessageHandler(messageHandler, sendHandler)
	eventHandler.AddSyncHandler(syncHandler, blacklistHandler)
	eventHandler.AddUnknownHandler(unknownHandler)

	mux := mux.NewRouter()
	mux.Handle("/events", http.HandlerFunc(headFunc)).Methods(http.MethodHead)
	mux.Handle("/events", http.HandlerFunc(eventHandler.Handle)).Methods(http.MethodPost)
	log.Printf("starting server")
	http.ListenAndServe(":8080", mux)
}

```

The inspiration for this model comes from a Slack RTM bot I've used in the past as well as the nlopes/slack RTM approach.
The idea is that you register handlers for each type of event you'll see (at this time only Message and Sync events are supported).
As Mandrill has two distinct types of wrapping events:

- [message](https://mandrill.zendesk.com/hc/en-us/articles/205583307-Message-Event-Webhook-format)
- [sync](https://mandrill.zendesk.com/hc/en-us/articles/205583297-Sync-Event-Webhook-format)

Sync Events come in two types:

- `blacklist`
- `whitelist`

with each have an associated `action`

Message events have a sub type associated with them (i.e. `send`, `hard_bounce`, etc)
You will have to do a type assertion on each type's `Data` field to get the concrete event and access its fields.

There are three buckets (currently) you can register your own code as a handler:

- `AddMessageHandler`: all `message` events are sent to handlers registered here
- `AddSyncHandler`: all `sync` events are sent to handlers registered here
- `AddUnknownHandler`: all failures in event parsing are sent to handlers registered here

(eventually I may allow for registering handlers for each specific event type)

Note that as these are handlers are called as part of an http request, you will get the original request context as well as the event.
This means if you need any data in the context (say, I dunno, for distributed tracing?) you'll have it in your function.

### Doing it all yourself

You can also do all the work yourself.
There are many functions for parsing out the incoming request from Mandrill, including validating the signature and extracting the JSON from the `mandrill_events` post.
These same functions are called by the previously mentioned handler approach.
Going this route may be useful if you don't want to handle any and all events at one endpoint.

If you go this route, you'll want to remember that you have to:

- Parse the form data with the key `mandrill_events` and unescape it (you can call `ParseRequest` for this) which will return a slice of `WebhookEvent`
- call `ParseInnerEvent` for each `WebhookEvent` to get either a `SyncEvent` or a `MessageEvent`
- type convert the `Data` field on each event to the concrete type of event

## Pending changes

I'm still working through the ergonomics of this package. I don't expect too much to change in terms of function signatures but I'm only working with data from Mandrill that is test data sent when testing your webhook.
