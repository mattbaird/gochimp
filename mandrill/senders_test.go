package mandrill

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendersList(t *testing.T) {
	resp := `
	[
    {
        "address": "sender.example@mandrillapp.com",
        "created_at": "2013-01-01 15:30:27",
        "sent": 42,
        "hard_bounces": 42,
        "soft_bounces": 42,
        "rejects": 42,
        "complaints": 42,
        "unsubs": 42,
        "opens": 42,
        "clicks": 42,
        "unique_opens": 42,
        "unique_clicks": 42
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
	senders, err := client.ListSenders()
	require.NoError(t, err)
	require.NotNil(t, senders)
	require.Len(t, senders, 1)
}

func TestSendersDomains(t *testing.T) {
	resp := `
	[
		{
			"domain": "example.com",
			"created_at": "2013-01-01 15:30:27",
			"last_tested_at": "2013-01-01 15:40:42",
			"spf": {
				"valid": true,
				"valid_after": "2013-01-01 15:45:23",
				"error": "example error"
			},
			"dkim": {
				"valid": true,
				"valid_after": "2013-01-01 15:45:23",
				"error": "example error"
			},
			"verified_at": "2013-01-01 15:50:21",
			"valid_signing": true
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
	senders, err := client.SendersDomains()
	require.NoError(t, err)
	require.NotNil(t, senders)
	require.Len(t, senders, 1)
}

func TestSenderInfo(t *testing.T) {
	resp := `
	{
		"address": "sender.example@mandrillapp.com",
		"created_at": "2013-01-01 15:30:27",
		"sent": 42,
		"hard_bounces": 42,
		"soft_bounces": 42,
		"rejects": 42,
		"complaints": 42,
		"unsubs": 42,
		"opens": 42,
		"clicks": 42,
		"stats": {
			"today": {
				"sent": 42,
				"hard_bounces": 42,
				"soft_bounces": 42,
				"rejects": 42,
				"complaints": 42,
				"unsubs": 42,
				"opens": 42,
				"unique_opens": 42,
				"clicks": 42,
				"unique_clicks": 42
			},
			"last_7_days": {
				"sent": 42,
				"hard_bounces": 42,
				"soft_bounces": 42,
				"rejects": 42,
				"complaints": 42,
				"unsubs": 42,
				"opens": 42,
				"unique_opens": 42,
				"clicks": 42,
				"unique_clicks": 42
			},
			"last_30_days": {
				"sent": 42,
				"hard_bounces": 42,
				"soft_bounces": 42,
				"rejects": 42,
				"complaints": 42,
				"unsubs": 42,
				"opens": 42,
				"unique_opens": 42,
				"clicks": 42,
				"unique_clicks": 42
			},
			"last_60_days": {
				"sent": 42,
				"hard_bounces": 42,
				"soft_bounces": 42,
				"rejects": 42,
				"complaints": 42,
				"unsubs": 42,
				"opens": 42,
				"unique_opens": 42,
				"clicks": 42,
				"unique_clicks": 42
			},
			"last_90_days": {
				"sent": 42,
				"hard_bounces": 42,
				"soft_bounces": 42,
				"rejects": 42,
				"complaints": 42,
				"unsubs": 42,
				"opens": 42,
				"unique_opens": 42,
				"clicks": 42,
				"unique_clicks": 42
			}
		}
	}
	`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	sender, err := client.GetSenderInfo("foo")
	require.NoError(t, err)
	require.NotNil(t, sender)
	require.Len(t, sender.Stats, 5)
}

func TestSendersTimeSeries(t *testing.T) {
	resp := `
	[
		{
			"time": "2013-01-01 15:30:27",
			"sent": 42,
			"hard_bounces": 42,
			"soft_bounces": 42,
			"rejects": 42,
			"complaints": 42,
			"opens": 42,
			"unique_opens": 42,
			"clicks": 42,
			"unique_clicks": 42
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
	ts, err := client.GetSenderTimeSeries("foo")
	require.NoError(t, err)
	require.NotNil(t, ts)
	require.Len(t, ts, 1)
}

func TestAddSendingDomain(t *testing.T) {
	resp := `
	{
		"domain": "example.com",
		"created_at": "2013-01-01 15:30:27",
		"last_tested_at": "2013-01-01 15:40:42",
		"spf": {
			"valid": true,
			"valid_after": "2013-01-01 15:45:23",
			"error": "example error"
		},
		"dkim": {
			"valid": true,
			"valid_after": "2013-01-01 15:45:23",
			"error": "example error"
		},
		"verified_at": "2013-01-01 15:50:21",
		"valid_signing": true
	}
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	ts, err := client.AddSendingDomain("foo")
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestCheckSendingDomain(t *testing.T) {
	resp := `
	{
		"domain": "example.com",
		"created_at": "2013-01-01 15:30:27",
		"last_tested_at": "2013-01-01 15:40:42",
		"spf": {
			"valid": true,
			"valid_after": "2013-01-01 15:45:23",
			"error": "example error"
		},
		"dkim": {
			"valid": true,
			"valid_after": "2013-01-01 15:45:23",
			"error": "example error"
		},
		"verified_at": "2013-01-01 15:50:21",
		"valid_signing": true
	}
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	ts, err := client.CheckSendingDomain("foo")
	require.NoError(t, err)
	require.NotNil(t, ts)
}

func TestVerifySendingDomain(t *testing.T) {
	resp := `
	{
		"status": "example status",
		"domain": "example domain",
		"email": "example email"
	}
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	s, e, err := client.VerifySendingDomain("foo", "bar")
	require.NoError(t, err)
	require.NotEmpty(t, s)
	require.NotEmpty(t, e)
}
