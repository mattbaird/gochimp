package mandrill

import (
	"context"
	"fmt"
	"time"

	"github.com/lusis/gochimp/mandrill/api"
)

// TemplateMessage represents a templated message in Mandrill
type TemplateMessage struct {
	TemplateName    string
	TemplateContent []TemplateVar
	Message         Message
}

// Content represents the results of a messages/content api call
type Content struct {
	FromEmail   string
	FromName    string
	Subject     string
	To          []To
	Text        string
	HTML        string
	Headers     map[string]string
	Timestamp   time.Time
	Tags        []string
	ID          string
	Attachments []Attachment
	Images      []Image
}

// Message represents a message in mandrill
type Message struct {
	HTML                    string
	Text                    string
	Subject                 string
	FromEmail               string
	FromName                string
	To                      []To
	Headers                 map[string]string
	Important               bool
	TrackOpens              bool
	TrackClicks             bool
	AutoText                bool
	AutoHTML                bool
	InlineCSS               bool
	URLStripQS              bool
	PreserveRecipients      bool
	ViewContentLink         bool
	BccAddress              string
	TrackingDomain          string
	SigningDomain           string
	ReturnPathDomain        string
	Merge                   bool
	MergeLanguage           string
	GlobalMergeVars         []Var
	MergeVars               []MergeVar
	Tags                    []string
	SubAccount              string
	GoogleAnalyticsDomains  []string
	GoogleAnalyticsCampaign []string
	MetaData                map[string]string
	RecipientMetaData       []RecipientMetaData
	Attachments             []Attachment
	Images                  []Image
	Async                   bool
	IPPool                  string
}

// TemplateVar is a variable for template content
type TemplateVar struct {
	Name    string
	Content string
}

// Var is an individual merge variable
type Var struct {
	Name    string
	Content interface{}
}

// MergeVar represents a destination-specific merge variable
type MergeVar struct {
	Rcpt string
	Vars []Var
}

// To represents the destination of an email
type To struct {
	Email string
	Name  string
	Type  string
}

// Attachment represents a message attachment
type Attachment struct {
	Type    string
	Name    string
	Content string
	Binary  bool
}

// Image represents an image in a message
type Image struct {
	Type    string
	Name    string
	Content string
}

// RecipientMetaData represents per-recipient custom key/value data
type RecipientMetaData struct {
	Rcpt   string
	Values map[string]string
}

// MessageStatus represents the send status of a message
type MessageStatus struct {
	Email        string
	Status       string
	RejectReason string
	ID           string
}

// MessageInfo represents the result of a message search
type MessageInfo struct {
	Timestamp     time.Time
	ID            string
	Sender        string
	Template      string
	Subject       string
	Email         string
	Tags          []string
	Opens         int
	OpensDetails  []ClicksOpensDetail
	Clicks        int
	ClicksDetails []ClicksOpensDetail
	State         string
	MetaData      map[string]string
	SMTPEvents    []SMTPEvent
}

// SMTPEvent represents the smtp_events details
type SMTPEvent struct {
	Timestamp     time.Time
	Type          string
	Diag          string
	SourceIP      string
	DestinationIP string
	Size          int32
}

// ClicksOpensDetail represents the clicks/opens details
type ClicksOpensDetail struct {
	Timestamp time.Time
	URL       string
	IP        string
	Location  string
	UserAgent string
}

// MessageSearchParams represents the params for searching messages
type MessageSearchParams struct {
	Query    string
	DateFrom *time.Time
	DateTo   *time.Time
	Tags     []string
	Senders  []string
	APIKeys  []string
	Limit    int
}

// MessageContent represents a call to messages/content
func (c *Client) MessageContent(id string) (*Content, error) {
	return c.MessageContentContext(context.TODO(), id)
}

// MessageContentContext makes a call to messages/content with the provided context
func (c *Client) MessageContentContext(ctx context.Context, id string) (*Content, error) {
	resp := &api.MessagesContentResponse{}
	req := &api.MessagesContentRequest{
		ID: id,
	}
	if err := c.postContext(ctx, "messages/content", req, resp); err != nil {
		return nil, err
	}
	res := &Content{
		ID:        resp.ID,
		Text:      resp.Text,
		HTML:      resp.HTML,
		Timestamp: time.Unix(int64(resp.TS), 0),
		FromEmail: resp.FromEmail,
		FromName:  resp.FromName,
		Subject:   resp.Subject,
		Tags:      resp.Tags,
		Headers:   resp.Headers,
	}
	for _, t := range resp.To {
		res.To = append(res.To, To{Email: t.Email, Name: t.Name})
	}
	for _, a := range resp.Attachments {
		res.Attachments = append(res.Attachments, Attachment{Name: a.Name, Type: a.Type, Binary: a.Binary, Content: a.Content})
	}
	for _, i := range resp.Images {
		res.Images = append(res.Images, Image{Name: i.Name, Type: i.Type, Content: i.Content})
	}
	return res, nil
}

