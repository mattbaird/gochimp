package mandrill

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMessageSend(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `[{
			"email":"test@test.com",
			"status":"delivered",
			"_id":"123456"
			}]`)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	resp, err := client.SendMessage(Message{
		Text:      "welcome",
		Subject:   "hello",
		FromEmail: "test@test.com",
		FromName:  "admin@test.com",
		To: []To{
			{
				Email: "test@test.com",
				Type:  "to",
				Name:  "a user",
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, "delivered", resp[0].Status)
	require.Equal(t, "test@test.com", resp[0].Email)
	require.Empty(t, resp[0].RejectReason)
	require.Equal(t, "123456", resp[0].ID)
}

func TestMessageSearch(t *testing.T) {
	resp := `
	[
    {
        "ts": 1365190000,
        "_id": "abc123abc123abc123abc123",
        "sender": "sender@example.com",
        "template": "example-template",
        "subject": "example subject",
        "email": "recipient.email@example.com",
        "tags": [
            "password-reset"
        ],
        "opens": 42,
        "opens_detail": [
            {
                "ts": 1365190001,
                "ip": "55.55.55.55",
                "location": "Georgia, US",
                "ua": "Linux/Ubuntu/Chrome/Chrome 28.0.1500.53"
            }
        ],
        "clicks": 42,
        "clicks_detail": [
            {
                "ts": 1365190001,
                "url": "http://www.example.com",
                "ip": "55.55.55.55",
                "location": "Georgia, US",
                "ua": "Linux/Ubuntu/Chrome/Chrome 28.0.1500.53"
            }
        ],
        "state": "sent",
        "metadata": {
            "user_id": "123",
            "website": "www.example.com"
        }
    }
	]
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	res, err := client.SearchMessages(MessageSearchParams{})
	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Len(t, res[0].ClicksDetails, 1)
	require.Len(t, res[0].OpensDetails, 1)
}

func TestMessageSearchParams(t *testing.T) {
	resp := `
	[
    {
        "ts": 1365190000,
        "_id": "abc123abc123abc123abc123",
        "sender": "sender@example.com",
        "template": "example-template",
        "subject": "example subject",
        "email": "recipient.email@example.com",
        "tags": [
            "password-reset"
        ],
        "opens": 42,
        "opens_detail": [
            {
                "ts": 1365190001,
                "ip": "55.55.55.55",
                "location": "Georgia, US",
                "ua": "Linux/Ubuntu/Chrome/Chrome 28.0.1500.53"
            }
        ],
        "clicks": 42,
        "clicks_detail": [
            {
                "ts": 1365190001,
                "url": "http://www.example.com",
                "ip": "55.55.55.55",
                "location": "Georgia, US",
                "ua": "Linux/Ubuntu/Chrome/Chrome 28.0.1500.53"
            }
        ],
        "state": "sent",
        "metadata": {
            "user_id": "123",
            "website": "www.example.com"
		},
		"smtp_events": [
			{
				"ts": 1365190001,
				"type": "sent",
				"diag": "250 OK"
			}
		]
    }
	]
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	now := time.Now()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	res, err := client.SearchMessages(MessageSearchParams{DateFrom: &now, DateTo: &now})
	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Len(t, res[0].ClicksDetails, 1)
	require.Len(t, res[0].OpensDetails, 1)
	require.Len(t, res[0].SMTPEvents, 1)
}

func TestMessageInfo(t *testing.T) {
	resp := `
	{
		"ts": 1365190000,
		"_id": "abc123abc123abc123abc123",
		"sender": "sender@example.com",
		"template": "example-template",
		"subject": "example subject",
		"email": "recipient.email@example.com",
		"tags": [
			"password-reset"
		],
		"opens": 42,
		"opens_detail": [
			{
				"ts": 1365190001,
				"ip": "55.55.55.55",
				"location": "Georgia, US",
				"ua": "Linux/Ubuntu/Chrome/Chrome 28.0.1500.53"
			}
		],
		"clicks": 42,
		"clicks_detail": [
			{
				"ts": 1365190001,
				"url": "http://www.example.com",
				"ip": "55.55.55.55",
				"location": "Georgia, US",
				"ua": "Linux/Ubuntu/Chrome/Chrome 28.0.1500.53"
			}
		],
		"state": "sent",
		"metadata": {
			"user_id": "123",
			"website": "www.example.com"
		},
		"smtp_events": [
			{
				"ts": 1365190001,
				"type": "sent",
				"diag": "250 OK"
			}
		]
	}
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	res, err := client.MessageInfo("12345")
	require.NoError(t, err)
	require.Len(t, res.ClicksDetails, 1)
	require.Len(t, res.OpensDetails, 1)
	require.Len(t, res.SMTPEvents, 1)
}

func TestMessageContent(t *testing.T) {
	resp := `
	{
		"subject": "Some Subject",
		"from_email": "sender@example.com",
		"from_name": "Sender Name",
		"to": [
			{
				"email": "recipient.email@example.com",
				"name": "Recipient Name"
			}
		],
		"headers": {
			"Reply-To": "replies@example.com"
		},
		"text": "Some text content",
		"html": "<p>Some HTML content</p>",
		"attachments": [
			{
				"name": "example.txt",
				"type": "text/plain",
				"binary": false,
				"content": "example non-binary content"
			}
		],
		"images": [
			{
				"name": "IMAGEID",
				"type": "image/png",
				"content": "ZXhhbXBsZSBmaWxl"
			}
		]
	}
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	res, err := client.MessageContent("123456")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Len(t, res.To, 1)
	require.Len(t, res.Attachments, 1)
	require.Len(t, res.Images, 1)
}
