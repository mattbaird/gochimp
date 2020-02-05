package main

import (
	"fmt"
	"net/http"
	"net/url"

	"log"

	"github.com/lusis/gochimp/mandrill/events"
)

func main() {
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodHead {
			// Mandrill sends a HEAD request when creating a webhook
			// if that doesn't work, it sends an empty mandrill_events array
			w.WriteHeader(http.StatusOK)
			return
		}
		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		msg := r.Form.Get("mandrill_events")

		data, err := url.QueryUnescape(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(data) == 2 {
			log.Printf("empty event body")
			w.WriteHeader(http.StatusOK)
			return
		}
		postedEvents, err := events.ParseEvents([]byte(data))
		if err != nil {
			log.Printf("error parsing event: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, entry := range postedEvents {
			/*
				switch event := entry.(type) {
				case events.ClickEvent:
					log.Printf("click event: %v", event.Event)
				case events.SendEvent:
					log.Printf("send event: %v", event.Event)
				default:
					log.Printf("unknown event: %v", event)
				}
			*/
			log.Printf("saw event type %s with inner type %s", entry.Type, entry.InnerEventType)
			evt, err := events.ParseInnerEvent(entry)
			if err != nil {
				log.Printf("unable to parse inner event: %s", err.Error())
			} else {
				log.Printf("got inner event: %v", evt)
			}

		}
		w.WriteHeader(http.StatusOK)
	})
	fmt.Println("starting server")
	http.ListenAndServe(":8080", nil)
}
