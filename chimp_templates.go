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

const (
	chimp_templates_list_endpoint string = "/templates/list.json"
)

func (a *ChimpAPI) TemplatesList(req TemplatesList) (TemplatesListResponse, error) {
	req.ApiKey = a.Key
	var response TemplatesListResponse
	err := parseChimpJson(a, chimp_templates_list_endpoint, req, &response)
	return response, err
}

type TemplateListType struct {
	User    bool `json:"user"`
	Gallery bool `json:"gallery"`
	Base    bool `json:"base"`
}

type TemplateListFilter struct {
	Category           string `json:"category"`
	FolderId           string `json:"folder_id"`
	IncludeInactive    bool   `json:"include_inactive"`
	InactiveOnly       bool   `json:"inactive_only"`
	IncludeDragAndDrop bool   `json:"include_drag_and_drop"`
}

type UserTemplate struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Layout       string `json:"layout"`
	Category     string `json:"category"`
	PreviewImage string `json:"preview_image"`
	DateCreated  string `json:"date_created"`
	Active       bool   `json:"active"`
	EditSource   bool   `json:"edit_source"`
	FolderId     bool   `json:"folder_id"`
}

type GalleryTemplate struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Layout       string `json:"layout"`
	Category     string `json:"category"`
	PreviewImage string `json:"preview_image"`
	DateCreated  string `json:"date_created"`
	Active       bool   `json:"active"`
	EditSource   bool   `json:"edit_source"`
}

type TemplatesList struct {
	ApiKey  string             `json:"apikey"`
	Types   TemplateListType   `json:"types"`
	Filters TemplateListFilter `json:"filters"`
}

type TemplatesListResponse struct {
	User    []UserTemplate    `json:"user"`
	Gallery []GalleryTemplate `json:"gallery"`
}

const (
	chimp_template_info_endpoint   string = "/templates/info.json"
	chimp_template_add_endpoint    string = "/templates/add.json"
	chimp_template_update_endpoint string = "/templates/update.json"
)

func (a *ChimpAPI) TemplatesInfo(req TemplateInfo) (TemplateInfoResponse, error) {
	req.ApiKey = a.Key
	var response TemplateInfoResponse
	err := parseChimpJson(a, chimp_template_info_endpoint, req, &response)
	return response, err
}

type TemplateInfo struct {
	ApiKey     string `json:"apikey"`
	TemplateID int    `json:"template_id"`
	Type       string `json:"type"`
}

type TemplateInfoResponse struct {
	DefaultContent interface{}
	Sections       interface{}
	Source         string
	Preview        string
}

func (a *ChimpAPI) TemplatesAdd(req TemplatesAdd) (TemplatesAddResponse, error) {
	req.ApiKey = a.Key
	var response TemplatesAddResponse
	err := parseChimpJson(a, chimp_template_add_endpoint, req, &response)
	return response, err
}

type TemplatesAdd struct {
	ApiKey   string `json:"apikey"`
	Name     string `json:"name"`
	HTML     string `json:"html"`
	FolderID int    `json:"folder_id,omitempty"`
}

type TemplatesAddResponse struct {
	TemplateID int `json:"template_id"`
}

func (a *ChimpAPI) TemplatesUpdate(req TemplatesUpdate) (TemplatesUpdateResponse, error) {
	req.ApiKey = a.Key
	var response TemplatesUpdateResponse
	err := parseChimpJson(a, chimp_template_update_endpoint, req, &response)
	return response, err
}

type TemplatesUpdate struct {
	ApiKey     string                `json:"apikey"`
	TemplateID int                   `json:"template_id"`
	Values     TemplatesUpdateValues `json:"values"`
}

type TemplatesUpdateValues struct {
	Name     string `json:"name"`
	HTML     string `json:"html"`
	FolderID int    `json:"folder_id,omitempty"`
}

type TemplatesUpdateResponse struct {
	Complete bool `json:"complete"`
}
