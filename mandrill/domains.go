package mandrill

import (
	"time"
)

type SendingDomain struct {
	Name         string
	CreatedAt    time.Time
	LastTestedAt time.Time
	SPF          SPF
	DKIM         DKIM
	VerifiedAt   time.Time
	ValidSigning bool
}

func (d *SendingDomain) IsValid() bool {
	if d.SPF.Valid && d.DKIM.Valid && d.ValidSigning {
		return true
	}
	return false
}

type SPF struct {
	Valid      bool
	ValidAfter time.Time
	Error      string
}

type DKIM struct {
	Valid      bool
	ValidAfter time.Time
	Error      string
}

type InboundDomain struct {
	Name      string
	CreatedAt time.Time
	ValidMX   bool
}

type CName struct {
	Valid      bool
	ValidAfter time.Time
	Error      string
}

type TrackingDomain struct {
	Domain        string
	CreatedAt     time.Time
	LastTestedAt  time.Time
	CName         CName
	ValidTracking bool
}

func (d *TrackingDomain) IsValid() (bool, error) {
	res, err := globalClient.CheckTrackingDomain(d.Domain)
	if err != nil {
		return false, err
	}
	if res.CName.Valid && res.ValidTracking && time.Now().After(res.CName.ValidAfter) && res.CName.Error == "" {
		return true, nil
	}
	return false, nil
}
