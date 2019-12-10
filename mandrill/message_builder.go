package mandrill

import (
	"context"
	"sync"
)

/*
The bare minimum needed to send a message through mandrill is the following JSON:

{
    "key": "<valid api key>",
    "message": {
        "from_email": "user@domain.com",
        "to": [
            {
                "email": "user@destination.com"
            }
        ]
    }
}

However not having a subject and not having a body tends to trigger spam filters
So the minimum JSON needed to send an mail is realistically

{
    "key": "<valid api key>",
    "message": {
        "text": "Example text content",
        "subject": "example subject",
        "from_email": "user@domain.com",
        "to": [
            {
                "email": "user@destination.com"
            }
        ]
    }
}

*/
// NewMessageBuilder returns a new MessageBuilder
func NewMessageBuilder(from string, name string) *MessageBuilder {
	// We're only ever going to have one "from" and the rest of the data may be variable
	// So we start here requiring a from_email/from_name
	builder := &MessageBuilder{}
	msg := &Message{
		FromEmail: from,
		FromName:  name,
	}
	builder.message = msg
	return builder
}

// MessageBuilder is a customized representation of a mandrill.Message
// The goal is to facilitate Message creation
type MessageBuilder struct {
	lock            sync.RWMutex // ensure thread-safe message building
	finalized       bool
	isTemplate      bool
	isRaw           bool
	recipients      []Recipient
	message         *Message
	templateName    string
	templateContent []TemplateVar
}

// Recipient is a custom type that collocates all information about a recipient.
//
// When sending a message in Mandrill, you are required to populate the To data
// However there may be commensurate data when using templates that needs to go
// into RecipientMetaData or MergeVars for just that recipient.
//
// The Recipient type allows specifying all the information up front for later distribution to
// the appropriate api format
type Recipient struct {
	Email     string
	Name      string
	Type      string
	MergeVars []Var
	MetaData  map[string]string
}

// AddRecipients adds a collection of recipients to a message
func (m *MessageBuilder) AddRecipients(recipients []Recipient) *MessageBuilder {
	m.lock.Lock()
	if m.recipients == nil {
		m.recipients = []Recipient{}
	}
	m.recipients = append(m.recipients, recipients...)
	m.lock.Unlock()
	return m
}

// AddRecipient adds a single recipient to a message
func (m *MessageBuilder) AddRecipient(recipient Recipient) *MessageBuilder {
	rcpt := []Recipient{}
	rcpt = append(rcpt, recipient)
	return m.AddRecipients(rcpt)
}

// WithSubject sets the subject of the email
func (m *MessageBuilder) WithSubject(subject string) *MessageBuilder {
	m.lock.Lock()
	m.message.Subject = subject
	m.lock.Unlock()
	return m
}

// MergeAs sets the merge language to be used and enables merging
// valid options are 'mailchimp' or 'handlebars'
func (m *MessageBuilder) MergeAs(language string) *MessageBuilder {
	m.lock.Lock()
	m.message.MergeLanguage = language
	m.message.Merge = true
	m.lock.Unlock()
	return m
}

// WithText specifies the text body of an email
func (m *MessageBuilder) WithText(text string) *MessageBuilder {
	m.lock.Lock()
	m.message.Text = text
	m.lock.Unlock()
	return m
}

// WithHTML specifies the html version of the email body
func (m *MessageBuilder) WithHTML(html string) *MessageBuilder {
	m.lock.Lock()
	m.message.HTML = html
	m.lock.Unlock()
	return m
}

// WithTemplate ensures that the minimum amount of required data is passed in
func (m *MessageBuilder) WithTemplate(name string, vars []TemplateVar) *MessageBuilder {
	m.lock.Lock()
	m.isTemplate = true
	m.templateName = name
	m.templateContent = vars
	m.lock.Unlock()
	return m
}

// WithHeaders specifies any custom headers to send with the message
func (m *MessageBuilder) WithHeaders(headers map[string]string) *MessageBuilder {
	m.lock.Lock()
	m.message.Headers = headers
	m.lock.Unlock()
	return m
}

// WithAutoHTML enables generating html versions of emails automatically
func (m *MessageBuilder) WithAutoHTML() *MessageBuilder {
	m.lock.Lock()
	m.message.AutoHTML = true
	m.lock.Unlock()
	return m
}

// WithAutoText enables generating text versions of emails automatically
func (m *MessageBuilder) WithAutoText() *MessageBuilder {
	m.lock.Lock()
	m.message.AutoText = true
	m.lock.Unlock()
	return m
}

func (m *MessageBuilder) finalize() {
	m.lock.Lock()
	// do work to build all the appropriate API representations
	m.rcptToMsg()
	m.finalized = true
	m.lock.Unlock()
}

// this converts the Recipient type to the format needed for Message
func (m *MessageBuilder) rcptToMsg() {
	// initialize our slices if needed
	if len(m.message.MergeVars) == 0 {
		m.message.MergeVars = []MergeVar{}
	}
	if len(m.message.To) == 0 {
		m.message.To = []To{}
	}
	if len(m.message.RecipientMetaData) == 0 {
		m.message.RecipientMetaData = []RecipientMetaData{}
	}
	for _, r := range m.recipients {
		// Append a new To from this recipient
		m.message.To = append(m.message.To, To{
			Email: r.Email,
			Name:  r.Name,
			Type:  r.Type,
		})
		// Populate MergeVars if needed
		if len(r.MergeVars) != 0 {
			m.message.MergeVars = append(m.message.MergeVars, MergeVar{
				Rcpt: r.Email,
				Vars: r.MergeVars,
			})
		}
		// Populate RecipientMetaData if needed
		if len(r.MetaData) != 0 {
			m.message.RecipientMetaData = append(m.message.RecipientMetaData, RecipientMetaData{
				Rcpt:   r.Email,
				Values: r.MetaData,
			})
		}
	}
}

// Send sends the current instances of the MessageBuilder
func (m *MessageBuilder) Send() ([]MessageStatus, error) {
	return m.SendContext(context.TODO())
}

// SendContext sends the current instances of the MessageBuilder
func (m *MessageBuilder) SendContext(ctx context.Context) ([]MessageStatus, error) {
	// This is somewhat unused right now
	// The future idea is an additional immutability check where we only send finalized messages
	if !m.finalized {
		m.finalize()
	}
	if m.isTemplate {
		t := TemplateMessage{
			TemplateName:    m.templateName,
			TemplateContent: m.templateContent,
			Message:         *m.message,
		}
		return globalClient.SendTemplateContext(ctx, t)
	}
	return globalClient.SendMessageContext(ctx, *m.message)
}
