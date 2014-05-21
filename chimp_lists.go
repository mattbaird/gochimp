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
