package mandrill

import (
	"github.com/lusis/gochimp/mandrill/api"
)

// TemplateMessage represents a templated message in Mandrill
type TemplateMessage struct {
	TemplateName    string
	TemplateContent []TemplateVar
	Message         Message
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

// SendMessage sends a mandrill.Message
func (c *Client) SendMessage(m Message) ([]MessageStatus, error) {
	msg := api.MessagesSendRequest{
		Message: m.toAPIMessage(),
	}
	return c.messageSend(msg)
}

// SendTemplate sends a mandrill.TemplateMessage
func (c *Client) SendTemplate(m TemplateMessage) ([]MessageStatus, error) {
	return c.templateSend(m.toAPITemplateMessage())
}

func (c *Client) messageSend(m api.MessagesSendRequest) ([]MessageStatus, error) {
	resp := &api.MessagesSendResponse{}
	if err := c.post("messages/send", &m, resp); err != nil {
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
	resp := &api.MessagesSendResponse{}
	if err := c.post("messages/send-template", &m, resp); err != nil {
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
