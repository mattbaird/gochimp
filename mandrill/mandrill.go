package mandrill

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"sync"

	"github.com/pkg/errors"
)

const (
	baseURL    = "https://mandrillapp.com/api"
	apiVersion = "1.0"
	debug      = false // nolint: deadcode,unused
)

var (
	// I'm not a fan of globals like this
	// but I'm a bigger fan of more comfortable
	// apis where I can operate directly on an instance
	// of a thing (i.e. foo.Delete() vs. client.DeleteFoo(foo))
	globalClient = &Client{}
)

// ClientOption represents a configuration option for Client
type ClientOption func(*Client) error

// Client represents a Mandrill API Client
type Client struct {
	subaccount string
	sync.Mutex
	apiKey     string
	debug      bool
	httpClient *http.Client
	endpoint   string
	doPing     bool
	logger     *log.Logger
}

// WithLogger sets a custom stdlib logger to use
func WithLogger(l *log.Logger) ClientOption {
	return func(c *Client) error {
		c.logger = l
		return nil
	}
}

// WithSubAccount scopes all operations to the provided
// subaccount (where applicable)
func WithSubAccount(s string) ClientOption {
	return func(c *Client) error {
		c.subaccount = s
		return nil
	}
}

// WithDebug enables debug logging
func WithDebug() ClientOption {
	return func(c *Client) error {
		c.debug = true
		return nil
	}
}

// WithHTTPClient lets you set a custom http.Client for use with requests
func WithHTTPClient(h *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = h
		return nil
	}
}

// WithEndpoint lets you set a custom endpoint to use
func WithEndpoint(s string) ClientOption {
	return func(c *Client) error {
		c.endpoint = s
		return nil
	}
}

// WithPing attempt to validate Mandrill connectivity at client creation time for fast failures
func WithPing() ClientOption {
	return func(c *Client) error {
		c.doPing = true
		return nil
	}
}

// Connect is a simplified global client connection creator
// for use with things like the MessageBuilder
func Connect(key string, opts ...ClientOption) error {
	c, err := New(key, opts...)
	globalClient = c
	return err
}

// New returns a new mandrill.Client
func New(key string, opts ...ClientOption) (*Client, error) {
	c := &Client{}
	for _, opt := range opts {
		c.Lock()
		if err := opt(c); err != nil {
			c.Unlock()
			return nil, errors.Wrap(err, "unable to apply option")
		}
		c.Unlock()
	}
	c.Lock()
	defer c.Unlock()
	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}
	if c.endpoint == "" {
		c.endpoint = fmt.Sprintf("%s/%s", baseURL, apiVersion)
	}
	if c.logger == nil {
		c.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
	c.apiKey = key
	if c.doPing {
		err := c.Ping()
		if err != nil {
			return nil, err
		}
	}
	globalClient = c
	return c, nil
}

func (c *Client) debugLog(msg string) {
	if c.debug {
		c.logger.Println("DEBUG: " + msg)
	}
}

func (c *Client) post(path string, t interface{}, d interface{}) error {
	c.debugLog(fmt.Sprintf("Request Type: %T", t))
	c.debugLog(fmt.Sprintf("Response Type: %T", d))
	c.debugLog(fmt.Sprintf("Request: %+v", t))
	// I dislike using reflect but
	// this is the "cleanest" way to wrap all JSON post bodies with the apikey
	// Another option would be to have all api types defined to satisfy
	// an interface with a SetKey(string)
	v := reflect.ValueOf(t).Elem().FieldByName("Key")
	if v.IsValid() {
		v.SetString(c.apiKey)
	} else {
		return fmt.Errorf("Unable to add api key to request")
	}

	body, err := json.Marshal(t)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to marshal requested type %T to json", t))
	}
	c.debugLog("Request JSON: " + string(body))
	u := fmt.Sprintf("%s/%s.json", c.endpoint, path)
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewBuffer(body))
	if err != nil {
		return errors.Wrap(err, "unable to create a new http request")
	}
	c.debugLog("URL: " + req.URL.String())
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "error performing http request to endpoint")
	}

	defer resp.Body.Close()
	c.debugLog(fmt.Sprintf("Status: %s | StatusCode: %d", resp.Status, resp.StatusCode))
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "unable to read response body")
	}
	c.debugLog(fmt.Sprintf("Body: %s", b))
	if resp.StatusCode == 500 {
		return parseErrorJSON(b)
	}
	if err := json.Unmarshal(b, d); err != nil {
		return errors.Wrap(err, fmt.Sprintf("unable to unmarshal json into requested type %T", d))
	}
	return nil
}
