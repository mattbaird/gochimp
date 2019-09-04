package mandrill

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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
