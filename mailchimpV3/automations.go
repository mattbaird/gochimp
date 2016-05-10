package mailchimpV3

import (
	"errors"
	"fmt"
)

const (
	automations_path       = "/automations"
	single_automation_path = automations_path + "/%s"
	pause_all_emails_path  = single_automation_path + "/actions/pause-all-emails"
	start_all_emails_path  = single_automation_path + "/actions/start-all-emails"

	automation_email_path        = single_automation_path + "/emails"
	single_automation_email_path = automation_email_path + "/%s"

	pause_single_email_path = single_automation_email_path + "/actions/pause"
	start_single_email_path = single_automation_email_path + "/actions/start"

	automation_queues_path       = single_automation_email_path + "/queue"
	single_automation_queue_path = automation_queues_path + "/%s"

	removed_subscribers_automation_path = single_automation_path + "/removed-subscribers"
)

type ListOfAutomations struct {
	baseList
	Automations []Automation `json:"automations"`
}

type Automation struct {
	ID              string                  `json:"id"`
	CreateTime      string                  `json:"create_time"`
	StartTime       string                  `json:"start_time"`
	Status          string                  `json:"status"`
	EmailsSent      int                     `json:"emails_sent"`
	Recipients      AutomationRecipient     `json:"recipients"`
	Settings        AutomationSettingsShort `json:"settings"`
	Tracking        AutomationTracking      `json:"tracking"`
	TriggerSettings WorkflowType            `json:"trigger_settings"`
	ReportSummary   ReportSummary           `json:"report_summary"`

	withLinks
	api *ChimpAPI
}

type AutomationRecipient struct {
	ListID         string            `json:"list_id"`
	SegmentOptions AutomationOptions `json:"segment_options"`
}

type AutomationOptions struct {
	SavedSegmentID int                  `json:"saved_segment_id"`
	Match          string               `json:"match"`
	Conditions     []SegmentConditional `json:"conditions"`
}

type AutomationSettingsShort struct {
	UseConversation bool   `json:"use_conversation"`
	ToName          string `json:"to_name"`
	Title           string `json:"title"`
	FromName        string `json:"from_name"`
	ReplyTo         string `json:"reply_to"`
	Authenticate    bool   `json:"authenticate"`
	AutoFooter      bool   `json:"auto_footer"`
	InlineCSS       bool   `json:"inline_css"`
}

type AutomationSettingsLong struct {
	Title        string   `json:"title"`
	FromName     string   `json:"from_name"`
	ReplyTo      string   `json:"reply_to"`
	Authenticate bool     `json:"authenticate"`
	AutoFooter   bool     `json:"auto_footer"`
	InlineCSS    bool     `json:"inline_css"`
	SubjectLine  string   `json:"subject_line"`
	AutoTweet    bool     `json:"auto_tweet"`
	AutoFBPost   []string `json:"auto_fb_post"`
	FBComments   bool     `json:"fb_comments"`
	TemplateID   int      `json:"template_id"`
	DragAndDrop  bool     `json:"drag_and_drop"`
}

type AutomationTracking struct {
	Opens           bool       `json:"opens"`
	HTMLClicks      bool       `json:"html_clicks"`
	TextClicks      bool       `json:"text_clicks"`
	GoalTracking    bool       `json:"goal_tracking"`
	Ecomm360        bool       `json:"ecomm360"`
	GoogleAnalytics string     `json:"google_analytics"`
	Clicktale       string     `json:"clicktale"`
	Salesforce      Salesforce `json:"salesforce"`
	Highrise        Highrise   `json:"highrise"`
	Capsule         Capsule    `json:"capsule"`
}

type Salesforce struct {
	Campaign bool `json:"campaign"`
	Notes    bool `json:"notes"`
}

type Highrise struct {
	Campaign bool `json:"campaign"`
	Notes    bool `json:"notes"`
}

type Capsule struct {
	Notes bool `json:"notes"`
}

type ReportSummary struct {
	Opens            int     `json:"opens"`
	UniqueOpens      int     `json:"unique_opens"`
	OpenRate         float64 `json:"open_rate"`
	Clicks           int     `json:"clicks"`
	SubscriberClicks int     `json:"subscriber_clicks"`
	ClickRate        float64 `json:"click_rate"`
}

func (auto Automation) CanMakeRequest() error {
	if auto.ID == "" {
		return errors.New("No ID provided")
	}

	return nil
}

