package mandrill

import (
	"context"
	"time"
)

// SendingDomain represents a sending domain in Mandrill
type SendingDomain struct {
	Name         string
	CreatedAt    time.Time
	LastTestedAt time.Time
	SPF          SPF
	DKIM         DKIM
	VerifiedAt   time.Time
	ValidSigning bool
}

// IsValid is a helper to determine if a sending domain is valid for use
func (d *SendingDomain) IsValid() bool {
	if d.SPF.Valid && d.DKIM.Valid && d.ValidSigning {
		return true
	}
	return false
}

// SPF represents SPF record validity for a sending domain
type SPF struct {
	Valid      bool
	ValidAfter time.Time
	Error      string
}

// DKIM represents DKIM record validity for a sending domain
type DKIM struct {
	Valid      bool
	ValidAfter time.Time
	Error      string
}

// InboundDomain represents an inbound domain in Mandrill
type InboundDomain struct {
	Name      string
	CreatedAt time.Time
	ValidMX   bool
}

// CName represents CName validity for a tracking domain
type CName struct {
	Valid      bool
	ValidAfter time.Time
	Error      string
}

// TrackingDomain represents a tracking domain in Mandrill
type TrackingDomain struct {
	Domain        string
	CreatedAt     time.Time
	LastTestedAt  time.Time
	CName         CName
	ValidTracking bool
}

// IsValid is a helper for verifying proper configuration of a tracking domain
func (d *TrackingDomain) IsValid() (bool, error) {
	return d.IsValidContext(context.TODO())
}

// IsValidContext is a helper for verifying proper configuration of a tracking domain
func (d *TrackingDomain) IsValidContext(ctx context.Context) (bool, error) {
	res, err := globalClient.CheckTrackingDomainContext(ctx, d.Domain)
	if err != nil {
		return false, err
	}
	if res.CName.Valid && res.ValidTracking && time.Now().After(res.CName.ValidAfter) && res.CName.Error == "" {
		return true, nil
	}
	return false, nil
}
