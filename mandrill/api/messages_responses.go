package api

type MessageSendReponse struct {
	Email        string `json:"email"`
	Status       string `json:"status"`
	RejectReason string `json:"reject_reason"`
	ID           string `json:"_id"`
}

type MessagesSendResponse []MessageSendReponse

type MessagesSendTemplateResponse []MessageSendReponse

type MessagesInfoResponse struct {
	TS          int32    `json:"ts"`
	ID          string   `json:"_id"`
	Sender      string   `json:"sender"`
	Template    string   `json:"template"`
	Subject     string   `json:"subject"`
	Email       string   `json:"email"`
	Tags        []string `json:"tags"`
	Opens       int      `json:"opens"`
	OpensDetail []struct {
		TS       int32  `json:"ts"`
		IP       string `json:"ip"`
		Location string `json:"location"`
		UA       string `json:"ua"`
	} `json:"opens_detail"`
	Clicks       int `json:"clicks"`
	ClicksDetail []struct {
		TS       int32  `json:"ts"`
		IP       string `json:"ip"`
		Location string `json:"location"`
		UA       string `json:"ua"`
	} `json:"clicks_detail"`
	State      string            `json:"state"`
	MetaData   map[string]string `json:"metadata"`
	SMTPEvents []struct {
		TS   int32  `json:"ts"`
		Type string `json:"type"`
		Diag string `json:"diag"`
	} `json:"smtp_events"`
}

type MessagesSearchResponse []MessagesInfoResponse

type MessageTimeSeriesSearchResponse struct {
	Time         Time `json:"apitime"`
	Sent         int  `json:"sent"`
	HardBounces  int  `json:"hard_bounces"`
	SoftBounces  int  `json:"soft_bounces"`
	Rejects      int  `json:"rejects"`
	Complaints   int  `json:"complaints"`
	Unsubs       int  `json:"unsubs"`
	Opens        int  `json:"opens"`
	UniqueOpens  int  `json:"unique_opens"`
	Clicks       int  `json:"clicks"`
	UniqueClicks int  `json:"unique_clicks"`
}
type MessagesSearchTimeSeriesResponse []MessageTimeSeriesSearchResponse

type MessagesContentResponse struct {
	TS        int32  `json:"ts"`
	ID        string `json:"_id"`
	FromEmail string `json:"from_email"`
	FromName  string `json:"from_name"`
	Subject   string `json:"subject"`
	To        struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"to"`
	Tags        []string          `json:"tags"`
	Headers     map[string]string `json:"headers"`
	Text        string            `json:"text"`
	HTML        string            `json:"html"`
	Attachments []struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Content string `json:"content"`
	} `json:"attachments"`
}

type MessagesParseResponse struct {
	TS        int32  `json:"ts"`
	ID        string `json:"_id"`
	FromEmail string `json:"from_email"`
	FromName  string `json:"from_name"`
	Subject   string `json:"subject"`
	To        struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"to"`
	Tags        []string          `json:"tags"`
	Headers     map[string]string `json:"headers"`
	Text        string            `json:"text"`
	HTML        string            `json:"html"`
	Attachments []struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Content string `json:"content"`
	} `json:"attachments"`
	Images []struct {
		Name    string `json:"name"`
		Type    string `json:"type"`
		Content string `json:"content"`
	} `json:"images"`
}

type MessageSendRawResponse struct {
	Email        string `json:"email"`
	Status       string `json:"status"`
	RejectReason string `json:"reject_reason"`
	ID           string `json:"_id"`
}

type MessagesSendRawResponse []MessageSendRawResponse

type MessagesScheduledMessageResponse struct {
	ID        string `json:"_id"`
	CreatedAt Time   `json:"created_at"`
	SendAt    Time   `json:"send_at"`
	FromEmail string `json:"from_email"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
}

type MessagesListScheduledResponse []MessagesScheduledMessageResponse

type MessagesCancelScheduledResponse MessagesScheduledMessageResponse

type MessagesRescheduleResponse MessagesScheduledMessageResponse
