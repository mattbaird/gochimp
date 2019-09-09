package mandrill

import (
	"fmt"
	"io/ioutil"
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
	client, err := New("abcdefg", WithEndpoint(srv.URL))
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
	err := Connect("abcdefg", WithEndpoint(srv.URL))
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

func TestPublishTemplate(t *testing.T) {
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
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	err = client.PublishTemplate("Example Template")
	require.NoError(t, err)
}

func TestGetTemplateInfo(t *testing.T) {
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
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	template, err := client.GetTemplateInfo("Example Template")
	require.NoError(t, err)
	require.NotNil(t, template)
	require.Equal(t, "example-template", template.Slug)
	require.Equal(t, "Example Template", template.Name)
	require.Equal(t, []string{"example-label"}, template.Labels)
	require.Equal(t, "<div mc:edit=\"editable\">editable content</div>", template.Code)
	require.Equal(t, "example subject", template.Subject)
	require.Equal(t, "from.email@example.com", template.FromEmail)
	require.Equal(t, "Example Name", template.FromName)
	require.Equal(t, "Example text", template.Text)
	require.Equal(t, "Example Template", template.PublishName)
	require.NotNil(t, template.PublishedAt)
}

func TestTemplatePublish(t *testing.T) {
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
	err := Connect("abcdefg", WithEndpoint(srv.URL))
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
	err = template.Publish()
	require.NoError(t, err)
}

func TestTemplateInfo(t *testing.T) {
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
	err := Connect("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	template, err := NewTemplate("test_template", WithLabels("foo", "bar"))
	require.NoError(t, err)
	require.NotNil(t, template)
	require.Equal(t, []string{"foo", "bar"}, template.Labels)
	err = template.Info()
	require.NoError(t, err)
	require.Equal(t, []string{"example-label"}, template.Labels)
}

func TestDeleteTemplate(t *testing.T) {
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
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	err = client.DeleteTemplate("Example Template")
	require.NoError(t, err)
}

func TestTemplateDelete(t *testing.T) {
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
	_, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	template, err := NewTemplate("foo")
	require.NoError(t, err)
	require.NotNil(t, template)
	err = template.Delete()
	require.NoError(t, err)
}

func TestListTemplate(t *testing.T) {
	respBody := `[
    {
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
    }
]`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	client, err := New("abcdefg", WithEndpoint(srv.URL))
	require.NoError(t, err)
	require.NotNil(t, client)
	templates, err := client.ListTemplates()
	require.NoError(t, err)
	require.Len(t, templates, 1)
	require.Equal(t, "example-template", templates[0].Slug)
	require.Equal(t, "Example Template", templates[0].Name)
	require.Len(t, templates[0].Labels, 1)
	require.NotNil(t, templates[0].PublishedAt)
}

func TestRenderTemplateTable(t *testing.T) {
	testCases := map[string]struct {
		Expected string
		Content  []map[string]string
		Vars     []map[string]string
	}{
		"all_params": {
			Content: []map[string]string{
				{"foo": "bar"},
			},
			Vars: []map[string]string{
				{"baz": "qux"},
			},
			Expected: `{"key":"abcdefg","template_name":"foo","template_content":[{"name":"foo","content":"bar"}],"merge_vars":[{"name":"baz","content":"qux"}]}`,
		},
		"no_vars": {
			Content: []map[string]string{
				{"foo": "bar"},
			},
			Vars:     []map[string]string{},
			Expected: `{"key":"abcdefg","template_name":"foo","template_content":[{"name":"foo","content":"bar"}]}`,
		},
		"no_content": {
			Content: []map[string]string{},
			Vars: []map[string]string{
				{"baz": "qux"},
			},
			Expected: `{"key":"abcdefg","template_name":"foo","template_content":null,"merge_vars":[{"name":"baz","content":"qux"}]}`,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			var results string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, err := ioutil.ReadAll(r.Body)
				require.NoError(t, err)
				defer r.Body.Close()
				results = string(b)
				fmt.Fprintln(w, fmt.Sprintf(`{"html": "%s"}`, name))
			}))
			defer srv.Close()
			client, err := New("abcdefg", WithEndpoint(srv.URL))
			require.NoError(t, err)
			require.NotNil(t, client)

			r, err := client.RenderTemplate("foo", tc.Content, tc.Vars)
			require.NoError(t, err)
			require.NotEmpty(t, r)
			require.Equal(t, tc.Expected, results)
			require.Equal(t, name, r)
		})
	}
}

func TestTemplateRender(t *testing.T) {
	respBody := `{"html":"foobar"}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}))
	defer srv.Close()
	err := Connect("abcdefg", WithEndpoint(srv.URL))
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
	res, err := template.Render([]map[string]string{{"foo": "bar"}}, []map[string]string{{"foo": "bar"}})
	require.NoError(t, err)
	require.Equal(t, "foobar", res)
}
