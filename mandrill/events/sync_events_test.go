package events

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseSyncEvent(t *testing.T) {
	testCases := map[string]struct {
		outerEvent   WebhookEvent
		expectedType interface{}
		err          error
	}{
		"blacklist_remove": {
			outerEvent: WebhookEvent{
				Type:           SyncEventType,
				InnerEventType: "blacklist",
				raw:            []byte(`{"type":"blacklist","action":"remove","reject":{"reason":"hard-bounce","detail":"Example detail","last_event_at":"2014-02-01 12:43:56","email":"example.webhook@mandrillapp.com","created_at":"2014-01-15 11:32:19","expires_at":"2020-04-02 12:09:18","expired":false,"subaccount":"example_subaccount","sender":"example.sender@mandrillapp.com"},"ts":1580762074}`),
			},
			expectedType: BlacklistEvent{},
			err:          nil,
		},
		"whitelist_add": {
			outerEvent: WebhookEvent{
				Type:           SyncEventType,
				InnerEventType: "whitelist",
				raw: []byte(`{
					"type": "whitelist",
					"action": "add",
					"entry": {
						"email": "example.webhook@mandrillapp.com",
						"detail": "example details",
						"created_at": "2014-01-15 12:03:19"
					},
					"ts": 1580762074
				}`),
			},
			expectedType: WhitelistEvent{},
			err:          nil,
		},
		"non_sync": {
			outerEvent: WebhookEvent{
				Type: MessageEventType,
			},
			err: InvalidEventType{},
		},
		"default": {
			outerEvent: WebhookEvent{
				Type:           MessageEventType,
				InnerEventType: "foobar",
			},
			err: InvalidEventType{},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := parseSyncEvent(testCase.outerEvent)
			if testCase.err == nil {
				require.NoError(t, err)
				require.IsType(t, testCase.expectedType, res.Data)
			} else {
				require.Error(t, err)
				require.IsType(t, testCase.err, err)
			}
		})
	}
}
