package mandrill

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebHookInfo(t *testing.T) {
	respBody := `
	{
		"id": 42,
		"url": "http://example/webhook-url",
		"description": "My Example Webhook",
		"auth_key": "gplJ8yWptFTqCoq5S1SHPA",
		"events": [
			"send",
			"open",
			"click"
		],
		"created_at": "2013-01-01 15:30:27",
		"last_sent_at": "2013-01-01 15:30:49",
		"batches_sent": 42,
		"events_sent": 42,
		"last_error": "example last_error"
	}
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	hook, err := client.WebhookInfo(42)
	require.NoError(t, err)
	require.NotNil(t, hook)
	require.Len(t, hook.Events, 3)
	require.Equal(t, int32(42), hook.BatchesSent)
	require.Equal(t, int32(42), hook.EventsSent)
}

func TestWebHookAdd(t *testing.T) {
	respBody := `
	{
		"id": 42,
		"url": "http://example/webhook-url",
		"description": "My Example Webhook",
		"auth_key": "XXXXXXX",
		"events": [
			"send",
			"open",
			"click"
		],
		"created_at": "2013-01-01 15:30:27",
		"last_sent_at": "2013-01-01 15:30:49",
		"batches_sent": 1,
		"events_sent": 1,
		"last_error": ""
	}
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	// We use different values here to ensure the values from the api call
	// are working (refreshed values)
	hook := &Webhook{
		URL:         "http://example",
		Description: "foo",
	}
	err = hook.Add()
	require.NoError(t, err)
	require.Len(t, hook.Events, 3)
	require.Equal(t, 42, hook.ID)
	require.Equal(t, int32(1), hook.BatchesSent)
	require.Equal(t, int32(1), hook.EventsSent)
	require.Equal(t, "http://example/webhook-url", hook.URL)
	require.Equal(t, "My Example Webhook", hook.Description)
	require.Equal(t, "XXXXXXX", hook.AuthKey)
	hook.URL = "changed value"
	err = hook.Info()
	require.NoError(t, err)
	require.Equal(t, "http://example/webhook-url", hook.URL)
}
