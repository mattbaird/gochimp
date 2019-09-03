package mandrill

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lusis/gochimp/mandrill/api"
	"github.com/stretchr/testify/require"
)

func testFailOption() ClientOption {
	return func(c *Client) error {
		return errors.New("failed option")
	}
}
func TestClientOptions(t *testing.T) {
	clientOpts := []ClientOption{
		WithDebug(),
		WithSubAccount("12345"),
		WithEndpoint("http://localhost:12345"),
	}
	c, err := New("abcdefg", clientOpts...)
	require.NoError(t, err)
	require.NotNil(t, c)
	require.True(t, c.debug)
	require.Equal(t, "12345", c.subaccount)
	require.Equal(t, "http://localhost:12345", c.endpoint)
	require.Equal(t, "abcdefg", c.apiKey)
}

func TestClientOptionFail(t *testing.T) {
	c, err := New("abcdefg", testFailOption())
	require.Error(t, err)
	require.Nil(t, c)
}

func TestConnect(t *testing.T) {
	err := Connect("12345")
	require.NoError(t, err)
	require.NotNil(t, globalClient)
	require.Equal(t, "https://mandrillapp.com/api/1.0", globalClient.endpoint)
	require.NotNil(t, globalClient.logger)
	require.Equal(t, "12345", globalClient.apiKey)
	require.NotNil(t, globalClient.httpClient)
}

func TestDoPing(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"PING":"PONG!"}`)
	}))
	defer srv.Close()
	err := Connect("abcdefg", WithPing(), WithEndpoint(srv.URL))
	require.NoError(t, err)
}

func TestDoPingFail(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		fmt.Fprintln(w, `{
			"status": "error",
			"code": -1,
			"name": "Invalid_Key",
			"message": "Invalid API key"
		}`)

	}))
	defer srv.Close()
	err := Connect("abcdefg", WithPing(), WithEndpoint(srv.URL))
	require.Error(t, err)
	require.IsType(t, &InvalidKeyError{}, err)
}

func TestWithLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "TEST LOGGER: ", log.LUTC)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"PING":"PONG!"}`)
	}))
	defer srv.Close()
	err := Connect("abcdefg", WithPing(), WithEndpoint(srv.URL), WithLogger(logger), WithDebug())
	require.NoError(t, err)

	require.Contains(t, buf.String(), "TEST LOGGER: DEBUG: Status: 200 OK | StatusCode: 200")
}

type testRoundTripper struct {
	r http.RoundTripper
}

func (rt testRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("X-Mandrill-Test", "foobar")
	return rt.r.RoundTrip(r)
}
func TestWithCustomClient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("X-Mandrill-Test")
		ret := fmt.Sprintf(`{"PING":"%s"}`, h)
		fmt.Fprintln(w, ret)
	}))
	defer srv.Close()
	httpclient := &http.Client{
		Transport: testRoundTripper{r: http.DefaultTransport},
	}
	err := Connect("abcdefg", WithHTTPClient(httpclient), WithEndpoint(srv.URL))
	require.NoError(t, err)
	req := &api.UsersPing2Request{}
	resp := &api.UsersPing2Response{}
	err = globalClient.post("users/ping2", req, resp)
	require.NoError(t, err)
	require.Equal(t, "foobar", resp.Ping)
}
