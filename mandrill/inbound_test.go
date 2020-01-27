package mandrill

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIPList(t *testing.T) {
	respBody := `[
		{
			"ip": "127.0.0.1",
			"created_at": "2013-01-01 15:50:21",
			"pool": "Main Pool",
			"domain": "mail1.example.mandrillapp.com",
			"custom_dns": {
				"enabled": true,
				"valid": true,
				"error": "example error"
			},
			"warmup": {
				"warming_up": true,
				"start_at": "2013-03-01 12:00:01",
				"end_at": "2013-03-31 12:00:01"
			}
		}
	]`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	md, err := client.IPList()
	require.NoError(t, err)
	require.Len(t, md, 1)
}

func TestIPInfo(t *testing.T) {
	respBody := `{
		"ip": "127.0.0.1",
		"created_at": "2013-01-01 15:50:21",
		"pool": "Main Pool",
		"domain": "mail1.example.mandrillapp.com",
		"custom_dns": {
			"enabled": true,
			"valid": true,
			"error": "example error"
		},
		"warmup": {
			"warming_up": true,
			"start_at": "2013-03-01 12:00:01",
			"end_at": "2013-03-31 12:00:01"
		}
	}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	//layout := `2006-01-02 15:04:05`
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	md, err := client.IPInfo("127.0.0.1")
	require.NoError(t, err)
	require.Equal(t, "127.0.0.1", md.IP)
	require.Equal(t, "Main Pool", md.Pool)
	require.True(t, md.CustomDNS.Enabled)
	require.True(t, md.CustomDNS.Valid)
	require.True(t, md.WarmUp.WarmingUp)
	require.Equal(t, "2013-01-01 15:50:21 +0000 UTC", fmt.Sprintf(md.CreatedAt.String()))
}

func TestIPProvision(t *testing.T) {
	respBody := `{
		"requested_at": "2013-01-01 01:52:21"
	}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	//layout := `2006-01-02 15:04:05`
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	md, err := client.IPProvision("foo", true)
	require.NoError(t, err)
	require.Equal(t, "2013-01-01 01:52:21 +0000 UTC", fmt.Sprintf(md.String()))
}

func TestIPStartWarmUp(t *testing.T) {
	respBody := `{
		"ip": "127.0.0.1",
		"created_at": "2013-01-01 15:50:21",
		"pool": "Main Pool",
		"domain": "mail1.example.mandrillapp.com",
		"custom_dns": {
			"enabled": true,
			"valid": true,
			"error": "example error"
		},
		"warmup": {
			"warming_up": true,
			"start_at": "2013-03-01 12:00:01",
			"end_at": "2013-03-31 12:00:01"
		}
	}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	//layout := `2006-01-02 15:04:05`
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	md, err := client.IPStartWarmUp("127.0.0.1")
	require.NoError(t, err)
	require.Equal(t, "127.0.0.1", md.IP)
	require.Equal(t, "Main Pool", md.Pool)
	require.True(t, md.CustomDNS.Enabled)
	require.True(t, md.CustomDNS.Valid)
	require.True(t, md.WarmUp.WarmingUp)
	require.Equal(t, "2013-01-01 15:50:21 +0000 UTC", fmt.Sprintf(md.CreatedAt.String()))
}
