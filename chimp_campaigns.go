// Copyright 2012 Matthew Baird
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gochimp

import (
	"fmt"
)

// see https://mandrillapp.com/api/docs/messages.html
const get_content_endpoint string = "/campaigns/content.%s" // Get the content (both html and text) for a campaign either as it would appear in the campaign archive or as the raw, original content

func (a *ChimpAPI) getContent(apiKey string, cid string, options map[string]interface{}, contentFormat string) ([]SendResponse, error) {
	var response []SendResponse
	var params map[string]interface{} = make(map[string]interface{})
	params["apikey"] = apiKey
	params["cid"] = cid
	params["options"] = options
	err := parseChimpJson(a, fmt.Sprintf(get_content_endpoint, contentFormat), params, &response)
	return response, err
}
