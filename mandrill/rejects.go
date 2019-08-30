package mandrill

import (
	"fmt"
	"time"

	"github.com/lusis/gochimp/mandrill/api"
)

// Reject represents a rejected email
type Reject struct {
	Email       string
	Reason      string
	Detail      string
	CreatedAt   time.Time
	LastEventAt time.Time
	ExpiresAt   time.Time
	Expired     bool
	Sender      Sender
	SubAccount  string
}

// AddReject adds an explicit rejection for the provided email
func (c *Client) AddReject(email string, comment string) error {
	req := &api.RejectsAddRequest{
		Email: email,
	}

	if c.subaccount != "" {
		req.SubAccount = c.subaccount
	}

	if comment != "" {
		req.Comment = comment
	}
	resp := &api.RejectsAddResponse{}
	err := c.post("rejects/add", req, resp)
	if err != nil || !resp.Added {
		return err
	}
	return nil
}

// ListRejects returns a list of all current Rejects
func (c *Client) ListRejects(email string, includeExpired bool) ([]*Reject, error) {
	req := &api.RejectsListRequest{
		Email:          email,
		IncludeExpired: includeExpired,
	}
	if c.subaccount != "" {
		req.SubAccount = c.subaccount
	}
	resp := &api.RejectsListResponse{}
	err := c.post("rejects/list", req, resp)
	if err != nil {
		return nil, err
	}
	rejects := make([]*Reject, len(*resp))
	for _, r := range *resp {
		reject := &Reject{
			Email:       r.Email,
			Reason:      r.Reason,
			Detail:      r.Detail,
			CreatedAt:   r.CreatedAt.Time,
			LastEventAt: r.LastEventAt.Time,
			ExpiresAt:   r.ExpiresAt.Time,
			Expired:     r.Expired,
			Sender: Sender{
				Address:      r.Sender.Address,
				CreatedAt:    r.Sender.CreatedAt.Time,
				Sent:         r.Sender.Sent,
				HardBounces:  r.Sender.HardBounces,
				SoftBounces:  r.Sender.SoftBounces,
				Rejects:      r.Sender.Rejects,
				Complaints:   r.Sender.Complaints,
				Unsubs:       r.Sender.Unsubs,
				Opens:        r.Sender.Opens,
				Clicks:       r.Sender.Clicks,
				UniqueOpens:  r.Sender.UniqueOpens,
				UniqueClicks: r.Sender.UniqueClicks,
			},
		}

		rejects = append(rejects, reject)
	}
	return rejects, nil
}

// Delete deletes a reject
func (r *Reject) Delete() error {
	req := &api.RejectsDeleteRequest{
		Email: r.Email,
	}
	if globalClient.subaccount != "" {
		req.SubAccount = globalClient.subaccount
	}
	resp := &api.RejectsDeleteResponse{}
	err := globalClient.post("rejects/delete", req, resp)
	if err != nil {
		return err
	}
	if !resp.Deleted {
		return fmt.Errorf("api call was successful but deleted reported as false")
	}
	return nil
}
