package gochimp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"time"
)

// URIFormat defines the endpoint for a single app
const URIFormat string = "%s.api.mailchimp.com"

// Version the latest API version
const Version string = "/3.0"

// DatacenterRegex defines which datacenter to hit
var DatacenterRegex = regexp.MustCompile("[^-]\\w+$")

// ChimpAPI represents the origin of the API
type ChimpAPI struct {
	Key       string
	Timeout   time.Duration
	Transport http.RoundTripper

	User  string
	Debug bool

	endpoint string
}

// NewChimp creates a ChimpAPI
func NewChimp(apiKey string, https bool) *ChimpAPI {
	u := url.URL{}
	u.Scheme = "http"
	if https {
		u.Scheme = "https"
	}

	u.Host = fmt.Sprintf(URIFormat, DatacenterRegex.FindString(apiKey))
	u.Path = Version

	return &ChimpAPI{
		User:     "gochimp",
		Key:      apiKey,
		endpoint: u.String(),
	}
}

// Request will make a call to the actual API.
func (api ChimpAPI) Request(method, path string, params QueryParams, body, response interface{}) error {
	client := &http.Client{Transport: api.Transport}
	if api.Timeout > 0 {
		client.Timeout = api.Timeout
	}

	requestURL := fmt.Sprintf("%s%s", api.endpoint, path)
	if api.Debug {
		log.Printf("Requesting %s: %s\n", method, requestURL)
	}

	var bodyBytes io.Reader
	var err error
	var data []byte
	if body != nil {
		data, err = json.Marshal(body)
		if err != nil {
			return err
		}
		bodyBytes = bytes.NewBuffer(data)
		if api.Debug {
			log.Printf("Adding body: %+v\n", body)
		}
	}

	req, err := http.NewRequest(method, requestURL, bodyBytes)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(api.User, api.Key)

	if params != nil && !reflect.ValueOf(params).IsNil() {
		queryParams := req.URL.Query()
		for k, v := range params.Params() {
			queryParams.Set(k, v)
		}
		if api.Debug {
			log.Printf("Adding query params: %q\n", req.URL.Query())
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if api.Debug {
		log.Printf("Got response: %d %s\n", resp.StatusCode, resp.Status)
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		// Do not unmarshall response is nil
		if response == nil || reflect.ValueOf(response).IsNil() {
			return nil
		}

		err = json.Unmarshal(data, response)
		if err != nil {
			return err
		}

		return nil
	}

	// This is an API Error
	return parseAPIError(data)
}

// RequestOk Make Request ignoring body and return true if HTTP status code is 2xx.
func (api ChimpAPI) RequestOk(method, path string) (bool, error) {
	err := api.Request(method, path, nil, nil, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func parseAPIError(data []byte) error {
	apiError := new(APIError)
	err := json.Unmarshal(data, apiError)
	if err != nil {
		return err
	}

	return apiError
}
