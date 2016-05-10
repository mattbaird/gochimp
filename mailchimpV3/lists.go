package mailchimpV3

import (
	"errors"
	"fmt"
)

const (
	lists_path       = "/lists"
	single_list_path = lists_path + "/%s"

	abuse_reports_path       = "/lists/%s/abuse_reports"
	single_abuse_report_path = abuse_reports_path + "/%s"

	activity_path = "/lists/%s/activity"
	clients_path  = "/lists/%s/clients"

	history_path        = "/lists/%s/growth-history"
	single_history_path = history_path + "/%s"

	interest_categories_path      = "/lists/%s/interest-categories"
	single_interest_category_path = interest_categories_path + "/%s"
)

type ListQueryParams struct {
	ExtendedQueryParams

	BeforeDateCreated      string
	SinceDateCreated       string
	BeforeCampaignLastSent string
	SinceCampaignLastSent  string
	Email                  string
}

func (q ListQueryParams) Params() map[string]string {
	m := q.ExtendedQueryParams.Params()
	m["before_date_created"] = q.BeforeDateCreated
	m["since_date_created"] = q.SinceDateCreated
	m["before_campaign_last_sent"] = q.BeforeCampaignLastSent
	m["since_campaign_last_sent"] = q.SinceCampaignLastSent
	m["email"] = q.Email
	return m
}

type ListOfLists struct {
	baseList
	Lists []ListResponse `json:"lists"`
}

type ListCreationRequest struct {
	Name                string           `json:"name"`
	Contact             Contact          `json:"contact"`
	PermissionReminder  string           `json:"permission_reminder"`
	UseArchiveBar       bool             `json:"use_archive_bar"`
	CampaignDefaults    CampaignDefaults `json:"campaign_defaults"`
	NotifyOnSubscribe   string           `json:"notify_on_subscribe"`
	NotifyOnUnsubscribe string           `json:"notify_on_unsubscribe"`
	EmailTypeOption     bool             `json:"email_type_option"`
	Visibilty           string           `json:"visiblity"`
}

type ListResponse struct {
	ListCreationRequest
	withLinks

	ID                string                   `json:"id"`
	DateCreated       string                   `json:"date_created"`
	ListRating        int                      `json:"list_rating"`
	SubscribeURLShort string                   `json:"subscribe_url_short"`
	SubscribeURLLong  string                   `json:"subscribe_url_long"`
	BeamerAddress     string                   `json:"beamer_address"`
	Modules           []map[string]interface{} `json:"modules"` // TODO undocumented
	Stats             Stats                    `json:"stats"`

	api *ChimpAPI
}

func (list ListResponse) CanMakeRequest() error {
	if list.ID == "" {
		return errors.New("No ID provided on list")
	}

	return nil
}

type Stats struct {
	MemberCount               int     `json:"member_count"`
	UnsubscribeCount          int     `json:"unsubscribe_count"`
	CleanedCount              int     `json:"cleaned_count"`
	MemberCountSinceSend      int     `json:"member_count_since_send"`
	UnsubscribeCountSinceSend int     `json:"unsubscribe_count_since_send"`
	CleanedCountSinceSend     int     `json:"cleaned_count_since_send"`
	CampaignCount             int     `json:"campaign_count"`
	CampaignLastSent          string  `json:"campaign_last_sent"`
	MergeFieldCount           int     `json:"merge_field_count"`
	AvgSubRate                float64 `json:"avg_sub_rate"`
	AvgUnsubRate              float64 `json:"avg_unsub_rate"`
	TargetSubRate             float64 `json:"target_sub_rate"`
	OpenRate                  float64 `json:"open_rate"`
	ClickRate                 float64 `json:"click_rate"`
	LastSubDate               string  `json:"last_sub_date"`
	LastUnsubDate             string  `json:"last_unsub_date"`
}

type CampaignDefaults struct {
	FromName  string `json:"from_name"`
	FromEmail string `json:"from_email"`
	Subject   string `json:"subject"`
	Language  string `json:"language"`
}

