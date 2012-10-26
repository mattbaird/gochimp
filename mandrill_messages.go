/**
* Copyright 2012 Matthew Baird
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*  distributed under the License is distributed on an "AS IS" BASIS,
*  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  See the License for the specific language governing permissions and
*  limitations under the License.
**/
package gochimp

import (
	"encoding/json"
)

// see https://mandrillapp.com/api/docs/messages.html
const messages_send_endpoint string = "/messages/send.json" //Send a new transactional message through Mandrill

func (a *MandrillAPI) Send(message Message, async bool) ([]SendResponse, error) {
	var response []SendResponse
	var params map[string]interface{} = make(map[string]interface{})
	params["message"] = message
	err := parseMandrillJson(a, messages_send_endpoint, params, &response)
	return response, err
}

type Message struct {
	Html                    string              `json:"html"`
	Text                    string              `json:"text"`
	Subject                 string              `json:"subject"`
	FromEmail               string              `json:"from_email"`
	FromName                string              `json:"from_name"`
	To                      []Recipient         `json:"to"`
	Headers                 map[string]string   `json:"headers,omitempty"`
	TrackOpens              bool                `json:"track_opens"`
	TrackClicks             bool                `json:"track_clicks"`
	AutoText                bool                `json:"auto_text"`
	UrlStripQS              bool                `json:"url_strip_qs"`
	PreserveRecipients      bool                `json:"preserve_recipients"`
	BCCAddress              string              `json:"bcc_address"`
	Merge                   bool                `json:"merge"`
	GlobalMergeVars         []Var               `json:"global_merge_vars,omitempty"`
	MergeVars               []MergeVars         `json:"merge_vars,omitempty"`
	Tags                    []string            `json:"tags,omitempty"`
	GoogleAnalyticsDomains  []string            `json:"google_analytics_domains,omitempty"`
	GoogleAnalyticsCampaign []string            `json:"google_analytics_campaign,omitempty"`
	Metadata                []map[string]string `json:"metadata,omitempty"`
	RecipientMetadata       []RecipientMetaData `json:"recipient_metadata,omitempty"`
	Attachments             []Attachment        `json:"attachments,omitempty"`
}

func (m *Message) String() string {
	b, _ := json.Marshal(m)
	return string(b)
}

func (m *Message) addRecipients(r ...Recipient) {
	m.To = append(m.To, r...)
}

func (m *Message) addGlobalMergeVar(globalvars ...Var) {
	m.GlobalMergeVars = append(m.GlobalMergeVars, globalvars...)
}

func (m *Message) addMergeVar(vars ...MergeVars) {
	m.MergeVars = append(m.MergeVars, vars...)
}

func (m *Message) AddTag(tags ...string) {
	m.Tags = append(m.Tags, tags...)
}

func (m *Message) addGoogleAnalyticsDomains(domains ...string) {
	m.GoogleAnalyticsDomains = append(m.GoogleAnalyticsDomains, domains...)
}

func (m *Message) addGoogleAnalyticsCampaign(campaigns ...string) {
	m.GoogleAnalyticsCampaign = append(m.GoogleAnalyticsCampaign, campaigns...)
}

func (m *Message) addMetadata(metadata ...map[string]string) {
	m.Metadata = append(m.Metadata, metadata...)
}

func (m *Message) addRecipientMetadata(metadata ...RecipientMetaData) {
	m.RecipientMetadata = append(m.RecipientMetadata, metadata...)
}

func (m *Message) addAttachments(attachement ...Attachment) {
	m.Attachments = append(m.Attachments, attachement...)
}

type Attachment struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
type RecipientMetaData struct {
	Recipient string            `json:"rcpt"`
	Vars      map[string]string `json:"values"`
}

type MergeVars struct {
	Recipient string `json:"rcpt"`
	Vars      []Var  `json:"vars"`
}

type Var struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Recipient struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type SendResponse struct {
	Email  string `json:"email"`
	Status string `json:"status"`
}
