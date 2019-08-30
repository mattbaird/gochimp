package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/lusis/gochimp/mandrill"
)

func main() {
	client, err := mandrill.New(os.Getenv("MANDRILL_KEY"), mandrill.WithDebug(), mandrill.WithPing())
	if err != nil {
		log.Fatalf("error creating client: %s", err.Error())
	}
	u, err := client.ListSenders()
	if err != nil {
		log.Fatalf("error getting info: %s", err.Error())
	}
	spew.Dump(u)
}
