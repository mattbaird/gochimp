package main

import (
	"fmt"
	"os"

	"github.com/mattbaird/gochimp"
)

func main() {
	apiKey := os.Getenv("MANDRILL_KEY")
	mandrillApi, err := gochimp.NewMandrill(apiKey)

	if err != nil {
		fmt.Println("Error instantiating client")
	}

	templateName := "welcome email"
	contentVar := gochimp.Var{
		Name:    "main",
		Content: "<h1>Welcome aboard!</h1>",
	}
	content := []gochimp.Var{contentVar}

	_, err = mandrillApi.TemplateAdd(templateName, fmt.Sprintf("%s", contentVar.Content), true)
	if err != nil {
		fmt.Printf("Error adding template: %s\n", err.Error())
		return
	}
	defer func() {
		_, _ = mandrillApi.TemplateDelete(templateName)
	}()
	renderedTemplate, err := mandrillApi.TemplateRender(templateName, content, nil)

	if err != nil {
		fmt.Printf("Error adding template: %s\n", err.Error())
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
