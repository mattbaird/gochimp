// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gochimp

// see http://apidocs.mailchimp.com/api/2.0/
const (
	lists_subscribe_endpoint         string = "/lists/subscribe.json"
	lists_unsubscribe_endpoint       string = "/lists/unsubscribe.json"
	lists_list_endpoint              string = "/lists/list.json"
	lists_update_member_endpoint     string = "/lists/update-member.json"
	lists_members_endpoint           string = "/lists/members.json"
	lists_member_info_endpoint       string = "/lists/member-info.json"
	lists_batch_unsubscribe_endpoint string = "/lists/batch-unsubscribe.json"
	lists_batch_subscribe_endpoint   string = "/lists/batch-subscribe.json"
)

func (a *ChimpAPI) BatchSubscribe(req BatchSubscribe) (BatchSubscribeResponse, error) {
	var response BatchSubscribeResponse
	req.ApiKey = a.Key
	err := parseChimpJson(a, lists_batch_subscribe_endpoint, req, &response)
	return response, err
}

func (a *ChimpAPI) BatchUnsubscribe(req BatchUnsubscribe) (BatchResponse, error) {
	var response BatchResponse
	req.ApiKey = a.Key
	err := parseChimpJson(a, lists_batch_unsubscribe_endpoint, req, &response)
	return response, err
}

func (a *ChimpAPI) ListsSubscribe(req ListsSubscribe) (Email, error) {
	var response Email
	req.ApiKey = a.Key
	err := parseChimpJson(a, lists_subscribe_endpoint, req, &response)
	return response, err
}

func (a *ChimpAPI) ListsUnsubscribe(req ListsUnsubscribe) error {
	req.ApiKey = a.Key
	return parseChimpJson(a, lists_unsubscribe_endpoint, req, nil)
}

func (a *ChimpAPI) ListsList(req ListsList) (ListsListResponse, error) {
	req.ApiKey = a.Key
	var response ListsListResponse
	err := parseChimpJson(a, lists_list_endpoint, req, &response)
	return response, err
}

func (a *ChimpAPI) UpdateMember(req UpdateMember) error {
	req.ApiKey = a.Key
	return parseChimpJson(a, lists_update_member_endpoint, req, nil)
}

func (a *ChimpAPI) Members(req ListsMembers) (ListsMembersResponse, error) {
	req.ApiKey = a.Key
	var response ListsMembersResponse
	err := parseChimpJson(a, lists_members_endpoint, req, &response)
	return response, err
}

func (a *ChimpAPI) MemberInfo(req ListsMemberInfo) (ListsMemberInfoResponse, error) {
	req.ApiKey = a.Key
	var response ListsMemberInfoResponse
	err := parseChimpJson(a, lists_member_info_endpoint, req, &response)
	return response, err
}

type BatchUnsubscribe struct {
	ApiKey       string  `json:"apikey"`
	ListId       string  `json:"id"`
	Batch        []Email `json:"batch"`
	DeleteMember bool    `json:"delete_member"`
	SendGoodbye  bool    `json:"send_goodbye"`
	SendNotify   bool    `json:"send_notify"`
}

type BatchSubscribe struct {
	ApiKey           string        `json:"apikey"`
	ListId           string        `json:"id"`
	Batch            []ListsMember `json:"batch"`
	DoubleOptin      bool          `json:"double_optin"`
	UpdateExisting   bool          `json:"update_existing"`
	ReplaceInterests bool          `json:"replace_interests"`
}

type ListsUnsubscribe struct {
	ApiKey       string `json:"apikey"`
	ListId       string `json:"id"`
	Email        Email  `json:"email"`
	DeleteMember bool   `json:"delete_member"`
	SendGoodbye  bool   `json:"send_goodbye"`
	SendNotify   bool   `json:"send_notify"`
}

type ListsSubscribe struct {
	ApiKey           string                 `json:"apikey"`
	ListId           string                 `json:"id"`
	Email            Email                  `json:"email"`
	MergeVars        map[string]interface{} `json:"merge_vars,omitempty"`
	EmailType        string                 `json:"email_type,omitempty"`
	DoubleOptIn      bool                   `json:"double_optin"`
	UpdateExisting   bool                   `json:"update_existing"`
	ReplaceInterests bool                   `json:"replace_interests"`
	SendWelcome      bool                   `json:"send_welcome"`
}

type ListFilter struct {
	ListId        string `json:"list_id"`
	ListName      string `json:"list_name"`
	FromName      string `json:"from_name"`
	FromEmail     string `json:"from_email"`
	FromSubject   string `json:"from_subject"`
	CreatedBefore string `json:"created_before"`
	CreatedAfter  string `json:"created_after"`
	Exact         bool   `json:"exact"`
}

type ListStat struct {
	MemberCount               float64 `json:"member_count"`
	UnsubscribeCount          float64 `json:"unsubscribe_count"`
	CleanedCount              float64 `json:"cleaned_count"`
	MemberCountSinceSend      float64 `json:"member_count_since_send"`
	UnsubscribeCountSinceSend float64 `json:"unsubscribe_count_since_send"`
	CleanedCountSinceSend     float64 `json:"cleaned_count_since_send"`
	CampaignCount             float64 `json:"campaign_count"`
	GroupingCount             float64 `json:"grouping_count"`
	GroupCount                float64 `json:"group_count"`
	MergeVarCount             float64 `json:"merge_var_count"`
	AvgSubRate                float64 `json:"avg_sub_rate"`
	AvgUnsubRate              float64 `json:"avg_unsub_rate"`
	TargetSubRate             float64 `json:"target_sub_rate"`
	OpenRate                  float64 `json:"open_rate"`
	ClickRate                 float64 `json:"click_rate"`
}

