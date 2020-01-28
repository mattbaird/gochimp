package mandrill

import (
	"context"
	"time"

	"github.com/lusis/gochimp/mandrill/api"
)

// Sender represents a sender in mandrill (a sending email address)
type Sender struct {
	Address      string
	CreatedAt    time.Time
	Sent         int32
	HardBounces  int32
	SoftBounces  int32
	Rejects      int32
	Complaints   int32
	Unsubs       int32
	Opens        int32
	Clicks       int32
	UniqueOpens  int32
	UniqueClicks int32
	Stats        map[string]Stats
}

// Info returns information about the current Sender
func (s *Sender) Info() (*Sender, error) {
	return globalClient.GetSenderInfo(s.Address)
}

// TimeSeries returns time-series data about the current Sender
func (s *Sender) TimeSeries() ([]TimeSeries, error) {
	return globalClient.GetSenderTimeSeries(s.Address)
}

// GetSenderTimeSeries returns time-series data about the provided email address
func (c *Client) GetSenderTimeSeries(address string) ([]TimeSeries, error) {
	return c.GetSenderTimeSeriesContext(context.TODO(), address)
}

// GetSenderTimeSeriesContext returns time-series data about the provided email address
func (c *Client) GetSenderTimeSeriesContext(ctx context.Context, address string) ([]TimeSeries, error) {
	req := &api.SendersTimeSeriesRequest{
		Address: address,
	}
	resp := &api.SendersTimeSeriesResponse{}
	err := c.postContext(ctx, "senders/time-series", req, resp)
	if err != nil {
		return nil, err
	}
	tsData := []TimeSeries{}
	for _, t := range *resp {
		tsData = append(tsData, TimeSeries{
			Time:         t.Time.Time,
			Sent:         t.Sent,
			HardBounces:  t.HardBounces,
			SoftBounces:  t.SoftBounces,
			Rejects:      t.Rejects,
			Complaints:   t.Complaints,
			Unsubs:       t.Unsubs,
			Opens:        t.Opens,
			Clicks:       t.Clicks,
			UniqueOpens:  t.UniqueOpens,
			UniqueClicks: t.UniqueClicks,
		})
	}
	return tsData, nil
}

// GetSenderInfo returns information about the provided sender email
func (c *Client) GetSenderInfo(address string) (*Sender, error) {
	return c.GetSenderInfoContext(context.TODO(), address)
}

// GetSenderInfoContext returns information about the provided sender email
func (c *Client) GetSenderInfoContext(ctx context.Context, address string) (*Sender, error) {
	req := &api.SendersInfoRequest{
		Address: address,
	}
	resp := &api.SendersInfoResponse{}
	err := c.postContext(ctx, "senders/info", req, resp)
	if err != nil {
		return nil, err
	}
	sender := &Sender{
		Address:     resp.Address,
		CreatedAt:   resp.CreatedAt.Time,
		Sent:        resp.Sent,
		HardBounces: resp.HardBounces,
		SoftBounces: resp.SoftBounces,
		Rejects:     resp.Rejects,
		Complaints:  resp.Complaints,
		Unsubs:      resp.Unsubs,
		Opens:       resp.Opens,
		Clicks:      resp.Clicks,
		Stats:       make(map[string]Stats),
	}
	for k, v := range resp.Stats {
		sender.Stats[k] = statsResponseToStats(v)
	}
	return sender, nil
}

// ListSenders lists all senders
func (c *Client) ListSenders() ([]*Sender, error) {
	req := &api.SendersListRequest{}
	resp := &api.SendersListResponse{}
	err := c.post("senders/list", req, resp)
	if err != nil {
		return nil, err
	}
	senders := []*Sender{}
	for _, s := range *resp {
		senders = append(senders, &Sender{
			Address:      s.Address,
			CreatedAt:    s.CreatedAt.Time,
			Sent:         s.Sent,
			HardBounces:  s.HardBounces,
			SoftBounces:  s.SoftBounces,
			Rejects:      s.Rejects,
			Complaints:   s.Complaints,
			Unsubs:       s.Unsubs,
			Opens:        s.Opens,
			Clicks:       s.Clicks,
			UniqueOpens:  s.UniqueOpens,
			UniqueClicks: s.UniqueClicks,
		})
	}
	return senders, nil
}

