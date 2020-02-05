package events

import (
	"github.com/lusis/gochimp/mandrill/api"
)

// MandrillEventJSON represents an unparsed mandrill event
// we only capture the fields we need to parse and provide basic information:
// - Type is used by Sync events
// - Event is used by Message events and Inbound Message events
// We capture timestamp for debugging
type MandrillEventJSON struct {
	Type  string `json:"type"`
	Event string `json:"event"`
	TS    api.TS `json:"ts"`
}

// MessageEventJSON is a set of common fields for Message Events
type MessageEventJSON struct {
	TS    api.TS `json:"ts"`
	Event string `json:"event"`
	ID    string `json:"_id"`
}

// SyncEventJSON is a set of common fields for Sync Events
type SyncEventJSON struct {
	TS     api.TS `json:"ts"`
	Type   string `json:"type"`
	Action string `json:"action"`
}

// MessageEventMsg represents the Msg field of a MessageEvent
type MessageEventMsg struct {
	TS       api.TS                 `json:"ts"`
	ID       string                 `json:"_id"`
	Version  string                 `json:"_version"`
	Subject  string                 `json:"subject"`
	Email    string                 `json:"email"`
	Sender   string                 `json:"sender"`
	Tags     []string               `json:"tags"`
	State    string                 `json:"state"`
	MetaData map[string]interface{} `json:"metadata"`
	Template string                 `json:"template"`
}

// MessageSMTPEvent represents the SMTP details of a MessageEvent
type MessageSMTPEvent struct {
	TS            api.TS `json:"ts"`
	DestinationIP string `json:"destination_ip"`
	Diag          string `json:"diag"`
	SourceIP      string `json:"source_ip"`
	Type          string `json:"type"`
	Size          int64  `json:"size"`
}

// MessageEventOpen represents the an Open entry in a MessageEvent
type MessageEventOpen struct {
	TS api.TS `json:"ts"`
}

// MessageEventClick represents a click entry in a MessageEvent
type MessageEventClick struct {
	TS  api.TS `json:"ts"`
	URL string `json:"url"`
}

// MessageEventUserAgentParsed represents the UserAgent details of a message event
type MessageEventUserAgentParsed struct {
	Mobile       bool   `json:"mobile"`
	OSCompany    string `json:"os_company"`
	OSCompanyURL string `json:"os_company_url"`
	OSFamily     string `json:"os_family"`
	OSIcon       string `json:"os_icon"`
	OSName       string `json:"os_name"`
	OSURL        string `json:"os_url"`
	Type         string `json:"type"`
	UACompany    string `json:"ua_company"`
	UACompanyURL string `json:"ua_company_url"`
	UAFamily     string `json:"ua_family"`
	UAIcon       string `json:"ua_icon"`
	UAName       string `json:"ua_name"`
	UAURL        string `json:"ua_url"`
	UAVersion    string `json:"ua_version"`
}

