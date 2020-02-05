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
