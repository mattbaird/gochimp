package mandrill

import (
	"context"
	"github.com/lusis/gochimp/mandrill/api"
	"time"
)

// IPPool represents a pool of dedicated ip addresses
type IPPool struct {
	Name      string
	CreatedAt time.Time
	IPs       []DedicatedIP
}

// DedicatedIP represents a dedicated IP in mandrill
type DedicatedIP struct {
	IP        string
	CreatedAt time.Time
	Pool      string
	Domain    string
	CustomDNS CustomDNS
	WarmUp    WarmUp
}

// CustomDNS represents the custom DNS associated with a dedicated IP
type CustomDNS struct {
	Enabled bool
	Valid   bool
	Error   string
}

// WarmUp represents the warmup status of a dedicated ip
type WarmUp struct {
	WarmingUp bool
	StartAt   time.Time
	EndAt     time.Time
}

// IPListContext returns a list of dedicated ips with the provided context
func (c *Client) IPListContext(ctx context.Context) ([]*DedicatedIP, error) {
	req := &api.IPsListRequest{}
	resp := &api.IPsListResponse{}
	err := c.postContext(ctx, "ips/list", req, resp)
	if err != nil {
		return nil, err
	}
	res := []*DedicatedIP{}
	for _, i := range *resp {
		ip := &DedicatedIP{
			IP:        i.IP,
			CreatedAt: i.CreatedAt.Time,
			Pool:      i.Pool,
			Domain:    i.Domain,
			CustomDNS: CustomDNS{
				Enabled: i.CustomDNS.Enabled,
				Valid:   i.CustomDNS.Valid,
				Error:   i.CustomDNS.Error,
			},
			WarmUp: WarmUp{
				WarmingUp: i.WarmUp.WarmingUp,
				StartAt:   i.WarmUp.StartAt.Time,
				EndAt:     i.WarmUp.EndAt.Time,
			},
		}
		res = append(res, ip)
	}
	return res, nil
}

// IPList returns a list of dedicated ips
func (c *Client) IPList() ([]*DedicatedIP, error) {
	return c.IPListContext(context.TODO())
}

// IPInfo returns information about a dedicated ip
func (c *Client) IPInfo(ip string) (*DedicatedIP, error) {
	return c.IPInfoContext(context.TODO(), ip)
}

// IPInfoContext returns information about a dedicated ip with provided context
func (c *Client) IPInfoContext(ctx context.Context, ip string) (*DedicatedIP, error) {
	req := &api.IPsInfoRequest{
		IP: ip,
	}
	resp := &api.IPInfoResponse{}
	err := c.postContext(ctx, "ips/list", req, resp)
	if err != nil {
		return nil, err
	}
	i := &DedicatedIP{
		IP:        resp.IP,
		CreatedAt: resp.CreatedAt.Time,
		Pool:      resp.Pool,
		Domain:    resp.Domain,
		CustomDNS: CustomDNS{
			Enabled: resp.CustomDNS.Enabled,
			Valid:   resp.CustomDNS.Valid,
			Error:   resp.CustomDNS.Error,
		},
		WarmUp: WarmUp{
			WarmingUp: resp.WarmUp.WarmingUp,
			StartAt:   resp.WarmUp.StartAt.Time,
			EndAt:     resp.WarmUp.EndAt.Time,
		},
	}
	return i, nil
}

// IPProvision provisions a new dedicated IP in the provided pool (optionally with warmup)
// Returns the server-side time the ip was requested
func (c *Client) IPProvision(pool string, warmup bool) (*time.Time, error) {
	return c.IPProvisionContext(context.TODO(), pool, warmup)
}

// IPProvisionContext provisions a new dedicated IP in the provided pool (optionally with warmup) with the provided context
// Returns the server-side time the ip was requested
func (c *Client) IPProvisionContext(ctx context.Context, pool string, warmup bool) (*time.Time, error) {
	req := &api.IPsProvisionRequest{
		Pool:   pool,
		WarmUp: warmup,
	}
	resp := &api.IPsProvisionResponse{}
	err := c.postContext(ctx, "ips/provision", req, resp)
	if err != nil {
		return nil, err
	}
	return &resp.RequestedAt.Time, nil
}

// IPStartWarmUp begins the warmup process for a dedicated IP
func (c *Client) IPStartWarmUp(ip string) (*DedicatedIP, error) {
	return c.IPStartWarmUpContext(context.TODO(), ip)
}

// IPStartWarmUpContext begins the warmup process for a dedicated IP with the provided context
func (c *Client) IPStartWarmUpContext(ctx context.Context, ip string) (*DedicatedIP, error) {
	req := &api.IPsStartWarmUpRequest{
		IP: ip,
	}
	resp := &api.IPsStartWarmUpResponse{}
	err := c.postContext(ctx, "ips/start-warmup", req, resp)
	if err != nil {
		return nil, err
	}
	i := &DedicatedIP{
		IP:        resp.IP,
		CreatedAt: resp.CreatedAt.Time,
		Pool:      resp.Pool,
		Domain:    resp.Domain,
		CustomDNS: CustomDNS{
			Enabled: resp.CustomDNS.Enabled,
			Valid:   resp.CustomDNS.Valid,
			Error:   resp.CustomDNS.Error,
		},
		WarmUp: WarmUp{
			WarmingUp: resp.WarmUp.WarmingUp,
			StartAt:   resp.WarmUp.StartAt.Time,
			EndAt:     resp.WarmUp.EndAt.Time,
		},
	}
	return i, nil
}
