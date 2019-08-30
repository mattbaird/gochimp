package mandrill

import (
	"time"

	"github.com/lusis/gochimp/mandrill/api"
)

// User is a Mandrill user
type User struct {
	Username    string
	CreatedAt   time.Time
	PublicID    string
	Reputation  int32
	HourlyQuota int32
	Backlog     int32
	Stats       Stats
}

// UserInfo returns the information about the API-connected user
func (c *Client) UserInfo() (*User, error) {
	userInfoReq := &api.UsersInfoRequest{}
	userInfoResp := &api.UsersInfoResponse{}
	err := c.post("users/info", userInfoReq, userInfoResp)
	if err != nil {
		return nil, err
	}

	return &User{
		Username:    userInfoResp.Username,
		CreatedAt:   userInfoResp.CreatedAt.Time,
		PublicID:    userInfoResp.PublicID,
		Reputation:  userInfoResp.Reputation,
		HourlyQuota: userInfoResp.HourlyQuota,
		Backlog:     userInfoResp.Backlog,
		Stats:       userInfoResp.Stats,
	}, nil
}

// Ping calls users/ping2 to validate connectivity
func (c *Client) Ping() error {
	req := &api.UsersPing2Request{}
	resp := &api.UsersPing2Response{}
	return c.post("users/ping2", req, resp)
}

// UserSenders calls users/senders
func (c *Client) UserSenders() ([]*Sender, error) {
	req := &api.UsersSendersRequest{}
	resp := &api.UsersSendersResponse{}
	err := c.post("users/senders", req, resp)
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
