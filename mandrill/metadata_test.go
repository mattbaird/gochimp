package mandrill

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListMetaData(t *testing.T) {
	respBody := `[
		{
			"name": "group_id",
			"state": "active",
			"view_template": "<a href=\"http://yourapplication.com/user/{{value}}\">{{value}}</a>"
		}
	]`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	md, err := client.ListMetaData()
	require.NoError(t, err)
	require.Len(t, md, 1)
}

func TestDeleteMetaData(t *testing.T) {
	respBody := `
		{
			"name": "group_id",
			"state": "active",
			"view_template": "<a href=\"http://yourapplication.com/user/{{value}}\">{{value}}</a>"
		}
	`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	err = client.DeleteMetaData("group_id")
	require.NoError(t, err)
}

func TestMetaDataDelete(t *testing.T) {
	respBody := `
	{
		"name": "group_id",
		"state": "active",
		"view_template": "<a href=\"http://yourapplication.com/user/{{value}}\">{{value}}</a>"
	}
`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	err := Connect("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	md := &MetaData{
		Name:         "group_id",
		State:        "active",
		ViewTemplate: "<a href=\"http://yourapplication.com/user/{{value}}\">{{value}}</a>",
	}
	err = md.Delete()
	require.NoError(t, err)
}
