package mandrill

import (
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
	Publish      bool
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

func NewTemplate(name string, opts ...TemplateOption) (*Template, error) {
	t := &Template{
		Name: name,
	}
	for _, opt := range opts {
		if err := opt(t); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (c *Client) AddTemplate(t *Template) error {
	req := api.TemplatesAddRequest{
		Name:      t.Name,
		FromEmail: t.FromEmail,
		FromName:  t.FromName,
		Labels:    t.Labels,
		Text:      t.Text,
		Code:      t.Code,
		Subject:   t.Subject,
	}
	resp := api.TemplatesAddResponse{}
	if err := c.post("templates/add", req, resp); err != nil {
		return err
	}
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
	return nil
}

func (t *Template) Add() error {
	return globalClient.AddTemplate(t)
}