// AddSendingDomain adds a new sending domain
func (c *Client) AddSendingDomain(domain string) (*SendingDomain, error) {
	req := &api.SendersAddDomainRequest{
		Domain: domain,
	}
	resp := &api.SendersAddDomainResponse{}
	err := c.post("senders/add-domain", req, resp)
	if err != nil {
		return nil, err
	}
	sDomain := &SendingDomain{
		Name:         resp.Domain,
		CreatedAt:    resp.CreatedAt.Time,
		LastTestedAt: resp.LastTestedAt.Time,
		SPF: SPF{
			Valid:      resp.SPF.Valid,
			ValidAfter: resp.SPF.ValidAfter.Time,
			Error:      resp.SPF.Error,
		},
		DKIM: DKIM{
			Valid:      resp.DKIM.Valid,
			ValidAfter: resp.DKIM.ValidAfter.Time,
			Error:      resp.DKIM.Error,
		},
		VerifiedAt:   resp.VerifiedAt.Time,
		ValidSigning: resp.ValidSigning,
	}
	return sDomain, nil
}

// CheckSendingDomain returns the status of the provided domain
func (c *Client) CheckSendingDomain(domain string) (*SendingDomain, error) {
	req := &api.SendersCheckDomainRequest{
		Domain: domain,
	}
	resp := &api.SendersCheckDomainResponse{}
	err := c.post("senders/check-domain", req, resp)
	if err != nil {
		return nil, err
	}
	sDomain := &SendingDomain{
		Name:         resp.Domain,
		CreatedAt:    resp.CreatedAt.Time,
		LastTestedAt: resp.LastTestedAt.Time,
		SPF: SPF{
			Valid:      resp.SPF.Valid,
			ValidAfter: resp.SPF.ValidAfter.Time,
			Error:      resp.SPF.Error,
		},
		DKIM: DKIM{
			Valid:      resp.DKIM.Valid,
			ValidAfter: resp.DKIM.ValidAfter.Time,
			Error:      resp.DKIM.Error,
		},
		VerifiedAt:   resp.VerifiedAt.Time,
		ValidSigning: resp.ValidSigning,
	}
	return sDomain, nil
}

// VerifySendingDomain verifies the provided domain by sending an email to the provided mailbox
func (c *Client) VerifySendingDomain(domain string, mailbox string) (string, string, error) {
	req := &api.SendersVerifyDomainRequest{
		Domain:  domain,
		Mailbox: mailbox,
	}
	resp := &api.SendersVerifyDomainResponse{}
	err := c.post("senders/verify-domain", req, resp)
	if err != nil {
		return "", "", err
	}
	return resp.Status, resp.Email, nil
}

// SendersDomains lists sending domains
func (c *Client) SendersDomains() ([]*SendingDomain, error) {
	return c.SendersDomainsContext(context.TODO())
}

// SendersDomainsContext lists sending domains with context
func (c *Client) SendersDomainsContext(ctx context.Context) ([]*SendingDomain, error) {
	req := &api.SendersDomainsRequest{}
	resp := &api.SendersDomainsResponse{}
	if err := c.postContext(ctx, "senders/domains", req, resp); err != nil {
		return nil, err
	}
	sendingDomains := []*SendingDomain{}
	for _, d := range *resp {
		sendingDomains = append(sendingDomains, &SendingDomain{
			Name:         d.Domain,
			CreatedAt:    d.CreatedAt.Time,
			LastTestedAt: d.LastTestedAt.Time,
			SPF: SPF{
				Valid:      d.SPF.Valid,
				ValidAfter: d.SPF.ValidAfter.Time,
				Error:      d.SPF.Error,
			},
			DKIM: DKIM{
				Valid:      d.DKIM.Valid,
				ValidAfter: d.DKIM.ValidAfter.Time,
				Error:      d.DKIM.Error,
			},
			VerifiedAt:   d.VerifiedAt.Time,
			ValidSigning: d.ValidSigning,
		})
	}
	return sendingDomains, nil
}
