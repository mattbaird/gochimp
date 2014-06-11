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
	lists_subscribe_endpoint     string = "/lists/subscribe.json"
	lists_unsubscribe_endpoint   string = "/lists/unsubscribe.json"
	lists_list_endpoint          string = "/lists/list.json"
	lists_update_member_endpoint string = "/lists/update-member.json"
)

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
	TargetSubRate             float64 `json:"target_sub_rate "`
	OpenRate                  float64 `json:"open_rate "`
	ClickRate                 float64 `json:"click_rate "`
}

type ListData struct {
	Id                string     `json:"id"`
	WebId             string     `json:"web_id"`
	Name              string     `json:"name"`
	DateCreated       string     `json:"date_created"`
	EmailTypeOption   string     `json:"email_type_option"`
	UseAwesomeBar     bool       `json:"use_awesomebar"`
	DefaultFromName   string     `json:"default_from_name"`
	DefaultFromEmail  string     `json:"default_from_email"`
	DefaultSubject    string     `json:"default_subject"`
	DefaultLanguage   string     `json:"default_language"`
	ListRating        float64    `json:"list_rating"`
	SubscribeShortUrl string     `json:"subscribe_url_short"`
	SubscribeLongUrl  string     `json:"subscribe_url_long"`
	BeamerAddress     string     `json:"beamer_address"`
	Visibility        string     `json:"visibility"`
	Stats             []ListStat `json:"stats"`
	Modules           []string   `json:"modules"`
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
	Filters       ListFilter `json:"filters"`
	Start         int        `json:"start"`
	Limit         int        `json:"limit"`
	SortField     string     `json:"sort_field"`
	SortDirection string     `json:"sort_direction"`
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