// MessageEventLocation represents the Location details of a message event
type MessageEventLocation struct {
	CountryShort string  `json:"country_short"`
	Country      string  `json:"country"`
	Region       string  `json:"region"`
	City         string  `json:"city"`
	PostalCode   string  `json:"postal_code"`
	TimeZone     string  `json:"timezone"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

// SendMessageEvent is a message send event
type SendMessageEvent struct {
	MessageEventJSON
	Msg struct {
		MessageEventMsg
		Opens  []MessageEventOpen  `json:"opens"`
		Clicks []MessageEventClick `json:"clicks"`
	} `json:"msg"`
}

// DeferralMessageEvent is a message defferal event
type DeferralMessageEvent struct {
	MessageEventJSON
	Msg struct {
		MessageEventMsg
		Opens      []MessageEventOpen  `json:"opens"`
		Clicks     []MessageEventClick `json:"clicks"`
		SMTPEvents []MessageSMTPEvent  `json:"smtp_events"`
	} `json:"msg"`
}

// HardBounceMessageEvent is a message hardbounce event
type HardBounceMessageEvent struct {
	MessageEventJSON
	Msg struct {
		MessageEventMsg
		BounceDescription string `json:"bounce_description"`
		BGToolsCode       int64  `json:"bgtools_code"`
		Diag              string `json:"diag"`
	} `json:"msg"`
}

// SoftBounceMessageEvent is a message softbounce event
type SoftBounceMessageEvent struct {
	MessageEventJSON
	Msg struct {
		MessageEventMsg
		BounceDescription string `json:"bounce_description"`
		BGToolsCode       int64  `json:"bgtools_code"`
		Diag              string `json:"diag"`
	} `json:"msg"`
}

// OpenMessageEvent is a message open event
type OpenMessageEvent struct {
	MessageEventJSON
	IP              string                      `json:"ip"`
	UserAgent       string                      `json:"user_agent"`
	Location        MessageEventLocation        `json:"location"`
	UserAgentParsed MessageEventUserAgentParsed `json:"user_agent_parsed"`
	Msg             struct {
		MessageEventMsg
		Opens  []MessageEventOpen  `json:"opens"`
		Clicks []MessageEventClick `json:"clicks"`
	} `json:"msg"`
}

// ClickMessageEvent is a message click event
type ClickMessageEvent struct {
	MessageEventJSON
	IP              string                      `json:"ip"`
	UserAgent       string                      `json:"user_agent"`
	Location        MessageEventLocation        `json:"location"`
	UserAgentParsed MessageEventUserAgentParsed `json:"user_agent_parsed"`
	Msg             struct {
		MessageEventMsg
		Opens  []MessageEventOpen  `json:"opens"`
		Clicks []MessageEventClick `json:"clicks"`
	} `json:"msg"`
	URL string `json:"url"`
}

// SpamMessageEvent is a message spam event
type SpamMessageEvent struct {
	MessageEventJSON
	Msg struct {
		MessageEventMsg
		Opens  []MessageEventOpen  `json:"opens"`
		Clicks []MessageEventClick `json:"clicks"`
	} `json:"msg"`
}

// UnsubMessageEvent is a message unsub event
type UnsubMessageEvent struct {
	MessageEventJSON
	Msg struct {
		MessageEventMsg
		Opens  []MessageEventOpen  `json:"opens"`
		Clicks []MessageEventClick `json:"clicks"`
	} `json:"msg"`
}

// RejectMessageEvent is a message reject event
type RejectMessageEvent struct {
	MessageEventJSON
	Msg struct {
		MessageEventMsg
		Opens  []MessageEventOpen  `json:"opens"`
		Clicks []MessageEventClick `json:"clicks"`
	} `json:"msg"`
}

// BlacklistMessageEvent is a blacklist sync event
type BlacklistMessageEvent struct {
	SyncEventJSON
	Reject struct {
		Reason      string   `json:"reason"`
		Detail      string   `json:"detail"`
		LastEventAt api.Time `json:"last_event_at"`
		Email       string   `json:"email"`
		CreatedAt   api.Time `json:"created_at"`
		ExpiresAt   api.Time `json:"expires_at"`
		Expired     bool     `json:"expired"`
		SubAccount  string   `json:"subaccount"`
		Sender      string   `json:"sender"`
	} `json:"reject"`
}

// WhitelistMessageEvent is a whitelist sync event
// https://mandrill.zendesk.com/hc/en-us/articles/205583297-Sync-Event-Webhook-format
type WhitelistMessageEvent struct {
	SyncEventJSON
	Entry struct {
		Email     string   `json:"email"`
		Detail    string   `json:"detail"`
		CreatedAt api.Time `json:"created_at"`
	} `json:"entry"`
}

// InboundMessageEvent is an inbound message event
type InboundMessageEvent struct {
	MessageEventJSON
	Msg struct {
		RawMsg    string   `json:"raw_msg"`
		Headers   []string `json:"headers"`
		Text      string   `json:"text"`
		HTML      string   `json:"html"`
		FromEmail string   `json:"from_email"`
		FromName  string   `json:"from_name"`
		To        []struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"to"`
		Email       string   `json:"email"`
		Subject     string   `json:"subject"`
		Tags        []string `json:"tags"`
		Sender      string   `json:"sender"`
		Attachments []struct {
			Name    string `json:"name"`
			Type    string `json:"type"`
			Content string `json:"content"`
			Base64  bool   `json:"base64"`
		} `json:"attachments"`
		Images []struct {
			Name    string `json:"name"`
			Type    string `json:"type"`
			Content string `json:"content"`
		} `json:"images"`
		SpamReport struct {
			Score        int64 `json:"score"`
			MatchedRules []struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				Score       int64  `json:"score"`
				SPF         struct {
					Result string `json:"result"`
					Detail string `json:"detail"`
				} `json:"spf"`

				DKIM struct {
					Signed bool `json:"signed"`
					Valid  bool `json:"valid"`
				} `json:"dkim"`
			} `json:"matched_rules"`
		}
	} `json:"msg"`
}