func (api ChimpAPI) GetAutomations(params *BasicQueryParams) (*ListOfAutomations, error) {
	response := new(ListOfAutomations)

	err := api.Request("GET", automations_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, l := range response.Automations {
		l.api = &api
	}

	return response, nil
}

// TODO query params?
func (api ChimpAPI) GetAutomation(id string) (*Automation, error) {
	endpoint := fmt.Sprintf(single_automation_path, id)

	response := new(Automation)
	response.api = &api

	return response, api.Request("GET", endpoint, nil, nil, response)
}

// ------------------------------------------------------------------------------------------------
// Actions for Sending Emails
// ------------------------------------------------------------------------------------------------

func (auto Automation) PauseSendingAll() (bool, error) {
	if err := auto.CanMakeRequest(); err != nil {
		return false, err
	}
	return auto.api.PauseSendingAll(auto.ID)
}

func (api ChimpAPI) PauseSendingAll(id string) (bool, error) {
	endpoint := fmt.Sprintf(pause_all_emails_path, id)
	return api.Do("POST", endpoint)
}

func (auto Automation) StartSendingAll() (bool, error) {
	if err := auto.CanMakeRequest(); err != nil {
		return false, err
	}
	return auto.api.StartSendingAll(auto.ID)
}

func (api ChimpAPI) StartSendingAll(id string) (bool, error) {
	endpoint := fmt.Sprintf(start_all_emails_path, id)
	return api.Do("POST", endpoint)
}

func (email AutomationEmail) PauseSending() (bool, error) {
	return email.api.PauseSending(email.WorkflowID, email.ID)
}

func (api ChimpAPI) PauseSending(workflowID, emailID string) (bool, error) {
	endpoint := fmt.Sprintf(pause_single_email_path, workflowID, emailID)
	return api.Do("POST", endpoint)
}

func (email AutomationEmail) StartSending() (bool, error) {
	return email.api.StartSending(email.WorkflowID, email.ID)
}

func (api ChimpAPI) StartSending(workflowID, emailID string) (bool, error) {
	endpoint := fmt.Sprintf(start_single_email_path, workflowID, emailID)
	return api.Do("POST", endpoint)
}

// ------------------------------------------------------------------------------------------------
// Automation Emails
// ------------------------------------------------------------------------------------------------

type ListOfEmails struct {
	baseList
	Emails []AutomationEmail `json:"emails"`
}

type AutomationEmail struct {
	ID              string                 `json:"id"`
	WorkflowID      string                 `json:"workflow_id"`
	Position        int                    `json:"position"`
	Delay           AutomationDelay        `json:"delay"`
	CreateTime      string                 `json:"create_time"`
	StartTime       string                 `json:"start_time"`
	ArchiveURL      string                 `json:"archive_url"`
	Status          string                 `json:"status"`
	EmailsSent      int                    `json:"emails_sent"`
	SendTime        string                 `json:"send_time"`
	ContentType     string                 `json:"content_type"`
	Recipients      AutomationRecipient    `json:"recipients"`
	Settings        AutomationSettingsLong `json:"settings"`
	Tracking        AutomationTracking     `json:"tracking"`
	SocialCard      SocialCard             `json:"social_card"`
	TriggerSettings WorkflowType           `json:"trigger_settings"`
	ReportSummary   ReportSummary          `json:"report_summary"`

	withLinks
	api *ChimpAPI
}

type SocialCard struct {
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

type AutomationDelay struct {
	Amount    int    `json:"amount"`
	Type      string `json:"type"`
	Direction string `json:"direction"`
	Action    string `json:"action"`
}

func (email AutomationEmail) CanMakeRequest() error {
	if email.ID == "" {
		return errors.New("No ID provided")
	}

	return nil
}

func (auto Automation) GetEmails() (*ListOfEmails, error) {
	if err := auto.CanMakeRequest(); err != nil {
		return nil, err
	}

	return auto.api.GetAutomationEmails(auto.ID)
}

func (api ChimpAPI) GetAutomationEmails(automationID string) (*ListOfEmails, error) {
	endpoint := fmt.Sprintf(automation_email_path, automationID)
	response := new(ListOfEmails)

	for _, l := range response.Emails {
		l.api = &api
	}

	return response, api.Request("GET", endpoint, nil, nil, response)
}

func (auto Automation) GetEmail(id string) (*AutomationEmail, error) {
	if err := auto.CanMakeRequest(); err != nil {
		return nil, err
	}

	return auto.api.GetAutomationEmail(auto.ID, id)
}

func (api ChimpAPI) GetAutomationEmail(automationID, emailID string) (*AutomationEmail, error) {
	endpoint := fmt.Sprintf(single_automation_email_path, automationID, emailID)
	response := new(AutomationEmail)
	response.api = &api

	return response, api.Request("GET", endpoint, nil, nil, response)
}

// ------------------------------------------------------------------------------------------------
// Queues
// ------------------------------------------------------------------------------------------------

type AutomationQueueRequest struct {
	EmailAddress string `json:"email_address"`
}

type ListOfAutomationQueues struct {
	baseList
	WorkflowID string            `json:"workflow_id"`
	EmailID    string            `json:"email_id"`
	Queues     []AutomationQueue `json:"queue"`
}

type AutomationQueue struct {
	ID           string `json:"id"`
	WorkflowID   string `json:"workflow_id"`
	EmailID      string `json:"email_id"`
	ListID       string `json:"list_id"`
	EmailAddress string `json:"email_address"`
	NextSend     string `json:"next_send"`
	withLinks

	api *ChimpAPI
}

func (email AutomationEmail) GetQueues() (*ListOfAutomationQueues, error) {
	if err := email.CanMakeRequest(); err != nil {
		return nil, err
	}

	return email.api.GetAutomationQueues(email.WorkflowID, email.ID)
}

func (api ChimpAPI) GetAutomationQueues(workflowID, emailID string) (*ListOfAutomationQueues, error) {
	endpoint := fmt.Sprintf(automation_queues_path, workflowID, emailID)

	response := new(ListOfAutomationQueues)
	for _, l := range response.Queues {
		l.api = &api
	}

	return response, api.Request("GET", endpoint, nil, nil, response)
}

func (email AutomationEmail) GetQueue(id string) (*AutomationQueue, error) {
	if err := email.CanMakeRequest(); err != nil {
		return nil, err
	}

	return email.api.GetAutomationQueue(email.WorkflowID, email.ID, id)
}

func (api ChimpAPI) GetAutomationQueue(workflowID, emailID, subsID string) (*AutomationQueue, error) {
	endpoint := fmt.Sprintf(single_automation_queue_path, workflowID, emailID, subsID)

	response := new(AutomationQueue)
	response.api = &api

	return response, api.Request("GET", endpoint, nil, nil, response)
}

func (email AutomationEmail) CreateQueue(emailAddress string) (*AutomationQueue, error) {
	if err := email.CanMakeRequest(); err != nil {
		return nil, err
	}

	return email.api.CreateAutomationEmailQueue(email.WorkflowID, email.ID, emailAddress)
}

func (api ChimpAPI) CreateAutomationEmailQueue(workflowID, emailID, emailAddress string) (*AutomationQueue, error) {
	endpoint := fmt.Sprintf(automation_queues_path, workflowID, emailID)
	response := new(AutomationQueue)

	body := &AutomationQueueRequest{
		EmailAddress: emailAddress,
	}

	err := api.Request("POST", endpoint, nil, body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ------------------------------------------------------------------------------------------------
// Removed Subscribers
// ------------------------------------------------------------------------------------------------

type RemovedSubscriberRequest struct {
	EmailAddress string `json:"email_address"`
}

type ListOfRemovedSubscribers struct {
	baseList
	WorkflowID  string              `json:"workflow_id"`
	Subscribers []RemovedSubscriber `json:"subscribers"`
}

type RemovedSubscriber struct {
	ID           string `json:"id"`
	WorkflowID   string `json:"workflow_id"`
	ListID       string `json:"list_id"`
	EmailAddress string `json:"email_address"`

	withLinks
}

func (auto Automation) GetRemovedSubscribers() (*ListOfRemovedSubscribers, error) {
	if err := auto.CanMakeRequest(); err != nil {
		return nil, err
	}

	return auto.api.GetAutomationRemovedSubscribers(auto.ID)
}

func (api ChimpAPI) GetAutomationRemovedSubscribers(workflowID string) (*ListOfRemovedSubscribers, error) {
	endpoint := fmt.Sprintf(removed_subscribers_automation_path, workflowID)

	response := new(ListOfRemovedSubscribers)

	return response, api.Request("GET", endpoint, nil, nil, response)
}

func (auto Automation) CreateRemovedSubscribers(emailAddress string) (*RemovedSubscriber, error) {
	if err := auto.CanMakeRequest(); err != nil {
		return nil, err
	}

	return auto.api.CreateAutomationRemovedSubscribers(auto.ID, emailAddress)
}

func (api ChimpAPI) CreateAutomationRemovedSubscribers(workflowID, emailAddress string) (*RemovedSubscriber, error) {
	endpoint := fmt.Sprintf(removed_subscribers_automation_path, workflowID)

	response := new(RemovedSubscriber)
	body := &RemovedSubscriberRequest{
		EmailAddress: emailAddress,
	}

	return response, api.Request("POST", endpoint, nil, body, response)
}
