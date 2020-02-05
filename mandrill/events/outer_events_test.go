package events

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// nolint: deadcode
var testData = `
[{
	"event": "send",
	"msg": {
		"ts": 1365109999,
		"subject": "This an example webhook message",
		"email": "example.webhook@mandrillapp.com",
		"sender": "example.sender@mandrillapp.com",
		"tags": ["webhook-example"],
		"opens": [],
		"clicks": [],
		"state": "sent",
		"metadata": {
			"user_id": 111
		},
		"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa",
		"_version": "exampleaaaaaaaaaaaaaaa"
	},
	"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa",
	"ts": 1580762074
}, {
	"event": "deferral",
	"msg": {
		"ts": 1365109999,
		"subject": "This an example webhook message",
		"email": "example.webhook@mandrillapp.com",
		"sender": "example.sender@mandrillapp.com",
		"tags": ["webhook-example"],
		"opens": [],
		"clicks": [],
		"state": "deferred",
		"metadata": {
			"user_id": 111
		},
		"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa1",
		"_version": "exampleaaaaaaaaaaaaaaa",
		"smtp_events": [{
			"destination_ip": "127.0.0.1",
			"diag": "451 4.3.5 Temporarily unavailable, try again later.",
			"source_ip": "127.0.0.1",
			"ts": 1365111111,
			"type": "deferred",
			"size": 0
		}]
	},
	"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa1",
	"ts": 1580762074
}, {
	"event": "hard_bounce",
	"msg": {
		"ts": 1365109999,
		"subject": "This an example webhook message",
		"email": "example.webhook@mandrillapp.com",
		"sender": "example.sender@mandrillapp.com",
		"tags": ["webhook-example"],
		"state": "bounced",
		"metadata": {
			"user_id": 111
		},
		"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa2",
		"_version": "exampleaaaaaaaaaaaaaaa",
		"bounce_description": "bad_mailbox",
		"bgtools_code": 10,
		"diag": "smtp;550 5.1.1 The email account that you tried to reach does not exist. Please try double-checking the recipient's email address for typos or unnecessary spaces."
	},
	"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa2",
	"ts": 1580762074
}, {
	"event": "soft_bounce",
	"msg": {
		"ts": 1365109999,
		"subject": "This an example webhook message",
		"email": "example.webhook@mandrillapp.com",
		"sender": "example.sender@mandrillapp.com",
		"tags": ["webhook-example"],
		"state": "soft-bounced",
		"metadata": {
			"user_id": 111
		},
		"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa3",
		"_version": "exampleaaaaaaaaaaaaaaa",
		"bounce_description": "mailbox_full",
		"bgtools_code": 22,
		"diag": "smtp;552 5.2.2 Over Quota"
	},
	"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa3",
	"ts": 1580762074
}, {
	"event": "open",
	"msg": {
		"ts": 1365109999,
		"subject": "This an example webhook message",
		"email": "example.webhook@mandrillapp.com",
		"sender": "example.sender@mandrillapp.com",
		"tags": ["webhook-example"],
		"opens": [{
			"ts": 1365111111
		}],
		"clicks": [{
			"ts": 1365111111,
			"url": "http:\/\/mandrill.com"
		}],
		"state": "sent",
		"metadata": {
			"user_id": 111
		},
		"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa4",
		"_version": "exampleaaaaaaaaaaaaaaa"
	},
	"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa4",
	"ip": "127.0.0.1",
	"location": {
		"country_short": "US",
		"country": "United States",
		"region": "Oklahoma",
		"city": "Oklahoma City",
		"latitude": 35.4675598145,
		"longitude": -97.5164337158,
		"postal_code": "73101",
		"timezone": "-05:00"
	},
	"user_agent": "Mozilla\/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.8) Gecko\/20100317 Postbox\/1.1.3",
	"user_agent_parsed": {
		"type": "Email Client",
		"ua_family": "Postbox",
		"ua_name": "Postbox 1.1.3",
		"ua_version": "1.1.3",
		"ua_url": "http:\/\/www.postbox-inc.com\/",
		"ua_company": "Postbox, Inc.",
		"ua_company_url": "http:\/\/www.postbox-inc.com\/",
		"ua_icon": "http:\/\/cdn.mandrill.com\/img\/email-client-icons\/postbox.png",
		"os_family": "OS X",
		"os_name": "OS X 10.6 Snow Leopard",
		"os_url": "http:\/\/www.apple.com\/osx\/",
		"os_company": "Apple Computer, Inc.",
		"os_company_url": "http:\/\/www.apple.com\/",
		"os_icon": "http:\/\/cdn.mandrill.com\/img\/email-client-icons\/macosx.png",
		"mobile": false
	},
	"ts": 1580762074
}, {
	"event": "click",
	"msg": {
		"ts": 1365109999,
		"subject": "This an example webhook message",
		"email": "example.webhook@mandrillapp.com",
		"sender": "example.sender@mandrillapp.com",
		"tags": ["webhook-example"],
		"opens": [{
			"ts": 1365111111
		}],
		"clicks": [{
			"ts": 1365111111,
			"url": "http:\/\/mandrill.com"
		}],
		"state": "sent",
		"metadata": {
			"user_id": 111
		},
		"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa5",
		"_version": "exampleaaaaaaaaaaaaaaa"
	},
	"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa5",
	"ip": "127.0.0.1",
	"location": {
		"country_short": "US",
		"country": "United States",
		"region": "Oklahoma",
		"city": "Oklahoma City",
		"latitude": 35.4675598145,
		"longitude": -97.5164337158,
		"postal_code": "73101",
		"timezone": "-05:00"
	},
	"user_agent": "Mozilla\/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.8) Gecko\/20100317 Postbox\/1.1.3",
	"user_agent_parsed": {
		"type": "Email Client",
		"ua_family": "Postbox",
		"ua_name": "Postbox 1.1.3",
		"ua_version": "1.1.3",
		"ua_url": "http:\/\/www.postbox-inc.com\/",
		"ua_company": "Postbox, Inc.",
		"ua_company_url": "http:\/\/www.postbox-inc.com\/",
		"ua_icon": "http:\/\/cdn.mandrill.com\/img\/email-client-icons\/postbox.png",
		"os_family": "OS X",
		"os_name": "OS X 10.6 Snow Leopard",
		"os_url": "http:\/\/www.apple.com\/osx\/",
		"os_company": "Apple Computer, Inc.",
		"os_company_url": "http:\/\/www.apple.com\/",
		"os_icon": "http:\/\/cdn.mandrill.com\/img\/email-client-icons\/macosx.png",
		"mobile": false
	},
	"url": "http:\/\/mandrill.com",
	"ts": 1580762074
}, {
	"event": "spam",
	"msg": {
		"ts": 1365109999,
		"subject": "This an example webhook message",
		"email": "example.webhook@mandrillapp.com",
		"sender": "example.sender@mandrillapp.com",
		"tags": ["webhook-example"],
		"opens": [{
			"ts": 1365111111
		}],
		"clicks": [{
			"ts": 1365111111,
			"url": "http:\/\/mandrill.com"
		}],
		"state": "sent",
		"metadata": {
			"user_id": 111
		},
		"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa6",
		"_version": "exampleaaaaaaaaaaaaaaa"
	},
	"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa6",
	"ts": 1580762074
}, {
	"event": "unsub",
	"msg": {
		"ts": 1365109999,
		"subject": "This an example webhook message",
		"email": "example.webhook@mandrillapp.com",
		"sender": "example.sender@mandrillapp.com",
		"tags": ["webhook-example"],
		"opens": [{
			"ts": 1365111111
		}],
		"clicks": [{
			"ts": 1365111111,
			"url": "http:\/\/mandrill.com"
		}],
		"state": "sent",
		"metadata": {
			"user_id": 111
		},
		"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa7",
		"_version": "exampleaaaaaaaaaaaaaaa"
	},
	"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa7",
	"ts": 1580762074
}, {
	"event": "reject",
	"msg": {
		"ts": 1365109999,
		"subject": "This an example webhook message",
		"email": "example.webhook@mandrillapp.com",
		"sender": "example.sender@mandrillapp.com",
		"tags": ["webhook-example"],
		"opens": [],
		"clicks": [],
		"state": "rejected",
		"metadata": {
			"user_id": 111
		},
		"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa8",
		"_version": "exampleaaaaaaaaaaaaaaa"
	},
	"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa8",
	"ts": 1580762074
}, {
	"type": "blacklist",
	"action": "add",
	"reject": {
		"reason": "hard-bounce",
		"detail": "Example detail",
		"last_event_at": "2014-02-01 12:43:56",
		"email": "example.webhook@mandrillapp.com",
		"created_at": "2014-01-15 11:32:19",
		"expires_at": "2020-04-02 12:09:18",
		"expired": false,
		"subaccount": "example_subaccount",
		"sender": "example.sender@mandrillapp.com"
	},
	"ts": 1580762074
}, {
	"type": "blacklist",
	"action": "change",
	"reject": {
		"reason": "hard-bounce",
		"detail": "Example detail",
		"last_event_at": "2014-02-01 12:43:56",
		"email": "example.webhook@mandrillapp.com",
		"created_at": "2014-01-15 11:32:19",
		"expires_at": "2020-04-02 12:09:18",
		"expired": false,
		"subaccount": "example_subaccount",
		"sender": "example.sender@mandrillapp.com"
	},
	"ts": 1580762074
}, {
	"type": "blacklist",
	"action": "remove",
	"reject": {
		"reason": "hard-bounce",
		"detail": "Example detail",
		"last_event_at": "2014-02-01 12:43:56",
		"email": "example.webhook@mandrillapp.com",
		"created_at": "2014-01-15 11:32:19",
		"expires_at": "2020-04-02 12:09:18",
		"expired": false,
		"subaccount": "example_subaccount",
		"sender": "example.sender@mandrillapp.com"
	},
	"ts": 1580762074
}, {
	"type": "whitelist",
	"action": "add",
	"entry": {
		"email": "example.webhook@mandrillapp.com",
		"detail": "example details",
		"created_at": "2014-01-15 12:03:19"
	},
	"ts": 1580762074
}, {
	"type": "whitelist",
	"action": "remove",
	"entry": {
		"email": "example.webhook@mandrillapp.com",
		"detail": "example details",
		"created_at": "2014-01-15 12:03:19"
	},
	"ts": 1580762074
}]
`

