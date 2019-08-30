package api

type MessageRequest struct {
	HTML                    string                     `json:"html,omitempty"`
	Text                    string                     `json:"text,omitempty"`
	Subject                 string                     `json:"subject,omitempty"`
	FromEmail               string                     `json:"from_email"`
	FromName                string                     `json:"from_name,omitempty"`
	To                      []MessageTo                `json:"to"`
	Headers                 map[string]string          `json:"headers,omitempty"`
	Important               bool                       `json:"important,omitempty"`
	TrackOpens              bool                       `json:"track_opens,omitempty"`
	TrackClicks             bool                       `json:"track_clicks,omitempty"`
	AutoText                bool                       `json:"auto_text,omitempty"`
	AutoHTML                bool                       `json:"auto_html,omitempty"`
	InlineCSS               bool                       `json:"inline_css,omitempty"`
	URLStripQS              bool                       `json:"url_strip_qs,omitempty"`
	PreserveRecipients      bool                       `json:"preserve_recipients,omitempty"`
	ViewContentLink         bool                       `json:"view_content_link,omitempty"`
	BccAddress              string                     `json:"bcc_address,omitempty"`
	TrackingDomain          string                     `json:"tracking_domain,omitempty"`
	SigningDomain           string                     `json:"signing_domain,omitempty"`
	ReturnPathDomain        string                     `json:"return_path_domain,omitempty"`
	Merge                   bool                       `json:"merge,omitempty"`
	MergeLanguage           string                     `json:"merge_language,omitempty"`
	GlobalMergeVars         []MessageVar               `json:"global_merge_vars,omitempty"`
	MergeVars               []MessageMergeVar          `json:"merge_vars,omitempty"`
	Tags                    []string                   `json:"tags,omitempty"`
	SubAccount              string                     `json:"subaccount,omitempty"`
	GoogleAnalyticsDomains  []string                   `json:"google_analytics_domains,omitempty"`
	GoogleAnalyticsCampaign []string                   `json:"google_analytics_campaign,omitempty"`
	MetaData                map[string]string          `json:"metadata,omitempty"`
	RecipientMetaData       []MessageRecipientMetaData `json:"recipient_metadata,omitempty"`
	Attachments             []MessageAttachment        `json:"attachments,omitempty"`
	Images                  []MessageAttachment        `json:"images,omitempty"`
	Async                   bool                       `json:"async,omitempty"`
	IPPool                  string                     `json:"ip_pool,omitempty"`
	SentAt                  string                     `json:"send_at,omitempty"`
}

type MessageTo struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
}
type MessageMergeVar struct {
	Rcpt string       `json:"rcpt"`
	Vars []MessageVar `json:"vars,omitempty"`
}
type MessageVar struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type MessageRecipientMetaData struct {
	Rcpt   string            `json:"rcpt"`
	Values map[string]string `json:"values"`
}

type MessageAttachment struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type MessagesSendRequest struct {
	Key     string         `json:"key"`
	Message MessageRequest `json:"message"`
}

type MessagesSendTemplateRequest struct {
	Key             string                    `json:"key"`
	TemplateName    string                    `json:"template_name"`
	TemplateContent []MessagesTemplateContent `json:"template_content"`
	Message         MessageRequest            `json:"message"`
}

type MessagesTemplateContent struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type MessagesSearchRequest struct {
	Key      string   `json:"key"`
	Query    string   `json:"query,omitempty"`
	DateFrom string   `json:"date_from,omitempty"`
	DateTo   string   `json:"date_to,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Senders  []string `json:"senders,omitempty"`
	APIKeys  []string `json:"apikeys,omitempty"`
	Limit    int      `json:"limit,omitempty"`
}

type MessagesSearchTimeSeriesRequest struct {
	Key      string   `json:"key"`
	Query    string   `json:"query,omitempty"`
	DateFrom string   `json:"date_from,omitempty"`
	DateTo   string   `json:"date_to,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Senders  []string `json:"senders,omitempty"`
	APIKeys  []string `json:"apikeys,omitempty"`
	Limit    int      `json:"limit,omitempty"`
}

type MessagesInfoRequest struct {
	Key string `json:"key"`
	ID  string `json:"id"`
}

type MessagesContentRequest struct {
	Key string `json:"key"`
	ID  string `json:"id"`
}

type MessagesParseRequest struct {
	Key        string `json:"key"`
	RawMessage string `json:"raw_message"`
}

type MessagesSendRawRequest struct {
	Key              string   `json:"key"`
	RawMessage       string   `json:"raw_message"`
	FromEmail        string   `json:"from_email,omitempty"`
	FromName         string   `json:"from_name,omitempty"`
	To               []string `json:"to,omitempty"`
	Async            bool     `json:"async,omitempty"`
	IPPool           string   `json:"ip_pool,omitempty"`
	SendAt           string   `json:"send_at,omitempty"`
	ReturnPathDomain string   `json:"return_path_domain,omitempty"`
}

type MessagesListScheduledRequest struct {
	Key string `json:"key"`
	To  string `json:"to,omitempty"`
}

type MessagesCancelScheduledRequest struct {
	Key string `json:"key"`
	ID  string `json:"id"`
}

type MessagesRescheduleRequest struct {
	Key    string `json:"key"`
	ID     string `json:"id"`
	SendAt string `json:"send_at"`
}
