package mandrill

import (
	"context"
	"sync"
	"time"

	"github.com/lusis/gochimp/mandrill/api"
)

// Template represents a template (published or unpublished)
type Template struct {
	sync.RWMutex // safety first
	Name         string
	FromEmail    string
	FromName     string
	Subject      string
	Code         string
	Text         string
	Labels       []string

	// the values below are populated after creating the template via api call
	// the are ignored when creating a new template
	Slug             string
	PublishName      string
	PublishCode      string
	PublishSubject   string
	PublishFromEmail string
	PublishFromName  string
	PublishText      string
	PublishedAt      time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// TemplateOption is an option you can pass to to template creation
type TemplateOption func(*Template) error

// WithCode sets the code for a template
func WithCode(code string) TemplateOption {
	return func(t *Template) error {
		t.Code = code
		return nil
	}
}

// WithText sets the template text
func WithText(text string) TemplateOption {
	return func(t *Template) error {
		t.Text = text
		return nil
	}
}

// WithSubject sets the template subject
func WithSubject(s string) TemplateOption {
	return func(t *Template) error {
		t.Subject = s
		return nil
	}
}

// WithFromEmail sets the template FromEmail
func WithFromEmail(s string) TemplateOption {
	return func(t *Template) error {
		t.FromEmail = s
		return nil
	}
}

// WithFromName sets the template FromName
func WithFromName(s string) TemplateOption {
	return func(t *Template) error {
		t.FromName = s
		return nil
	}
}

// WithLabels adds labels to the template
func WithLabels(labels ...string) TemplateOption {
	return func(t *Template) error {
		t.Labels = labels
		return nil
	}
}

// NewTemplate creates a new Template
func NewTemplate(name string, opts ...TemplateOption) (*Template, error) {
	t := &Template{
		Name: name,
	}
	for _, opt := range opts {
		t.Lock()
		if err := opt(t); err != nil {
			return nil, err
		}
		t.Unlock()
	}
	return t, nil
}

// AddTemplate adds the provided Template to Mandrill via the API
// all templates added via this library are unpublished
func (c *Client) AddTemplate(t *Template) error {
	return c.AddTemplateContext(context.TODO(), t)
}

// AddTemplateContext adds the provided Template to Mandrill via the API
// all templates added via this library are unpublished
func (c *Client) AddTemplateContext(ctx context.Context, t *Template) error {
	req := &api.TemplatesAddRequest{
		Name:      t.Name,
		FromEmail: t.FromEmail,
		FromName:  t.FromName,
		Labels:    t.Labels,
		Text:      t.Text,
		Code:      t.Code,
		Subject:   t.Subject,
	}
	resp := &api.TemplatesAddResponse{}
	if err := c.postContext(ctx, "templates/add", req, resp); err != nil {
		return err
	}
	t.Lock()
	t.Slug = resp.Slug
	t.CreatedAt = resp.CreatedAt.Time
	t.UpdatedAt = resp.UpdatedAt.Time
	t.PublishedAt = resp.PublishedAt.Time
	t.PublishName = resp.PublishName
	t.PublishCode = resp.PublishCode
	t.PublishSubject = resp.PublishSubject
	t.PublishFromEmail = resp.PublishFromEmail
	t.PublishFromName = resp.PublishFromName
	t.PublishText = resp.PublishText
	t.Unlock()
	return nil
}

// Add adds the current template by calling the Mandrill API
func (t *Template) Add() error {
	return globalClient.AddTemplate(t)
}

// PublishTemplate publishes the named template
func (c *Client) PublishTemplate(name string) error {
	return c.PublishTemplateContext(context.TODO(), name)
}

// PublishTemplateContext publishes the named template with context
func (c *Client) PublishTemplateContext(ctx context.Context, name string) error {
	req := &api.TemplatesPublishRequest{
		Name: name,
	}
	resp := &api.TemplatesPublishResponse{}
	if err := c.postContext(ctx, "templates/publish", req, resp); err != nil {
		return err
	}
	return nil
}

// Publish publishes the current template
func (t *Template) Publish() error {
	return globalClient.PublishTemplate(t.Name)
}

// GetTemplateInfo gets the information about a template
func (c *Client) GetTemplateInfo(name string) (*Template, error) {
	return c.GetTemplateInfoContext(context.TODO(), name)
}

// GetTemplateInfoContext gets the information about a template
func (c *Client) GetTemplateInfoContext(ctx context.Context, name string) (*Template, error) {
	req := &api.TemplatesInfoRequest{
		Name: name,
	}
	resp := &api.TemplatesInfoResponse{}
	if err := c.postContext(ctx, "templates/info", req, resp); err != nil {
		return nil, err
	}
	t := &Template{
		Slug:             resp.Slug,
		Name:             resp.Name,
		Labels:           resp.Labels,
		Code:             resp.Code,
		Subject:          resp.Subject,
		FromEmail:        resp.FromEmail,
		FromName:         resp.FromName,
		Text:             resp.Text,
		PublishName:      resp.PublishName,
		PublishCode:      resp.PublishCode,
		PublishSubject:   resp.PublishSubject,
		PublishFromEmail: resp.PublishFromEmail,
		PublishFromName:  resp.PublishFromName,
		PublishText:      resp.PublishText,
		PublishedAt:      resp.PublishedAt.Time,
		CreatedAt:        resp.CreatedAt.Time,
		UpdatedAt:        resp.UpdatedAt.Time,
	}
	return t, nil

}

// Info gets info from the api about the current template
// in reality this replaces the current template values with the version from the api
func (t *Template) Info() error {
	resp, err := globalClient.GetTemplateInfo(t.Name)
	if err != nil {
		return err
	}
	t.Lock()
	t.Slug = resp.Slug
	t.Name = resp.Name
	t.Labels = resp.Labels
	t.Code = resp.Code
	t.Subject = resp.Subject
	t.FromEmail = resp.FromEmail
	t.FromName = resp.FromName
	t.Text = resp.Text
	t.PublishName = resp.PublishName
	t.PublishCode = resp.PublishCode
	t.PublishSubject = resp.PublishSubject
	t.PublishFromEmail = resp.PublishFromEmail
	t.PublishFromName = resp.PublishFromName
	t.PublishText = resp.PublishText
	t.PublishedAt = resp.PublishedAt
	t.CreatedAt = resp.CreatedAt
	t.UpdatedAt = resp.UpdatedAt
	t.Unlock()
	return nil
}

// DeleteTemplate deletes the named template
func (c *Client) DeleteTemplate(name string) error {
	req := &api.TemplatesDeleteRequest{
		Name: name,
	}
	resp := &api.TemplatesDeleteResponse{}
	if err := c.post("templates/delete", req, resp); err != nil {
		return err
	}
	return nil
}

// Delete deletes the current template
func (t *Template) Delete() error {
	return globalClient.DeleteTemplate(t.Name)
}

// ListTemplates lists all templates available to the current user
func (c *Client) ListTemplates() ([]*Template, error) {
	req := &api.TemplatesListRequest{}
	resp := &api.TemplatesListResponse{}
	if err := c.post("templates/list", req, resp); err != nil {
		return nil, err
	}
	all := make([]*Template, len(*resp))
	for idx, template := range *resp {
		t := &Template{
			Slug:             template.Slug,
			Name:             template.Name,
			Labels:           template.Labels,
			Code:             template.Code,
			Subject:          template.Subject,
			FromEmail:        template.FromEmail,
			FromName:         template.FromName,
			Text:             template.Text,
			PublishName:      template.PublishName,
			PublishCode:      template.PublishCode,
			PublishSubject:   template.PublishSubject,
			PublishFromEmail: template.PublishFromEmail,
			PublishFromName:  template.PublishFromName,
			PublishText:      template.PublishText,
			PublishedAt:      template.PublishedAt.Time,
			CreatedAt:        template.CreatedAt.Time,
			UpdatedAt:        template.UpdatedAt.Time,
		}
		all[idx] = t
	}
	return all, nil
}

// RenderTemplate renders the template with the provided values
// content and vars maps will be converted like so:
// given: {"foo": "bar"}
// becomes: {Name: "foo", Content: "bar"}
func (c *Client) RenderTemplate(name string, content []map[string]string, vars []map[string]string) (string, error) {
	req := &api.TemplatesRenderRequest{
		TemplateName: name,
	}
	resp := &api.TemplatesRenderResponse{}
	for _, c := range content {
		for k, v := range c {
			req.TemplateContent = append(req.TemplateContent, api.TemplatesRenderVars{Name: k, Content: v})
		}
	}
	for _, c := range vars {
		for k, v := range c {
			req.MergeVars = append(req.MergeVars, api.TemplatesRenderVars{Name: k, Content: v})
		}
	}
	if err := c.post("templates/render", req, resp); err != nil {
		return "", err
	}
	return resp.HTML, nil
}

// Render renders the current template with the supplied vars
// content and vars maps will be converted like so:
// given: {"foo": "bar"}
// becomes: {Name: "foo", Content: "bar"}
func (t *Template) Render(content []map[string]string, vars []map[string]string) (string, error) {
	return globalClient.RenderTemplate(t.Name, content, vars)
}
