gochimp
=======

[![GoDoc](https://godoc.org/github.com/mattbaird/gochimp?status.svg)](https://godoc.org/github.com/mattbaird/gochimp)

Go based API for Mailchimp, starting with Mandrill.

To run tests, set a couple env variables.
(replacing values with your own mandrill credentials):
```bash
$ export MANDRILL_KEY=111111111-1111-1111-1111-111111111
$ export MANDRILL_USER=user@domain.com
```

Mandrill Status
===============
* API Feature complete on Oct 26/2012
* Adding tests, making naming conventions consistent, and refactoring error handling

Chimp Status
============
* Not started

Getting Started
===============
Below is an example approach to rendering custom content into a Mandrill
template called "welcome email" and sending the rendered email.

```
package main

import (
	"fmt"
	"github.com/mattbaird/gochimp"
	"os"
)

func main() {
	apiKey := os.Getenv("MANDRILL_KEY")
	mandrillApi, err := gochimp.NewMandrill(apiKey)

	if err != nil {
		fmt.Println("Error instantiating client")
	}

	templateName := "welcome email"
	contentVar := gochimp.Var{"main", "<h1>Welcome aboard!</h1>"}
	content := []gochimp.Var{contentVar}

	_, err = mandrillApi.TemplateAdd(templateName, fmt.Sprintf("%s", contentVar.Content), true)
	if err != nil {
		fmt.Println("Error adding template: %v", err)
		return
	}
	defer mandrillApi.TemplateDelete(templateName)
	renderedTemplate, err := mandrillApi.TemplateRender(templateName, content, nil)

	if err != nil {
		fmt.Println("Error rendering template: %v", err)
		return
	}

	recipients := []gochimp.Recipient{
		gochimp.Recipient{Email: "person@place.com"},
	}

	message := gochimp.Message{
		Html:      renderedTemplate,
		Subject:   "Welcome aboard!",
		FromEmail: "person@place.com",
		FromName:  "Boss Man",
		To:        recipients,
	}

	_, err = mandrillApi.MessageSend(message, false)

	if err != nil {
		fmt.Println("Error sending message")
	}
}

```
