package mandrill

import (
	"time"

	"github.com/lusis/gochimp/mandrill/api"
)

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
	Stats        Stats
}

func (s *Sender) Info() (*Sender, error) {
	return globalClient.GetSenderInfo(s.Address)
}

func (s *Sender) TimeSeries() ([]TimeSeries, error) {
	return globalClient.GetSenderTimeSeries(s.Address)
}

func (c *Client) GetSenderTimeSeries(address string) ([]TimeSeries, error) {
	req := &api.SendersTimeSeriesRequest{
		Address: address,
	}
	resp := &api.SendersTimeSeriesResponse{}
	err := c.post("senders/time-series", req, resp)
	if err != nil {
		return nil, err
	}
	tsData := make([]TimeSeries, len(*resp))
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

func (c *Client) GetSenderInfo(address string) (*Sender, error) {
	req := &api.SendersInfoRequest{
		Address: address,
	}
	resp := &api.SendersInfoResponse{}
	err := c.post("senders/info", req, resp)
	if err != nil {
		return nil, err
	}
	sender := &Sender{
		Address:     resp.Address,
		CreatedAt:   resp.CreatedAt,
		Sent:        resp.Sent,
		HardBounces: resp.HardBounces,
		SoftBounces: resp.SoftBounces,
		Rejects:     resp.Rejects,
		Complaints:  resp.Complaints,
		Unsubs:      resp.Unsubs,
		Opens:       resp.Opens,
		Clicks:      resp.Clicks,
		Stats:       resp.Stats,
	}
	return sender, nil
}

func (c *Client) ListSenders() ([]*Sender, error) {
	req := &api.SendersListRequest{}
	resp := &api.SendersListResponse{}
	err := c.post("senders/list", req, resp)
	if err != nil {
		return nil, err
	}
	senders := make([]*Sender, len(*resp))
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