// SearchMessages calls messages/search
func (c *Client) SearchMessages(m MessageSearchParams) ([]MessageInfo, error) {
	return c.SearchMessagesContext(context.TODO(), m)
}

// SearchMessagesContext calls messages/search with the provided context
func (c *Client) SearchMessagesContext(ctx context.Context, m MessageSearchParams) ([]MessageInfo, error) {
	resp := &api.MessagesSearchResponse{}
	req := api.MessagesSearchRequest{
		Query:   m.Query,
		Tags:    m.Tags,
		Senders: m.Senders,
		APIKeys: m.APIKeys,
		Limit:   m.Limit,
	}
	if m.DateTo != nil {
		req.DateTo = fmt.Sprintf("%d-%d-%d", m.DateTo.Year(), m.DateTo.Month(), m.DateTo.Day())
	}
	if m.DateFrom != nil {
		req.DateFrom = fmt.Sprintf("%d-%d-%d", m.DateFrom.Year(), m.DateFrom.Month(), m.DateFrom.Day())
	}
	if err := c.postContext(ctx, "messages/search", &req, resp); err != nil {
		return nil, err
	}
	results := []MessageInfo{}
	for _, res := range *resp {
		r := toMessageInfo(res)
		results = append(results, r)
	}
	return results, nil
}

// MessageInfo returns the messages/info api call
func (c *Client) MessageInfo(id string) (*MessageInfo, error) {
	return c.MessageInfoContext(context.TODO(), id)
}

// MessageInfoContext returns the messages/info api call with provided context
func (c *Client) MessageInfoContext(ctx context.Context, id string) (*MessageInfo, error) {
	resp := &api.MessagesInfoResponse{}
	req := &api.MessagesInfoRequest{ID: id}
	if err := c.postContext(ctx, "messages/info", req, resp); err != nil {
		return nil, err
	}
	r := toMessageInfo(*resp)
	return &r, nil
}

func toMessageInfo(res api.MessagesInfoResponse) MessageInfo {
	r := MessageInfo{
		ID:       res.ID,
		Sender:   res.Sender,
		Template: res.Template,
		Subject:  res.Subject,
		Email:    res.Email,
		Tags:     res.Tags,
		State:    res.State,
		MetaData: res.MetaData,
		Opens:    res.Opens,
		Clicks:   res.Clicks,
	}
	r.Timestamp = time.Unix(int64(res.TS), 0)
	for _, d := range res.OpensDetail {
		od := ClicksOpensDetail{
			Timestamp: time.Unix(int64(d.TS), 0),
			IP:        d.IP,
			UserAgent: d.UA,
			Location:  d.Location,
		}
		r.OpensDetails = append(r.OpensDetails, od)
	}
	for _, d := range res.ClicksDetail {
		cd := ClicksOpensDetail{
			Timestamp: time.Unix(int64(d.TS), 0),
			IP:        d.IP,
			UserAgent: d.UA,
			Location:  d.Location,
			URL:       d.URL,
		}
		r.ClicksDetails = append(r.ClicksDetails, cd)
	}
	for _, d := range res.SMTPEvents {
		se := SMTPEvent{
			Timestamp:     time.Unix(int64(d.TS), 0),
			Type:          d.Type,
			Diag:          d.Diag,
			SourceIP:      d.SourceIP,
			DestinationIP: d.DestinationIP,
			Size:          d.Size,
		}
		r.SMTPEvents = append(r.SMTPEvents, se)
	}
	return r
}

// SendMessage sends a mandrill.Message
func (c *Client) SendMessage(m Message) ([]MessageStatus, error) {
	return c.SendMessageContext(context.TODO(), m)
}

// SendMessageContext sends a mandrill.Message with context
func (c *Client) SendMessageContext(ctx context.Context, m Message) ([]MessageStatus, error) {
	msg := api.MessagesSendRequest{
		Message: m.toAPIMessage(),
	}
	return c.messageSendContext(ctx, msg)
}

// SendTemplate sends a mandrill.TemplateMessage
func (c *Client) SendTemplate(m TemplateMessage) ([]MessageStatus, error) {
	return c.SendTemplateContext(context.TODO(), m)
}

// SendTemplateContext sends a mandrill.TemplateMessage
func (c *Client) SendTemplateContext(ctx context.Context, m TemplateMessage) ([]MessageStatus, error) {
	return c.templateSendContext(ctx, m.toAPITemplateMessage())
}

func (c *Client) messageSend(m api.MessagesSendRequest) ([]MessageStatus, error) {
	return c.messageSendContext(context.TODO(), m)
}

