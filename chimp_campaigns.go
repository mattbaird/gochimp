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

import (
	"fmt"
)

const (
	get_content_endpoint     string = "/campaigns/content.%s"
	campaign_create_endpoint string = "/campaigns/create.json"
	campaign_send_endpoint   string = "/campaigns/send.json"
)

func (a *ChimpAPI) getContent(apiKey string, cid string, options map[string]interface{}, contentFormat string) ([]SendResponse, error) {
	var response []SendResponse
	var params map[string]interface{} = make(map[string]interface{})
	params["apikey"] = apiKey
	params["cid"] = cid
	params["options"] = options
	err := parseChimpJson(a, fmt.Sprintf(get_content_endpoint, contentFormat), params, &response)
	return response, err
}

func (a *ChimpAPI) CampaignCreate(req CampaignCreate) (CampaignCreateResponse, error) {
	req.ApiKey = a.Key
	var response CampaignCreateResponse
	err := parseChimpJson(a, campaign_create_endpoint, req, &response)
	return response, err
}

func (a *ChimpAPI) CampaignSend(cid string) (CampaignSendResponse, error) {
	req := campaignSend{
		ApiKey:     a.Key,
		CampaignId: cid,
	}
	var response CampaignSendResponse
	err := parseChimpJson(a, campaign_send_endpoint, req, &response)
	return response, err
}

type campaignSend struct {
	ApiKey     string `json:"apikey"`
	CampaignId string `json:"cid"`
}

type CampaignSendResponse struct {
	Complete bool `json:"complete"`
}

type CampaignCreate struct {
	ApiKey  string                `json:"apikey"`
	Type    string                `json:"type"`
	Options CampaignCreateOptions `json:"options"`
	Content CampaignCreateContent `json:"content"`
}

type CampaignCreateOptions struct {
	// ListID is the list to send this campaign to
	ListID string `json:"list_id"`

	// Subject is the subject line for your campaign message
	Subject string `json:"subject"`

	// FromEmail is the From: email address for your campaign message
	FromEmail string `json:"from_email"`

	// FromName is the From: name for your campaign message (not an email address)
	FromName string `json:"from_name"`

	// ToName is the To: name recipients will see (not email address)
	ToName string `json:"to_name"`
}

type CampaignCreateContent struct {
	// HTML is the raw/pasted HTML content for the campaign
	HTML string `json:"html"`

	// When using a template instead of raw HTML, each key
	// in the map should be the unique mc:edit area name from
	// the template.
	Sections map[string]string `json:"sections,omitempty"`

	// Text is the plain-text version of the body
	Text string `json:"text"`

	// MailChimp will pull in content from this URL. Note,
	// this will override any other content options - for lists
	// with Email Format options, you'll need to turn on
	// generate_text as well
	URL string `json:"url,omitempty"`

	// A Base64 encoded archive file for MailChimp to import all
	// media from. Note, this will override any other content
	// options - for lists with Email Format options, you'll
	// need to turn on generate_text as well
	Archive string `json:"archive,omitempty"`

	// ArchiveType only applies to the Archive field. Supported
	// formats are: zip, tar.gz, tar.bz2, tar, tgz, tbz.
	// If not included, we will default to zip
	ArchiveType string `json:"archive_options,omitempty"`
}

type CampaignCreateResponse struct {
	Id                 string           `json:"id"`
	WebId              int              `json:"web_id"`
	ListId             string           `json:"list_id"`
	FolderId           int              `json:"folder_id"`
	TemplateId         int              `json:"template_id"`
	ContentType        string           `json:"content_type"`
	ContentEditedBy    string           `json:"content_edited_by"`
	Title              string           `json:"title"`
	Type               string           `json:"type"`
	CreateTime         string           `json:"create_time"`
	SendTime           string           `json:"send_time"`
	ContentUpdatedTime string           `json:"content_updated_time"`
	Status             string           `json:"status"`
	FromName           string           `json:"from_name"`
	FromEmail          string           `json:"from_email"`
	Subject            string           `json:"subject"`
	ToName             string           `json:"to_name"`
	ArchiveURL         string           `json:"archive_url"`
	ArchiveURLLong     string           `json:"archive_url_long"`
	EmailsSent         int              `json:"emails_sent"`
	Analytics          string           `json:"analytics"`
	AnalyticsTag       string           `json:"analytics_tag"`
	InlineCSS          bool             `json:"inline_css"`
	Authenticate       bool             `json:authenticate"`
	Ecommm360          bool             `json:"ecomm360"`
	AutoTweet          bool             `json:"auto_tweet"`
	AutoFacebookPort   string           `json:"auto_fb_post"`
	AutoFooter         bool             `json:"auto_footer"`
	Timewarp           bool             `json:"timewarp"`
	TimewarpSchedule   string           `json:"timewarp_schedule,omitempty"`
	Tracking           CampaignTracking `json:"tracking"`
	ParentId           string           `json:"parent_id"`
	IsChild            bool             `json:"is_child"`
	TestsSent          int              `json:"tests_sent"`
	TestsRemaining     int              `json:"tests_remain"`
	SegmentText        string           `json:"segment_text"`
}

type CampaignTracking struct {
	HTMLClicks bool `json:"html_clicks"`
	TextClicks bool `json:"text_clicks"`
	Opens      bool `json:"opens"`
}
