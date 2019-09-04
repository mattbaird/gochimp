package mandrill

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTemplate(t *testing.T) {
	template, err := NewTemplate("test_template")
	require.NoError(t, err)
	require.NotNil(t, template)
}

func testTemplateOptionError() TemplateOption {
	return func(t *Template) error {
		return fmt.Errorf("this failed")
	}
}

func TestNewTemplateError(t *testing.T) {
	template, err := NewTemplate("test_template", testTemplateOptionError())
	require.Error(t, err)
	require.Nil(t, template)
}

func TestNewTemplateWithOptions(t *testing.T) {
	template, err := NewTemplate("test_template",
		WithLabels("foo", "bar"),
		WithFromEmail("test@test.com"),
		WithFromName("administrator"),
		WithSubject("my subject"),
		WithText("abcdefg"),
		WithCode("<div>foo</div>"),
	)

	require.NoError(t, err)
	require.NotNil(t, template)
	require.Equal(t, []string{"foo", "bar"}, template.Labels)
	require.Equal(t, "test@test.com", template.FromEmail)
	require.Equal(t, "administrator", template.FromName)
	require.Equal(t, "my subject", template.Subject)
	require.Equal(t, "abcdefg", template.Text)
	require.Equal(t, "<div>foo</div>", template.Code)
}

func TestTemplateAddTemplate(t *testing.T) {
	respBody := `{
		"slug": "example-template",
		"name": "Example Template",
		"labels": [
			"example-label"
		],
		"code": "<div mc:edit=\"editable\">editable content</div>",
		"subject": "example subject",
		"from_email": "from.email@example.com",
		"from_name": "Example Name",
		"text": "Example text",
		"publish_name": "Example Template",
		"publish_code": "<div mc:edit=\"editable\">different than draft content</div>",
		"publish_subject": "example publish_subject",
		"publish_from_email": "from.email.published@example.com",
		"publish_from_name": "Example Published Name",
		"publish_text": "Example published text",
		"published_at": "2013-01-01 15:30:40",
		"created_at": "2013-01-01 15:30:27",
		"updated_at": "2013-01-01 15:30:49"
	}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL), WithDebug())
	require.NoError(t, err)
	require.NotNil(t, client)
	template, err := NewTemplate("test_template",
		WithLabels("foo", "bar"),
		WithFromEmail("test@test.com"),
		WithFromName("administrator"),
		WithSubject("my subject"),
		WithText("abcdefg"),
		WithCode("<div>foo</div>"),
	)
	require.NoError(t, err)
	require.NotNil(t, template)
	err = client.AddTemplate(template)
	require.NoError(t, err)
}

func TestTemplateAdd(t *testing.T) {
	respBody := `{
		"slug": "example-template",
		"name": "Example Template",
		"labels": [
			"example-label"
		],
		"code": "<div mc:edit=\"editable\">editable content</div>",
		"subject": "example subject",
		"from_email": "from.email@example.com",
		"from_name": "Example Name",
		"text": "Example text",
		"publish_name": "Example Template",
		"publish_code": "<div mc:edit=\"editable\">different than draft content</div>",
		"publish_subject": "example publish_subject",
		"publish_from_email": "from.email.published@example.com",
		"publish_from_name": "Example Published Name",
		"publish_text": "Example published text",
		"published_at": "2013-01-01 15:30:40",
		"created_at": "2013-01-01 15:30:27",
		"updated_at": "2013-01-01 15:30:49"
	}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	err := Connect("abcdefg", WithEndpoint(srv.URL), WithDebug())
	require.NoError(t, err)
	template, err := NewTemplate("test_template",
		WithLabels("foo", "bar"),
		WithFromEmail("test@test.com"),
		WithFromName("administrator"),
		WithSubject("my subject"),
		WithText("abcdefg"),
		WithCode("<div>foo</div>"),
	)
	require.NoError(t, err)
	require.NotNil(t, template)
	err = template.Add()
	require.NoError(t, err)
}
