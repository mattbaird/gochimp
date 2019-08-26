// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
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

	"github.com/stretchr/testify/require"
)

func testMandrillClient(t *testing.T) *MandrillAPI {
	if os.Getenv("MANDRILL_KEY") == "" {
		t.Skip()
	}
	mandrill, err := NewMandrill(os.Getenv("MANDRILL_KEY"))
	require.NoError(t, err)
	require.NotNil(t, mandrill)
	return mandrill
}

var testUser string = os.Getenv("MANDRILL_USER")

const testTemplateName string = "test_transactional_template"

func TestPing(t *testing.T) {
	mandrill := testMandrillClient(t)
	response, err := mandrill.Ping()
	require.NoError(t, err)
	require.Equal(t, "PONG!", response)
}

func TestUserInfo(t *testing.T) {
	mandrill := testMandrillClient(t)
	response, err := mandrill.UserInfo()
	require.NoError(t, err)
	require.Equal(t, testUser, response.Username)
}

func TestUserSenders(t *testing.T) {
	mandrill := testMandrillClient(t)
	response, err := mandrill.UserSenders()
	require.NoError(t, err)
	require.NotNil(t, response)

}

func TestMessageSending(t *testing.T) {
	mandrill := testMandrillClient(t)

	require.NotEmpty(t, testUser)
	var message Message = Message{Html: "<b>hi there</b>", Text: "hello text", Subject: "Test Mail", FromEmail: testUser,
		FromName: testUser}
	message.AddRecipients(Recipient{Email: testUser, Name: testUser, Type: "to"})
	message.AddRecipients(Recipient{Email: testUser, Name: testUser, Type: "cc"})
	message.AddRecipients(Recipient{Email: testUser, Name: testUser, Type: "bcc"})
	response, err := mandrill.MessageSend(message, false)
	require.NoError(t, err)
	require.NotNil(t, response)
	require.Len(t, response, 3)
	require.Equal(t, testUser, response[0])
	require.Equal(t, testUser, response[1])
	require.Equal(t, testUser, response[2])
}

func TestTemplateAdd(t *testing.T) {
	mandrill := testMandrillClient(t)

	// delete the test template if it exists already
	_, _ = mandrill.TemplateDelete(testTemplateName)
	template, err := mandrill.TemplateAdd(testTemplateName, readTemplate("templates/transactional_basic.html"), true)
	require.NoError(t, err)
	require.Equal(t, testTemplateName, template.Name)

	// try recreating, should error out
	template, err = mandrill.TemplateAdd(testTemplateName, readTemplate("templates/transactional_basic.html"), true)
	require.Error(t, err)
}

func TestTemplateList(t *testing.T) {
	mandrill := testMandrillClient(t)

	_, err := mandrill.TemplateAdd("listTest", "testing 123", true)
	require.NoError(t, err)

	templates, err := mandrill.TemplateList()
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(templates), 1)
	_, err = mandrill.TemplateDelete("listTest")
	require.NoError(t, err)
}

func TestTemplateInfo(t *testing.T) {
	mandrill := testMandrillClient(t)

	template, err := mandrill.TemplateInfo(testTemplateName)
	require.NoError(t, err)
	require.Equal(t, "test_transactional_template", template.Name)
}

func TestTemplateUpdate(t *testing.T) {
	mandrill := testMandrillClient(t)

	// add a simple template
	template, err := mandrill.TemplateAdd("updateTest", "testing 123", true)
	require.NoError(t, err)
	require.NotNil(t, template)
	template, err = mandrill.TemplateUpdate("updateTest", "testing 321", true)
	require.NoError(t, err)
	require.Equal(t, "updateTest", template.Name)
	require.Equal(t, "testing 321", template.Code)
	// be nice and tear down after test
	_, err = mandrill.TemplateDelete("updateTest")
	require.NoError(t, err)
}

func TestTemplatePublish(t *testing.T) {
	mandrill := testMandrillClient(t)

	_, _ = mandrill.TemplateDelete("publishTest")
	// add a simple template
	template, err := mandrill.TemplateAdd("publishTest", "testing 123", false)
	require.NoError(t, err)
	require.Equal(t, "publishTest", template.Name)
	require.Empty(t, template.PublishCode)

	template, err = mandrill.TemplatePublish("publishTest")
	require.NoError(t, err)
	require.Equal(t, "publishTest", template.Name)
	require.NotEmpty(t, template.PublishCode)

	_, err = mandrill.TemplateDelete("publishTest")
	require.NoError(t, err)
}

func TestTemplateRender(t *testing.T) {
	mandrill := testMandrillClient(t)

	//make sure it's freshly added
	_, _ = mandrill.TemplateDelete("renderTest")
	template, err := mandrill.TemplateAdd("renderTest", "*|MC:SUBJECT|*", true)
	require.NoError(t, err)
	require.NotNil(t, template)
	mergeVars := []Var{*NewVar("SUBJECT", "Hello, welcome")}
	result, err := mandrill.TemplateRender("renderTest", nil, mergeVars)
	require.NoError(t, err)
	require.Equal(t, "Hello, welcome", result)
}

