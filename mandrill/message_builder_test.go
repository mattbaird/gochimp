package mandrill

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMessageBuilder(t *testing.T) {
	mb := NewMessageBuilder("test@test.com", "Test Sender").
		AddRecipient(Recipient{
			Email: "recipient@test.com",
			Name:  "Test Recipient",
			Type:  "cc",
			MergeVars: []Var{
				{
					Name:    "foovar",
					Content: "foocontent",
				},
			},
			MetaData: map[string]string{
				"fookey": "foovar",
			},
		}).
		WithSubject("test subject").
		WithText("text body").
		WithHTML("<p>html body</p>").
		WithHeaders(map[string]string{"fooheader": "fooheaderval"}).
		WithTemplate("footemplate", []TemplateVar{{Name: "templatevar", Content: "templatecontent"}}).
		MergeAs("handlebars")
	mb.finalize()
	require.Len(t, mb.message.To, 1)
	require.Equal(t, "recipient@test.com", mb.message.To[0].Email)
	require.Equal(t, "Test Recipient", mb.message.To[0].Name)
	require.Equal(t, "cc", mb.message.To[0].Type)
	require.Len(t, mb.message.MergeVars, 1)
	require.Equal(t, "recipient@test.com", mb.message.MergeVars[0].Rcpt)
	require.Len(t, mb.message.MergeVars[0].Vars, 1)
	require.Equal(t, "foovar", mb.message.MergeVars[0].Vars[0].Name)
	require.Equal(t, "foocontent", mb.message.MergeVars[0].Vars[0].Content)
	require.Len(t, mb.message.RecipientMetaData, 1)
	require.Equal(t, "recipient@test.com", mb.message.RecipientMetaData[0].Rcpt)
	require.Contains(t, mb.message.RecipientMetaData[0].Values, "fookey")
	require.Equal(t, "foovar", mb.message.RecipientMetaData[0].Values["fookey"])
	require.Len(t, mb.message.Headers, 1)
	require.Equal(t, "fooheaderval", mb.message.Headers["fooheader"])
	require.True(t, mb.isTemplate)
	require.True(t, mb.message.Merge)
	require.Equal(t, "handlebars", mb.message.MergeLanguage)
}

func TestMessageBuilderSendTemplate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `[{
			"email":"test@test.com",
			"status":"delivered",
			"_id":"123456"
			}]`)
	}))
	defer srv.Close()
	err := Connect("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	resp, err := NewMessageBuilder("test@test.com", "Test Sender").
		AddRecipient(Recipient{
			Email: "recipient@test.com",
			Name:  "Test Recipient",
			Type:  "cc",
			MergeVars: []Var{
				{
					Name:    "foovar",
					Content: "foocontent",
				},
			},
			MetaData: map[string]string{
				"fookey": "foovar",
			},
		}).
		WithSubject("test subject").
		WithText("text body").
		WithHTML("<p>html body</p>").
		WithHeaders(map[string]string{"fooheader": "fooheaderval"}).
		WithTemplate("footemplate", []TemplateVar{{Name: "templatevar", Content: "templatecontent"}}).
		MergeAs("handlebars").
		Send()
	require.NoError(t, err)
	require.Equal(t, "delivered", resp[0].Status)
	require.Equal(t, "test@test.com", resp[0].Email)
	require.Empty(t, resp[0].RejectReason)
	require.Equal(t, "123456", resp[0].ID)
}

func TestMessageBuilderSend(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `[{
			"email":"test@test.com",
			"status":"delivered",
			"_id":"123456"
			}]`)
	}))
	defer srv.Close()
	err := Connect("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	resp, err := NewMessageBuilder("test@test.com", "Test Sender").
		AddRecipient(Recipient{
			Email: "recipient@test.com",
			Name:  "Test Recipient",
			Type:  "cc",
			MergeVars: []Var{
				{
					Name:    "foovar",
					Content: "foocontent",
				},
			},
			MetaData: map[string]string{
				"fookey": "foovar",
			},
		}).
		WithSubject("test subject").
		WithText("text body").
		WithHTML("<p>html body</p>").
		WithHeaders(map[string]string{"fooheader": "fooheaderval"}).
		MergeAs("handlebars").
		Send()
	require.NoError(t, err)
	require.Equal(t, "delivered", resp[0].Status)
	require.Equal(t, "test@test.com", resp[0].Email)
	require.Empty(t, resp[0].RejectReason)
	require.Equal(t, "123456", resp[0].ID)
}
