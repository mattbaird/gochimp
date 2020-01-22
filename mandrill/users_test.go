package mandrill

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserInfo(t *testing.T) {
	respBody := `
	{
		"username": "myusername",
		"created_at": "2013-01-01 15:30:27",
		"public_id": "aaabbbccc112233",
		"reputation": 42,
		"hourly_quota": 42,
		"backlog": 42,
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
			},
			"all_time": {
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
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	i, err := client.UserInfo()
	require.NoError(t, err)
	require.NotNil(t, i)
	require.Len(t, i.Stats, 6)
	require.NotEmpty(t, i.Stats["all_time"])
	require.Equal(t, i.Stats["all_time"].Sent, int32(42))
}

func TestUsersSenders(t *testing.T) {
	respBody := `
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
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	i, err := client.UserSenders()
	require.NoError(t, err)
	require.NotNil(t, i)
	require.Len(t, i, 1)
}