func (api ChimpAPI) GetLists(params *ListQueryParams) (*ListOfLists, error) {
	response := new(ListOfLists)

	err := api.Request("GET", lists_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, l := range response.Lists {
		l.api = &api
	}

	return response, nil
}

func (api ChimpAPI) GetList(id string, params *BasicQueryParams) (*ListResponse, error) {
	endpoint := fmt.Sprintf(single_list_path, id)

	response := new(ListResponse)
	response.api = &api

	return response, api.Request("GET", endpoint, params, nil, response)
}

func (api ChimpAPI) CreateList(body *ListCreationRequest) (*ListResponse, error) {
	response := new(ListResponse)
	response.api = &api
	return response, api.Request("POST", lists_path, nil, body, response)
}

func (api ChimpAPI) UpdateList(id string, body *ListCreationRequest) (*ListResponse, error) {
	endpoint := fmt.Sprintf(single_list_path, id)

	response := new(ListResponse)
	response.api = &api

	return response, api.Request("PATCH", endpoint, nil, body, response)
}

func (api ChimpAPI) DeleteList(id string) (bool, error) {
	endpoint := fmt.Sprintf(single_list_path, id)
	return api.Do("DELETE", endpoint)
}

// ------------------------------------------------------------------------------------------------
// Abuse Reports
// ------------------------------------------------------------------------------------------------

type ListOfAbuseReports struct {
	baseList

	ListID  string        `json:"list_id"`
	Reports []AbuseReport `json:"abuse_reports"`
}

type AbuseReport struct {
	ID           string `json:"id"`
	CampaignID   string `json:"campaign_id"`
	ListID       string `json:"list_id"`
	EmailID      string `json:"email_id"`
	EmailAddress string `json:"email_address"`
	Date         string `json:"date"`

	withLinks
}

func (list ListResponse) GetAbuseReports(params *ExtendedQueryParams) (*ListOfAbuseReports, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(abuse_reports_path, list.ID)
	response := new(ListOfAbuseReports)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

func (list ListResponse) GetAbuseReport(id string, params *ExtendedQueryParams) (*AbuseReport, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(single_abuse_report_path, list.ID, id)
	response := new(AbuseReport)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

// ------------------------------------------------------------------------------------------------
// Activity
// ------------------------------------------------------------------------------------------------

type ListOfActivity struct {
	baseList

	ListID     string     `json:"list_id"`
	Activities []Activity `json:"activity"`
}

type Activity struct {
	Day             string `json:"day"`
	EmailsSent      int    `json:"emails_sent"`
	UniqueOpens     int    `json:"unique_opens"`
	RecipientClicks int    `json:"recipient_clicks"`
	HardBounce      int    `json:"hard_bounce"`
	SoftBounce      int    `json:"soft_bounce"`
	Subs            int    `json:"subs"`
	Unsubs          int    `json:"unsubs"`
	OtherAdds       int    `json:"other_adds"`
	OtherRemoves    int    `json:"other_removes"`

	withLinks
}

func (list ListResponse) GetActivity(params *BasicQueryParams) (*ListOfActivity, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(activity_path, list.ID)
	response := new(ListOfActivity)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

// ------------------------------------------------------------------------------------------------
// Clients
// ------------------------------------------------------------------------------------------------

type ListOfClients struct {
	baseList

	ListID  string   `json:"list_id"`
	Clients []Client `json:"clients"`
}

type Client struct {
	Client  string `json:"client"`
	Members int    `json:"members"`
	ListID  string `json:"list_id"`

	withLinks
}

func (list ListResponse) GetClients(params *BasicQueryParams) (*ListOfClients, error) {
	if list.ID == "" {
		return nil, errors.New("No ID provided on list")
	}

	endpoint := fmt.Sprintf(clients_path, list.ID)
	response := new(ListOfClients)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

// ------------------------------------------------------------------------------------------------
// Growth History
// ------------------------------------------------------------------------------------------------

type ListOfGrownHistory struct {
	baseList

	ListID  string          `json:"list_id"`
	History []GrowthHistory `json:"history"`
}

type GrowthHistory struct {
	ListID   string `json:"list_id"`
	Month    string `json:"month"`
	Existing int    `json:"existing"`
	Imports  int    `json:"imports"`
	OptIns   int    `json:"optins"`

	withLinks
}

func (list ListResponse) GetGrowthHistory(params *ExtendedQueryParams) (*ListOfGrownHistory, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(history_path, list.ID)
	response := new(ListOfGrownHistory)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

func (list ListResponse) GetGrowthHistoryForMonth(month string, params *BasicQueryParams) (*GrowthHistory, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(single_history_path, list.ID, month)
	response := new(GrowthHistory)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

// ------------------------------------------------------------------------------------------------
// Interest Categories
// ------------------------------------------------------------------------------------------------

type ListOfInterestCategories struct {
	baseList
	ListID     string             `json:"list_id"`
	Categories []InterestCategory `json:"categories"`
}

type InterestCategoryRequest struct {
	Title        string `json:"title"`
	DisplayOrder int    `json:"display_order"`
	Type         string `json:"type"`
}

type InterestCategory struct {
	InterestCategoryRequest

	ListID string `json:"list_id"`
	ID     string `json:"id"`

	withLinks
}

type InterestCategoriesQueryParams struct {
	ExtendedQueryParams

	Type string `json:"type"`
}

func (q InterestCategoriesQueryParams) Params() map[string]string {
	m := q.ExtendedQueryParams.Params()
	m["type"] = q.Type
	return m
}

func (list ListResponse) GetInterestCategories(params *InterestCategoriesQueryParams) (*ListOfInterestCategories, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(interest_categories_path, list.ID)
	response := new(ListOfInterestCategories)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

func (list ListResponse) GetInterestCategory(id string, params *BasicQueryParams) (*InterestCategory, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(single_interest_category_path, list.ID, id)
	response := new(InterestCategory)

	return response, list.api.Request("GET", endpoint, params, nil, response)
}

func (list ListResponse) CreateInterestCategory(body *InterestCategoryRequest) (*InterestCategory, error) {
	if err := list.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(interest_categories_path, list.ID)
	response := new(InterestCategory)

	return response, list.api.Request("POST", endpoint, nil, body, response)
}

func (list ListResponse) UpdateInterestCategory(id string, body *InterestCategoryRequest) (*InterestCategory, error) {
	if list.ID == "" {
		return nil, errors.New("No ID provided on list")
	}

	endpoint := fmt.Sprintf(single_interest_category_path, list.ID, id)
	response := new(InterestCategory)

	return response, list.api.Request("PATCH", endpoint, nil, body, response)
}

func (list ListResponse) DeleteInterestCategory(id string) (bool, error) {
	if list.ID == "" {
		return false, errors.New("No ID provided on list")
	}

	endpoint := fmt.Sprintf(single_interest_category_path, list.ID, id)
	return list.api.Do("DELETE", endpoint)
}

// ------------------------------------------------------------------------------------------------
// Merge Fields
// ------------------------------------------------------------------------------------------------
// TODO
