package main

import (
	"fmt"

	"github.com/mattbaird/gochimp"
)

func main() {
	// Set the mailchimp key
	chimpKey := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-us99"

	chimpApi := gochimp.NewChimp(chimpKey, true)

	// Sample to get the lists
	chimpLists, err := chimpApi.ListsList(gochimp.ListsList{})
	if err != nil {
		fmt.Println("Error getting the lists. ", err)
	}

	for _, chimpList := range chimpLists.Data {
		fmt.Println("List: ", chimpList)
	}

	// Sample to subscribe

	// Set the mailchimp list id
	listId := "0a0a0a0a0a"

	chimpEmail := gochimp.Email{Email: "sample@email.com"}

	subscriber := gochimp.ListsSubscribe{ListId: listId, Email: chimpEmail, DoubleOptIn: false}

	chimpSubscribed, err := chimpApi.ListsSubscribe(subscriber)
	if err != nil {
		fmt.Println("Error subscribing: ", err)
	}

	fmt.Println("Subscribed Email: ", chimpSubscribed)
}
