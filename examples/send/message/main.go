package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/lusis/gochimp/mandrill"
)

func main() {
	if len(os.Getenv("MANDRILL_FROM")) == 0 ||
		len(os.Getenv("MANDRILL_TO")) == 0 ||
		len(os.Getenv("MANDRILL_KEY")) == 0 {
		log.Fatal("please set the environment variables MANDRILL_FROM, MANDRILL_TO to appropriate email addresses and set MANDRILL_KEY to your api key")
		os.Exit(1)
	}
	client, err := mandrill.New(os.Getenv("MANDRILL_KEY"), mandrill.WithDebug(), mandrill.WithPing())
	if err != nil {
		log.Fatalf("error creating client: %s", err.Error())
	}
	msg := mandrill.Message{
		Subject:   "Email subject",
		Text:      "welcome api",
		FromEmail: os.Getenv("MANDRILL_FROM"),
		FromName:  "Mandrill Sender",
		To: []mandrill.To{
			mandrill.To{
				Email: os.Getenv("MANDRILL_TO"),
				Name:  "Mandrill Recipient",
			},
		},
	}

	resp, err := client.SendMessage(msg)
	if err != nil {
		log.Fatalf("error sending message: %s", err.Error())
	}
	spew.Dump(resp)
}