func TestBasicParseOuterEvents(t *testing.T) {
	testCases := map[string]struct {
		data      string
		eventType string
		innerType string
		err       error
	}{
		"unmarshal_error": {
			data: "foo",
			err:  UnmarshallError{},
		},
		"unknown_message": {
			data:      `{"event":"unknown", "ts":1580762074}`,
			eventType: MessageEventType,
			err:       InvalidEventType{},
		},
		"unknown_sync": {
			data:      `{"type":"unknown", "ts":1580762074}`,
			eventType: SyncEventType,
			err:       InvalidEventType{},
		},
		MessageEventType: {
			data: `{
				"event": "send",
				"msg": {
					"ts": 1365109999,
					"subject": "This an example webhook message",
					"email": "example.webhook@mandrillapp.com",
					"sender": "example.sender@mandrillapp.com",
					"tags": ["webhook-example"],
					"opens": [],
					"clicks": [],
					"state": "sent",
					"metadata": {
						"user_id": 111
					},
					"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa",
					"_version": "exampleaaaaaaaaaaaaaaa"
				},
				"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa",
				"ts": 1580762074
			}`,
			eventType: MessageEventType,
			innerType: "send",
			err:       nil,
		},
		SyncEventType: {
			data:      `{"type":"blacklist","action":"remove","reject":{"reason":"hard-bounce","detail":"Example detail","last_event_at":"2014-02-01 12:43:56","email":"example.webhook@mandrillapp.com","created_at":"2014-01-15 11:32:19","expires_at":"2020-04-02 12:09:18","expired":false,"subaccount":"example_subaccount","sender":"example.sender@mandrillapp.com"},"ts":1580762074}`,
			eventType: SyncEventType,
			innerType: "blacklist",
			err:       nil,
		},
	}
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			res, err := parseOuterEvent([]byte(testCase.data))
			if testCase.err == nil {
				require.NoError(t, err)
				require.Equal(t, testName, res.Type)
				require.Equal(t, testCase.data, string(res.raw))
				require.Equal(t, testCase.innerType, res.InnerEventType)
				require.NotNil(t, res.Timestamp)
			} else {
				require.Error(t, err)
				require.IsType(t, testCase.err, err)
			}
		})
	}
}