func (c *Client) messageSendContext(ctx context.Context, m api.MessagesSendRequest) ([]MessageStatus, error) {
	resp := &api.MessagesSendResponse{}
	if err := c.postContext(ctx, "messages/send", &m, resp); err != nil {
		return nil, err
	}
	status := []MessageStatus{}
	for _, s := range *resp {
		status = append(status, MessageStatus{
			Email:        s.Email,
			Status:       s.Status,
			RejectReason: s.RejectReason,
			ID:           s.ID,
		})
	}
	return status, nil
}

func (c *Client) templateSend(m api.MessagesSendTemplateRequest) ([]MessageStatus, error) {
	return c.templateSendContext(context.TODO(), m)
}

func (c *Client) templateSendContext(ctx context.Context, m api.MessagesSendTemplateRequest) ([]MessageStatus, error) {
	resp := &api.MessagesSendResponse{}
	if err := c.postContext(ctx, "messages/send-template", &m, resp); err != nil {
		return nil, err
	}
	status := []MessageStatus{}
	for _, s := range *resp {
		status = append(status, MessageStatus{
			Email:        s.Email,
			Status:       s.Status,
			RejectReason: s.RejectReason,
			ID:           s.ID,
		})
	}
	return status, nil
}

// converts our representation of a template messsage to the api format
func (m TemplateMessage) toAPITemplateMessage() api.MessagesSendTemplateRequest {
	mst := api.MessagesSendTemplateRequest{
		TemplateName: m.TemplateName,
		Message:      m.Message.toAPIMessage(),
	}
	for _, tc := range m.TemplateContent {
		mst.TemplateContent = append(mst.TemplateContent, api.MessagesTemplateContent(tc))
	}
	return mst
}

// converts our representation of a message to the api format
func (m Message) toAPIMessage() api.MessageRequest {
	apiMsg := api.MessageRequest{
		HTML:                    m.HTML,
		Text:                    m.Text,
		Subject:                 m.Subject,
		FromEmail:               m.FromEmail,
		FromName:                m.FromName,
		Headers:                 m.Headers,
		Important:               m.Important,
		TrackOpens:              m.TrackOpens,
		TrackClicks:             m.TrackClicks,
		AutoText:                m.AutoText,
		AutoHTML:                m.AutoHTML,
		InlineCSS:               m.InlineCSS,
		URLStripQS:              m.URLStripQS,
		PreserveRecipients:      m.PreserveRecipients,
		ViewContentLink:         m.ViewContentLink,
		BccAddress:              m.BccAddress,
		TrackingDomain:          m.TrackingDomain,
		SigningDomain:           m.SigningDomain,
		ReturnPathDomain:        m.ReturnPathDomain,
		Merge:                   m.Merge,
		MergeLanguage:           m.MergeLanguage,
		Tags:                    m.Tags,
		GoogleAnalyticsDomains:  m.GoogleAnalyticsDomains,
		GoogleAnalyticsCampaign: m.GoogleAnalyticsCampaign,
		MetaData:                m.MetaData,
		Async:                   m.Async,
		IPPool:                  m.IPPool,
	}
	// subaccount assignment
	if globalClient.subaccount != "" {
		apiMsg.SubAccount = globalClient.subaccount
	}
	// convert message.To to api.To
	apiMsg.To = []api.MessageTo{}
	for _, t := range m.To {
		apiMsg.To = append(apiMsg.To, api.MessageTo{
			Email: t.Email,
			Name:  t.Name,
			Type:  t.Type,
		})
	}
	// convert message.GlobalMergeVars to api.GlobalMergeVars
	apiMsg.GlobalMergeVars = []api.MessageVar{}
	for _, gmv := range m.GlobalMergeVars {
		apiMsg.GlobalMergeVars = append(apiMsg.GlobalMergeVars, api.MessageVar(gmv))
	}

	// convert message.MergeVars to api.MessageMergeVars
	apiMergeVars := []api.MessageMergeVar{}
	for _, mv := range m.MergeVars {
		amv := []api.MessageVar{}
		for _, a := range mv.Vars {
			amv = append(amv, api.MessageVar(a))
		}
		apiMergeVars = append(apiMergeVars, api.MessageMergeVar{
			Rcpt: mv.Rcpt,
			Vars: amv,
		})
	}
	apiMsg.MergeVars = apiMergeVars

	// convert message.MetaData to api.MessageRecipientMetaData
	apiMsg.RecipientMetaData = []api.MessageRecipientMetaData{}
	for _, md := range m.RecipientMetaData {
		apiMsg.RecipientMetaData = append(apiMsg.RecipientMetaData, api.MessageRecipientMetaData(md))
	}
	return apiMsg
}
