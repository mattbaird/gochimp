package mandrill

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTagsList(t *testing.T) {
	resp := `
	[
    {
        "tag": "example-tag",
        "reputation": 42,
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
	tags, err := client.ListTags()
	require.NoError(t, err)
	require.Len(t, tags, 1)
}

func TestTagDelete(t *testing.T) {
	resp := `
    {
        "tag": "example-tag",
        "reputation": 42,
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
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	tag := &Tag{}
	err = tag.Delete()
	require.NoError(t, err)
}

func TestTagInfo(t *testing.T) {
	resp := `
    {
        "tag": "example-tag",
        "reputation": 42,
        "sent": 42,
        "hard_bounces": 42,
        "soft_bounces": 42,
        "rejects": 42,
        "complaints": 42,
        "unsubs": 42,
        "opens": 42,
        "clicks": 42,
        "unique_opens": 42,
		"unique_clicks": 42,
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
	tag, err := client.TagInfo("foo")
	require.NoError(t, err)
	require.NotNil(t, tag)
	require.Len(t, tag.Stats, 5)
}