func TestTemplateRender2(t *testing.T) {
	mandrill := testMandrillClient(t)

	_, _ = mandrill.TemplateDelete("renderTest")
	template, err := mandrill.TemplateAdd("renderTest", "<div mc:edit=\"std_content00\"></div>", true)
	require.NoError(t, err)
	require.NotNil(t, template)

	templateContent := []Var{*NewVar("std_content00", "Hello, welcome")}
	result, err := mandrill.TemplateRender("renderTest", templateContent, nil)
	require.NoError(t, err)
	require.Equal(t, "<div>Hello, welcome</div>", result)
}

func TestMessageTemplateSend(t *testing.T) {
	mandrill := testMandrillClient(t)

	//make sure it's freshly added
	_, _ = mandrill.TemplateDelete(testTemplateName)
	template, err := mandrill.TemplateAdd(testTemplateName, readTemplate("templates/transactional_basic.html"), true)
	require.NoError(t, err)
	require.NotNil(t, template)

	templateContent := []Var{*NewVar("std_content00", "Hello, welcome")}
	mergeVars := []Var{*NewVar("SUBJECT", "Hello, welcome")}
	var message Message = Message{Subject: "Test Template Mail", FromEmail: testUser,
		FromName: testUser, GlobalMergeVars: mergeVars}
	message.AddRecipients(Recipient{Email: testUser, Name: testUser})
	_, err = mandrill.MessageSendTemplate(testTemplateName, templateContent, message, true)
	require.NoError(t, err)
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
	mandrill := testMandrillClient(t)
	//make sure it's freshly added
	results, err := mandrill.SenderList()
	require.NoError(t, err)
	require.True(t, len(results) > 0)
	require.Contains(t, results, testUser)
}

// incoming tests

func TestInboundDomainListAddCheckDelete(t *testing.T) {
	mandrill := testMandrillClient(t)
	domainName := "improbable.example.com"
	domains, err := mandrill.InboundDomainList()
	require.NoError(t, err)

	originalCount := len(domains)
	domain, err := mandrill.InboundDomainAdd(domainName)
	require.NoError(t, err)

	domains, err = mandrill.InboundDomainList()
	require.NoError(t, err)

	newCount := len(domains)
	require.Equal(t, originalCount+1, newCount)

	newDomain, err := mandrill.InboundDomainCheck(domainName)
	require.NoError(t, err)
	require.Equal(t, domain.CreatedAt, newDomain.CreatedAt)
	require.Equal(t, domain.Domain, newDomain.Domain)
	require.Equal(t, domain.ValidMx, newDomain.ValidMx)

	_, err = mandrill.InboundDomainDelete(domainName)
	require.NoError(t, err)

	domains, err = mandrill.InboundDomainList()
	require.NoError(t, err)

	deletedCount := len(domains)
	require.Equal(t, originalCount, deletedCount)
}

func TestInboundDomainRoutesAndRaw(t *testing.T) {
	mandrill := testMandrillClient(t)
	domainName := "www.google.com"
	emailAddress := "test"
	webhookUrl := fmt.Sprintf("http://%v/", domainName)
	_, err := mandrill.InboundDomainAdd(domainName)
	require.NoError(t, err)

	routeList, err := mandrill.RouteList(domainName)
	require.NoError(t, err)

	count := len(routeList)
	require.Equal(t, 0, count)

	route, err := mandrill.RouteAdd(domainName, emailAddress, webhookUrl)
	require.NoError(t, err)
	require.Equal(t, emailAddress, route.Pattern)
	require.Equal(t, webhookUrl, route.Url)

	newDomainName := "www.google.com"
	newEmailAddress := "test2"
	newWebhookUrl := fmt.Sprintf("http://%v/", newDomainName)
	_, err = mandrill.InboundDomainCheck(newDomainName)
	require.NoError(t, err)

	route, err = mandrill.RouteUpdate(route.Id, newDomainName, newEmailAddress, newWebhookUrl)
	require.NoError(t, err)
	require.Equal(t, newEmailAddress, route.Pattern)
	require.Equal(t, newWebhookUrl, route.Url)

	route, err = mandrill.RouteDelete(route.Id)
	require.NoError(t, err)

	routeList, err = mandrill.RouteList(domainName)
	require.NoError(t, err)

	newCount := len(routeList)
	require.Equal(t, count, newCount)

	/* skipping this for now
	// not sure if this actually sends. Should come up with something better if so
	rawMessage := "From: sender@example.com\nTo: test2@www.google.com\nSubject: Some Subject\n\nSome content."
	_, err = mandrill.SendRawMIME(rawMessage, []string{"test2@www.google.com"}, "test@www.google.com", "", "127.0.0.1")
	require.NoError(t, err)
	*/

	_, err = mandrill.InboundDomainDelete(domainName)
	require.NoError(t, err)

}