type ListData struct {
	Id                string   `json:"id"`
	WebId             int      `json:"web_id"`
	Name              string   `json:"name"`
	DateCreated       string   `json:"date_created"`
	EmailTypeOption   bool     `json:"email_type_option"`
	UseAwesomeBar     bool     `json:"use_awesomebar"`
	DefaultFromName   string   `json:"default_from_name"`
	DefaultFromEmail  string   `json:"default_from_email"`
	DefaultSubject    string   `json:"default_subject"`
	DefaultLanguage   string   `json:"default_language"`
	ListRating        float64  `json:"list_rating"`
	SubscribeShortUrl string   `json:"subscribe_url_short"`
	SubscribeLongUrl  string   `json:"subscribe_url_long"`
	BeamerAddress     string   `json:"beamer_address"`
	Visibility        string   `json:"visibility"`
	Stats             ListStat `json:"stats"`
	Modules           []string `json:"modules"`
}

type BatchResponse struct {
	Success     int          `json:"success_count"`
	ErrorCount  int          `json:"error_count"`
	BatchErrors []BatchError `json:"errors"`
}

type BatchSubscribeResponse struct {
	AddCount    int                    `json:"add_count"`
	Adds        []Email                `json:"adds"`
	UpdateCount int                    `json:"update_count"`
	Updates     []Email                `json:"updates"`
	ErrorCount  int                    `json:"error_count"`
	Error       []BatchSubscriberError `json:"errors"`
}

type BatchError struct {
	Emails Email  `json:"email"`
	Code   int    `json:"code"`
	Error  string `json:"error"`
}

type BatchSubscriberError struct {
	Emails Email       `json:"email"`
	Code   int         `json:"code"`
	Error  string      `json:"error"`
	Row    interface{} `json:"row"`
}

type ListError struct {
	Param string `json:"param"`
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type ListsListResponse struct {
	Total  int         `json:"total"`
	Data   []ListData  `json:"data"`
	Errors []ListError `json:"errors"`
}

type ListsList struct {
	ApiKey        string     `json:"apikey"`
	Filters       ListFilter `json:"filters,omitempty"`
	Start         int        `json:"start,omitempty"`
	Limit         int        `json:"limit,omitempty"`
	SortField     string     `json:"sort_field,omitempty"`
	SortDirection string     `json:"sort_dir,omitempty"`
}

type UpdateMember struct {
	ApiKey           string                 `json:"apikey"`
	ListId           string                 `json:"id"`
	Email            Email                  `json:"email"`
	MergeVars        map[string]interface{} `json:"merge_vars,omitempty"`
	EmailType        string                 `json:"email_type,omitempty"`
	ReplaceInterests bool                   `json:"replace_interests"`
}

type Email struct {
	Email string `json:"email"`
	Euid  string `json:"euid"`
	Leid  string `json:"leid"`
}

type ListsMembers struct {
	ApiKey  string          `json:"apikey"`
	ListId  string          `json:"id"`
	Status  string          `json:"status"`
	Options ListsMembersOpt `json:"opts,omitempty"`
}

type ListsMember struct {
	Email     Email                  `json:"email"`
	EmailType string                 `json:"emailtype"`
	MergeVars map[string]interface{} `json:"merge_vars,omitempty"`
}

type ListsMembersOpt struct {
	Start         int    `json:"start,omitempty"`
	Limit         int    `json:"limit,omitempty"`
	SortField     string `json:"sort_field,omitempty"`
	SortDirection string `json:"sort_dir,omitempty"`
}

type ListsMembersResponse struct {
	Total int          `json:"total"`
	Data  []MemberInfo `json:"data"`
}

type ListsMemberInfo struct {
	ApiKey string  `json:"apikey"`
	ListId string  `json:"id"`
	Emails []Email `json:"emails"`
}

type ListsMemberInfoResponse struct {
	SuccessCount      int          `json:"success_count"`
	ErrorCount        int          `json:"error_count"`
	Errors            []ListError  `json:"errors"`
	MemberInfoRecords []MemberInfo `json:"data"`
}

type MemberInfo struct {
	Email           string                 `json:"email"`
	Euid            string                 `json:"euid"`
	EmailType       string                 `json:"email_type"`
	IpSignup        string                 `json:"ip_signup,omitempty"`
	TimestampSignup string                 `json:"timestamp_signup,omitempty"`
	IpOpt           string                 `json:"ip_opt"`
	TimestampOpt    string                 `json:"timestamp_opt"`
	MemberRating    int                    `json:"member_rating"`
	InfoChanged     string                 `json:"info_changed"`
	Leid            int                    `json:"leid"`
	Language        string                 `json:"language,omitempty"`
	ListId          string                 `json:"list_id"`
	ListName        string                 `json:"list_name"`
	Merges          map[string]interface{} `json:"merges"`
	Status          string                 `json:"status"`
	Timestamp       string                 `json:"timestamp"`
}
