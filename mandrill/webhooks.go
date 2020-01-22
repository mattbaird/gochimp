package mandrill

import (
	"context"
	"github.com/lusis/gochimp/mandrill/api"
	"time"
)

// Webhook represents a Mandrill webhook
type Webhook struct {
	ID          int
	URL         string
	Description string
	Events      []string
	AuthKey     string
	CreatedAt   time.Time
	LastSentAt  time.Time
	BatchesSent int32
	EventsSent  int32
	LastError   string
}

// Info gets info about the current webhook
// in reality this refreshes the current webhook from the api
func (w *Webhook) Info() error {
	return w.InfoContext(context.TODO())
}

// InfoContext gets information about the current webhook
// in reality this refreshes the current webhook from the api
func (w *Webhook) InfoContext(ctx context.Context) error {
	h, err := globalClient.WebhookInfoContext(ctx, w.ID)
	*w = *h
	return err
}

// Add adds the current webhook
func (w *Webhook) Add() error {
	return w.AddContext(context.TODO())
}

// AddContext adds the current webhook
func (w *Webhook) AddContext(ctx context.Context) error {
	return globalClient.AddWebHookContext(ctx, w)
}

// Update updates the current webhook
func (w *Webhook) Update() error {
	return w.UpdateContext(context.TODO())
}

// UpdateContext updates the current webhook
func (w *Webhook) UpdateContext(ctx context.Context) error {
	return globalClient.UpdateWebhookContext(ctx, w)
}

// Delete deletes the current webhook
func (w *Webhook) Delete() error {
	return w.DeleteContext(context.TODO())
}

// DeleteContext deletes the current webhook
func (w *Webhook) DeleteContext(ctx context.Context) error {
	return globalClient.DeleteWebhookContext(ctx, w.ID)
}

// WebhookInfo returns information about the webhook by id
func (c *Client) WebhookInfo(id int) (*Webhook, error) {
	return c.WebhookInfoContext(context.TODO(), id)
}

// WebhookInfoContext returns information about the webhook by id
func (c *Client) WebhookInfoContext(ctx context.Context, id int) (*Webhook, error) {
	req := &api.WebhooksInfoRequest{ID: id}
	resp := &api.WebhooksInfoResponse{}
	err := c.postContext(ctx, "webhooks/info", req, resp)
	if err != nil {
		return nil, err
	}
	return &Webhook{
		ID:          resp.ID,
		URL:         resp.URL,
		Description: resp.Description,
		Events:      resp.Events,
		AuthKey:     resp.AuthKey,
		CreatedAt:   resp.CreatedAt.Time,
		LastSentAt:  resp.LastSentAt.Time,
		BatchesSent: resp.BatchesSent,
		EventsSent:  resp.EventsSent,
		LastError:   resp.LastError,
	}, nil
}

// ListWebhooks lists all webhooks
func (c *Client) ListWebhooks() ([]*Webhook, error) {
	return c.ListWebhooksContext(context.TODO())
}

// ListWebhooksContext lists all webhooks
func (c *Client) ListWebhooksContext(ctx context.Context) ([]*Webhook, error) {
	return nil, nil
}

// AddWebhook adds a webhook
func (c *Client) AddWebhook(w *Webhook) error {
	return c.AddWebHookContext(context.TODO(), w)
}

// AddWebHookContext adds a webhook
func (c *Client) AddWebHookContext(ctx context.Context, w *Webhook) error {
	req := &api.WebhooksAddRequest{
		URL:         w.URL,
		Description: w.Description,
		Events:      w.Events,
	}
	resp := &api.WebhooksAddResponse{}
	err := c.postContext(ctx, "webhooks/add", req, resp)
	if err != nil {
		return err
	}
	newhook, err := c.WebhookInfoContext(ctx, resp.ID)
	if err != nil {
		return err
	}
	*w = *newhook
	return nil
}

// UpdateWebhook updates a webhook
func (c *Client) UpdateWebhook(w *Webhook) error {
	return c.UpdateWebhookContext(context.TODO(), w)
}

// UpdateWebhookContext updates a webhook
func (c *Client) UpdateWebhookContext(ctx context.Context, w *Webhook) error {
	return nil
}

// DeleteWebhook deletes a webhook
func (c *Client) DeleteWebhook(id int) error {
	return c.DeleteWebhookContext(context.TODO(), id)
}

// DeleteWebhookContext deletes a webhook
func (c *Client) DeleteWebhookContext(ctx context.Context, id int) error {
	req := &api.WebhooksDeleteRequest{ID: id}
	resp := &api.WebhooksDeleteResponse{}
	return c.postContext(ctx, "webhooks/delete", req, resp)
}
