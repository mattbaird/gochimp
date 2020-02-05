package events

import (
	"encoding/json"
	"strings"

	api "github.com/lusis/gochimp/mandrill/api/events"
)

func parseMessageEvent(e WebhookEvent) (MessageEvent, error) {
	me := MessageEvent{}
	if e.Type != MessageEventType {
		return me, InvalidEventType{eventType: e.Type}
	}
	evt, ok := messageEventMapping[e.InnerEventType]
	if !ok {
		return me, InvalidEventType{eventType: e.InnerEventType}
	}
	apiEvt, ok := jsonMessageEventMapping[e.InnerEventType]
	if !ok {
		return me, InvalidEventType{eventType: e.InnerEventType}
	}
	decoder := json.NewDecoder(strings.NewReader(string(e.raw)))
	decoder.DisallowUnknownFields()
	switch d := apiEvt.(type) {
	case api.SendMessageEvent:
		if err := decoder.Decode(&d); err != nil {
			return me, UnmarshallError{data: string(e.raw), msg: err.Error()}
		}
		e := evt.(SendEvent)
		e.parse(d)
		me.Data = e
	case api.DeferralMessageEvent:
		if err := decoder.Decode(&d); err != nil {
			return me, UnmarshallError{data: string(e.raw), msg: err.Error()}
		}
		e := evt.(DeferralEvent)
		e.parse(d)
		me.Data = e
	case api.HardBounceMessageEvent:
		if err := decoder.Decode(&d); err != nil {
			return me, UnmarshallError{data: string(e.raw), msg: err.Error()}
		}
		e := evt.(HardBounceEvent)
		e.parse(d)
		me.Data = e
	case api.SoftBounceMessageEvent:
		if err := decoder.Decode(&d); err != nil {
			return me, UnmarshallError{data: string(e.raw), msg: err.Error()}
		}
		e := evt.(SoftBounceEvent)
		e.parse(d)
		me.Data = e
	case api.OpenMessageEvent:
		if err := decoder.Decode(&d); err != nil {
			return me, UnmarshallError{data: string(e.raw), msg: err.Error()}
		}
		e := evt.(OpenEvent)
		e.parse(d)
		me.Data = e
	case api.ClickMessageEvent:
		if err := decoder.Decode(&d); err != nil {
			return me, UnmarshallError{data: string(e.raw), msg: err.Error()}
		}
		e := evt.(ClickEvent)
		e.parse(d)
		me.Data = e
	case api.SpamMessageEvent:
		if err := decoder.Decode(&d); err != nil {
			return me, UnmarshallError{data: string(e.raw), msg: err.Error()}
		}
		e := evt.(SpamEvent)
		e.parse(d)
		me.Data = e
	case api.UnsubMessageEvent:
		if err := decoder.Decode(&d); err != nil {
			return me, UnmarshallError{data: string(e.raw), msg: err.Error()}
		}
		e := evt.(UnsubEvent)
		e.parse(d)
		me.Data = e
	case api.RejectMessageEvent:
		if err := decoder.Decode(&d); err != nil {
			return me, UnmarshallError{data: string(e.raw), msg: err.Error()}
		}
		e := evt.(RejectEvent)
		e.parse(d)
		me.Data = e
	}
	return me, nil
}

func (s *SendEvent) parse(e api.SendMessageEvent) {
	s.Event = e.Event
	s.Timestamp = e.TS.Time
	s.ID = e.ID
	s.Msg = SendMsg{
		MsgCommon{
			Timestamp: e.Msg.TS.Time,
			ID:        e.Msg.ID,
			Version:   e.Msg.Version,
			Subject:   e.Msg.Subject,
			Email:     e.Msg.Email,
			Sender:    e.Msg.Sender,
			Tags:      e.Msg.Tags,
			State:     e.Msg.State,
			MetaData:  e.Msg.MetaData,
			Template:  e.Msg.Template,
		},
		OpensClicks{},
	}
	for _, o := range e.Msg.Opens {
		s.Msg.Opens = append(s.Msg.Opens, Open{Timestamp: o.TS.Time})
	}
	for _, c := range e.Msg.Clicks {
		s.Msg.Clicks = append(s.Msg.Clicks, Click{Timestamp: c.TS.Time, URL: c.URL})
	}
}

