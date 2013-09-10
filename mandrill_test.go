// Copyright 2012 Matthew Baird
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gochimp

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var mandrill, err = NewMandrill(os.Getenv("MANDRILL_KEY"))
var user string = os.Getenv("MANDRILL_USER")

func TestPing(t *testing.T) {
	response, err := mandrill.Ping()
	if response != "PONG!" {
		t.Error(fmt.Sprintf("failed to return PONG!, returned [%s]", response), err)
	}
}

func TestUserInfo(t *testing.T) {
	response, err := mandrill.UserInfo()
	if err != nil {
		t.Error("Error:", err)
	}
	if response.Username != user {
		t.Error("wrong user")
	}
}

func TestUserSenders(t *testing.T) {
	response, err := mandrill.UserSenders()
	if response == nil {
		t.Error("response was nil", err)
	}
	if err != nil {
		t.Error("Error:", err)
	}
}

func TestMessageSending(t *testing.T) {
	var message Message = Message{Html: "<b>hi there</b>", Text: "hello text", Subject: "Test Mail", FromEmail: user,
		FromName: user}
	message.AddRecipients(Recipient{Email: user, Name: user})
	response, err := mandrill.MessageSend(message, false)
	if err != nil {
		t.Error("Error:", err)
	}
	if response[0].Email != user {
		t.Errorf("Wrong email recipient, expecting %s, got %s", user, response[0].Email)
	}
}

const testTemplateName string = "test_transactional_template"

func TestTemplateAdd(t *testing.T) {
	// delete the test template if it exists already
	mandrill.TemplateDelete(testTemplateName)
	template, err := mandrill.TemplateAdd(testTemplateName, readTemplate("templates/transactional_basic.html"), true)
	if err != nil {
		t.Error("Error:", err)
	}
	if template.Name != "test_transactional_template" {
		t.Errorf("Wrong template name, expecting %s, got %s", testTemplateName, template.Name)
	}
	// try recreating, should error out
	template, err = mandrill.TemplateAdd(testTemplateName, readTemplate("templates/transactional_basic.html"), true)
	if err == nil {
		t.Error("Should have error'd on duplicate template")
	}
}

func TestTemplateList(t *testing.T) {
	_, err := mandrill.TemplateAdd("listTest", "testing 123", true)
	if err != nil {
		t.Error("Error:", err)
	}
	templates, err := mandrill.TemplateList()
	if err != nil {
		t.Error("Error:", err)
	}
	if len(templates) <= 0 {
		t.Errorf("Should have retrieved templates")
	}
	mandrill.TemplateDelete("listTest")
}

func TestTemplateInfo(t *testing.T) {
	template, err := mandrill.TemplateInfo(testTemplateName)
	if err != nil {
		t.Error("Error:", err)
	}
	if template.Name != "test_transactional_template" {
		t.Errorf("Wrong template name, expecting %s, got %s", testTemplateName, template.Name)
	}
}

func TestTemplateUpdate(t *testing.T) {
	// add a simple template
	template, err := mandrill.TemplateAdd("updateTest", "testing 123", true)
	template, err = mandrill.TemplateUpdate("updateTest", "testing 321", true)
	if err != nil {
		t.Error("Error:", err)
	}
	if template.Name != "updateTest" {
		t.Errorf("Wrong template name, expecting %s, got %s", "updateTest", template.Name)
	}
	if template.Code != "testing 321" {
		t.Errorf("Wrong template code, expecting %s, got %s", "testing 321", template.Code)
	}
	// be nice and tear down after test
	mandrill.TemplateDelete("updateTest")
}

func TestTemplatePublish(t *testing.T) {
	mandrill.TemplateDelete("publishTest")
	// add a simple template
	template, err := mandrill.TemplateAdd("publishTest", "testing 123", false)
	if err != nil {
		t.Error("Error:", err)
	}
	if template.Name != "publishTest" {
		t.Errorf("Wrong template name, expecting %s, got %s", testTemplateName, template.Name)
	}
	if template.PublishCode != "" {
		t.Errorf("Template should not have a publish code, got %s", template.PublishCode)
	}
	template, err = mandrill.TemplatePublish("publishTest")
	if err != nil {
		t.Error("Error:", err)
	}
	if template.Name != "publishTest" {
		t.Errorf("Wrong template name, expecting %s, got %s", testTemplateName, template.Name)
	}
	if template.PublishCode == "" {
		t.Errorf("Template should have a publish code, got %s", template.PublishCode)
	}
	mandrill.TemplateDelete("publishTest")
}

func TestTemplateRender(t *testing.T) {
	//make sure it's freshly added
	mandrill.TemplateDelete("renderTest")
	mandrill.TemplateAdd("renderTest", "*|MC:SUBJECT|*", true)
	//weak - should check results
	mergeVars := []Var{*NewVar("SUBJECT", "Hello, welcome")}
	result, err := mandrill.TemplateRender("renderTest", nil, mergeVars)
	if err != nil {
		t.Error("Error:", err)
	}
	// mandrill adds DOCTYPE
	if result != "<!DOCTYPE html>\nHello, welcome" {
		t.Errorf("Rendered Result incorrect, expecting %s, got %s", "<!DOCTYPE html>Hello, welcome", result)
	}
}

func TestTemplateRender2(t *testing.T) {
	//make sure it's freshly added
	mandrill.TemplateDelete("renderTest")
	mandrill.TemplateAdd("renderTest", "<div mc:edit=\"std_content00\"></div>", true)
	//weak - should check results
	templateContent := []Var{*NewVar("std_content00", "Hello, welcome")}
	result, err := mandrill.TemplateRender("renderTest", templateContent, nil)
	if err != nil {
		t.Error("Error:", err)
	}
	// mandrill adds DOCTYPE
	if result != "<!DOCTYPE html>\n<div>Hello, welcome</div>" {
		t.Errorf("Rendered Result incorrect, expecting %s, got %s", "<!DOCTYPE html>\n<div>Hello, welcome</div>", result)
	}
}

func TestMessageTemplateSend(t *testing.T) {
	//make sure it's freshly added
	mandrill.TemplateDelete(testTemplateName)
	mandrill.TemplateAdd(testTemplateName, readTemplate("templates/transactional_basic.html"), true)
	//weak - should check results
	templateContent := []Var{*NewVar("std_content00", "Hello, welcome")}
	mergeVars := []Var{*NewVar("SUBJECT", "Hello, welcome")}
	var message Message = Message{Subject: "Test Template Mail", FromEmail: user,
		FromName: user, GlobalMergeVars: mergeVars}
	message.AddRecipients(Recipient{Email: user, Name: user})
	_, err := mandrill.MessageSendTemplate(testTemplateName, templateContent, message, true)
	if err != nil {
		t.Error("Error:", err)
	}
	//todo - how do we test this better?
}

func readTemplate(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// senders tests

func TestSendersList(t *testing.T) {
	//make sure it's freshly added
	results, err := mandrill.SenderList()
	if err != nil {
		t.Error("Error:", err)
	}
	var foundUser = false
	for i := range results {
		var info Sender = results[i]
		if info.Address == user {
			foundUser = true
		}
	}
	if !foundUser {
		t.Errorf("should have found User %s in [%s] length array", user, len(results))
	}
}
