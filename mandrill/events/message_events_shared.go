package events

import (
	"time"
)

// MessageEventCommon represents the default fields of all Message events
type MessageEventCommon struct {
	Timestamp time.Time
	Event     string
	ID        string
}

// MsgCommon represents the common fields between the Msg field of all Message events
type MsgCommon struct {
	Timestamp time.Time
	ID        string
	Version   string
	Subject   string
	Email     string
	Sender    string
	Tags      []string
	State     string
	MetaData  map[string]interface{}
	Template  string
}

// OpensClicks represents the Opens and Clicks in a Msg field
type OpensClicks struct {
	Opens  []Open
	Clicks []Click
}

// Open represents opens in a Msg field
type Open struct {
	Timestamp time.Time
}

// Click represents click in a Msg field
type Click struct {
	Timestamp time.Time
	URL       string
}

// SMTPEvent represents the SMTPEvents field of a message event
type SMTPEvent struct {
	Timestamp     time.Time
	DestinationIP string
	Diag          string
	SourceIP      string
	Type          string
	Size          int64
}

// BounceMsg is the msg field of a hard_bounce or soft_bounce event
type BounceMsg struct {
	MsgCommon
	BounceDescription string
	BGToolsCode       int64
	Diag              string
}

// Location represents location data in a message event
type Location struct {
	CountryShort string
	Country      string
	Region       string
	City         string
	PostalCode   string
	TimeZone     string
	Latitude     float64
	Longitude    float64
}

// UserAgentParsed represents useragent data in a message event
type UserAgentParsed struct {
	Mobile       bool
	OSCompany    string
	OSCompanyURL string
	OSFamily     string
	OSIcon       string
	OSName       string
	OSURL        string
	Type         string
	UACompany    string
	UACompanyURL string
	UAFamily     string
	UAIcon       string
	UAName       string
	UAURL        string
	UAVersion    string
}