func (s *DeferralEvent) parse(e api.DeferralMessageEvent) {
	s.Event = e.Event
	s.Timestamp = e.TS.Time
	s.ID = e.ID
	s.Msg = DeferralMsg{}
	s.Msg.MsgCommon =
		MsgCommon{
			Timestamp: e.Msg.TS.Time,
			ID:        e.Msg.ID,
			Version:   e.Msg.Version,
			Subject:   e.Msg.Subject,
			Email:     e.Msg.Email,
			Sender:    e.Msg.Sender,
			Tags:      e.Msg.Tags,
			State:     e.Msg.State,
			MetaData:  e.Msg.MetaData,
			Template:  e.Msg.Template,
		}
	for _, se := range e.Msg.SMTPEvents {
		s.Msg.SMTPEvents = append(s.Msg.SMTPEvents, SMTPEvent{
			Timestamp:     se.TS.Time,
			DestinationIP: se.DestinationIP,
			Diag:          se.Diag,
			SourceIP:      se.SourceIP,
			Type:          se.Type,
			Size:          se.Size,
		})
	}
	for _, o := range e.Msg.Opens {
		s.Msg.Opens = append(s.Msg.Opens, Open{Timestamp: o.TS.Time})
	}
	for _, c := range e.Msg.Clicks {
		s.Msg.Clicks = append(s.Msg.Clicks, Click{Timestamp: c.TS.Time, URL: c.URL})
	}
}

func (s *HardBounceEvent) parse(e api.HardBounceMessageEvent) {
	s.Event = e.Event
	s.Timestamp = e.TS.Time
	s.ID = e.ID
	s.Msg = BounceMsg{}
	s.Msg.MsgCommon =
		MsgCommon{
			Timestamp: e.Msg.TS.Time,
			ID:        e.Msg.ID,
			Version:   e.Msg.Version,
			Subject:   e.Msg.Subject,
			Email:     e.Msg.Email,
			Sender:    e.Msg.Sender,
			Tags:      e.Msg.Tags,
			State:     e.Msg.State,
			MetaData:  e.Msg.MetaData,
			Template:  e.Msg.Template,
		}
	s.Msg.BGToolsCode = e.Msg.BGToolsCode
	s.Msg.BounceDescription = e.Msg.BounceDescription
	s.Msg.Diag = e.Msg.Diag
}
func (s *SoftBounceEvent) parse(e api.SoftBounceMessageEvent) {
	s.Event = e.Event
	s.Timestamp = e.TS.Time
	s.ID = e.ID
	s.Msg = BounceMsg{}
	s.Msg.MsgCommon =
		MsgCommon{
			Timestamp: e.Msg.TS.Time,
			ID:        e.Msg.ID,
			Version:   e.Msg.Version,
			Subject:   e.Msg.Subject,
			Email:     e.Msg.Email,
			Sender:    e.Msg.Sender,
			Tags:      e.Msg.Tags,
			State:     e.Msg.State,
			MetaData:  e.Msg.MetaData,
			Template:  e.Msg.Template,
		}
	s.Msg.BGToolsCode = e.Msg.BGToolsCode
	s.Msg.BounceDescription = e.Msg.BounceDescription
	s.Msg.Diag = e.Msg.Diag
}

