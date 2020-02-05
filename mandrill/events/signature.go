package events

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"sort"
	"strings"
)

// ValidateWebhookSignature generates a webhook signature
func ValidateWebhookSignature(url string, key string, req *http.Request) error {
	signature := req.Header.Get("x-mandrill-signature")
	if len(signature) == 0 {
		return &ErrNoSignature{}
	}
	toSign := []string{url}
	if req.Method != http.MethodPost {
		return &ErrNotPost{method: req.Method}
	}
	err := req.ParseForm()
	if err != nil {
		return &ErrParse{parseErr: err.Error()}
	}
	if len(req.Form.Get("mandrill_events")) == 0 {
		return &ErrMissingMandrillEvents{postForm: req.Form.Encode()}
	}
	keys := make([]string, 0, len(req.Form))
	for k := range req.Form {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		toSign = append(toSign, k, req.Form.Get(k))
	}
	mac := hmac.New(sha1.New, []byte(key))
	_, err = mac.Write([]byte(strings.Join(toSign, "")))
	if err != nil {
		return &ErrCryptoErr{err: err.Error()}
	}
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	if res != signature {
		return &ErrInvalidSignature{provided: signature, computed: res}
	}
	return nil
}
