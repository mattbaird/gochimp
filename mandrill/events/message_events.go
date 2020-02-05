package events

// MessageEvent represents a basic sync event
type MessageEvent struct {
	Data interface{}
}

// SendEvent is a message send event
type SendEvent struct {
	MessageEventCommon
	Msg SendMsg
}

// SendMsg represents the Msg field of a send event
type SendMsg struct {
	MsgCommon
	OpensClicks
}

// DeferralEvent is a message deferral event
type DeferralEvent struct {
	MessageEventCommon
	Msg DeferralMsg
}

// DeferralMsg is the format of the msg field of deferral events
type DeferralMsg struct {
	MsgCommon
	OpensClicks
	SMTPEvents []SMTPEvent
}

// HardBounceEvent is a message hard_bounce event
type HardBounceEvent struct {
	MessageEventCommon
	Msg BounceMsg
}

// SoftBounceEvent is a message soft_bounce event
type SoftBounceEvent struct {
	MessageEventCommon
	Msg BounceMsg
}

// OpenEvent is a message open event
type OpenEvent struct {
	MessageEventCommon
	IP              string
	UserAgent       string
	Location        Location
	UserAgentParsed UserAgentParsed
	Msg             OpenMsg
}

// OpenMsg represents the Msg field of an open event
type OpenMsg struct {
	MsgCommon
	OpensClicks
}

// ClickEvent represents a click event
type ClickEvent struct {
	MessageEventCommon
	IP              string
	UserAgent       string
	Location        Location
	UserAgentParsed UserAgentParsed
	Msg             ClickMsg
}

// ClickMsg represents the msg field of a ClickEvent
type ClickMsg struct {
	MsgCommon
	OpensClicks
}

// SpamEvent is a message spam event
type SpamEvent struct {
	MessageEventCommon
	Msg SpamMsg
}

// SpamMsg represents the msg field of a spam event
type SpamMsg struct {
	MsgCommon
	OpensClicks
}

// UnsubEvent is a message unsub event
type UnsubEvent struct {
	MessageEventCommon
	Msg UnsubMsg
}

// UnsubMsg represents the msg field of an unsub event
type UnsubMsg struct {
	MsgCommon
	OpensClicks
}

// RejectEvent is a message reject event
type RejectEvent struct {
	MessageEventCommon
	Msg RejectMsg
}

// RejectMsg represents the msg field of a reject event
type RejectMsg struct {
	MsgCommon
	OpensClicks
}

// InboundEvent is a message inbound event
type InboundEvent struct{}