func (s *OpenEvent) parse(e api.OpenMessageEvent) {
	s.Event = e.Event
	s.Timestamp = e.TS.Time
	s.ID = e.ID
	s.Msg = OpenMsg{}
	s.IP = e.IP
	s.Location = Location(e.Location)
	s.UserAgentParsed = UserAgentParsed(e.UserAgentParsed)
	s.Msg.MsgCommon =
		MsgCommon{
			Timestamp: e.Msg.TS.Time,
			ID:        e.Msg.ID,
			Version:   e.Msg.Version,
			Subject:   e.Msg.Subject,
			Email:     e.Msg.Email,
			Sender:    e.Msg.Sender,
			Tags:      e.Msg.Tags,
			State:     e.Msg.State,
			MetaData:  e.Msg.MetaData,
			Template:  e.Msg.Template,
		}
	for _, o := range e.Msg.Opens {
		s.Msg.Opens = append(s.Msg.Opens, Open{Timestamp: o.TS.Time})
	}
	for _, c := range e.Msg.Clicks {
		s.Msg.Clicks = append(s.Msg.Clicks, Click{Timestamp: c.TS.Time, URL: c.URL})
	}
}
func (s *ClickEvent) parse(e api.ClickMessageEvent) {
	s.Event = e.Event
	s.Timestamp = e.TS.Time
	s.ID = e.ID
	s.Msg = ClickMsg{}
	s.IP = e.IP
	s.Location = Location(e.Location)
	s.UserAgentParsed = UserAgentParsed(e.UserAgentParsed)
	s.Msg.MsgCommon =
		MsgCommon{
			Timestamp: e.Msg.TS.Time,
			ID:        e.Msg.ID,
			Version:   e.Msg.Version,
			Subject:   e.Msg.Subject,
			Email:     e.Msg.Email,
			Sender:    e.Msg.Sender,
			Tags:      e.Msg.Tags,
			State:     e.Msg.State,
			MetaData:  e.Msg.MetaData,
			Template:  e.Msg.Template,
		}
	for _, o := range e.Msg.Opens {
		s.Msg.Opens = append(s.Msg.Opens, Open{Timestamp: o.TS.Time})
	}
	for _, c := range e.Msg.Clicks {
		s.Msg.Clicks = append(s.Msg.Clicks, Click{Timestamp: c.TS.Time, URL: c.URL})
	}
}
func (s *SpamEvent) parse(e api.SpamMessageEvent) {
	s.Event = e.Event
	s.Timestamp = e.TS.Time
	s.ID = e.ID
	s.Msg = SpamMsg{
		MsgCommon{
			Timestamp: e.Msg.TS.Time,
			ID:        e.Msg.ID,
			Version:   e.Msg.Version,
			Subject:   e.Msg.Subject,
			Email:     e.Msg.Email,
			Sender:    e.Msg.Sender,
			Tags:      e.Msg.Tags,
			State:     e.Msg.State,
			MetaData:  e.Msg.MetaData,
			Template:  e.Msg.Template,
		},
		OpensClicks{},
	}
	for _, o := range e.Msg.Opens {
		s.Msg.Opens = append(s.Msg.Opens, Open{Timestamp: o.TS.Time})
	}
	for _, c := range e.Msg.Clicks {
		s.Msg.Clicks = append(s.Msg.Clicks, Click{Timestamp: c.TS.Time, URL: c.URL})
	}
}
func (s *UnsubEvent) parse(e api.UnsubMessageEvent) {
	s.Event = e.Event
	s.Timestamp = e.TS.Time
	s.ID = e.ID
	s.Msg = UnsubMsg{
		MsgCommon{
			Timestamp: e.Msg.TS.Time,
			ID:        e.Msg.ID,
			Version:   e.Msg.Version,
			Subject:   e.Msg.Subject,
			Email:     e.Msg.Email,
			Sender:    e.Msg.Sender,
			Tags:      e.Msg.Tags,
			State:     e.Msg.State,
			MetaData:  e.Msg.MetaData,
			Template:  e.Msg.Template,
		},
		OpensClicks{},
	}
	for _, o := range e.Msg.Opens {
		s.Msg.Opens = append(s.Msg.Opens, Open{Timestamp: o.TS.Time})
	}
	for _, c := range e.Msg.Clicks {
		s.Msg.Clicks = append(s.Msg.Clicks, Click{Timestamp: c.TS.Time, URL: c.URL})
	}
}
func (s *RejectEvent) parse(e api.RejectMessageEvent) {
	s.Event = e.Event
	s.Timestamp = e.TS.Time
	s.ID = e.ID
	s.Msg = RejectMsg{
		MsgCommon{
			Timestamp: e.Msg.TS.Time,
			ID:        e.Msg.ID,
			Version:   e.Msg.Version,
			Subject:   e.Msg.Subject,
			Email:     e.Msg.Email,
			Sender:    e.Msg.Sender,
			Tags:      e.Msg.Tags,
			State:     e.Msg.State,
			MetaData:  e.Msg.MetaData,
			Template:  e.Msg.Template,
		},
		OpensClicks{},
	}
	for _, o := range e.Msg.Opens {
		s.Msg.Opens = append(s.Msg.Opens, Open{Timestamp: o.TS.Time})
	}
	for _, c := range e.Msg.Clicks {
		s.Msg.Clicks = append(s.Msg.Clicks, Click{Timestamp: c.TS.Time, URL: c.URL})
	}
}
