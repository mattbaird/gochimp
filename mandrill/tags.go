package mandrill

import (
	"github.com/lusis/gochimp/mandrill/api"
)

// Tag represents a tag and its stats
type Tag struct {
	Name         string
	Reputation   int32
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
}

// ListTags lists all tags
func (c *Client) ListTags() ([]*Tag, error) {
	req := &api.TagsListRequest{}
	resp := &api.TagsListResponse{}
	err := c.post("tags/list", req, resp)
	if err != nil {
		return nil, err
	}
	tags := make([]*Tag, len(*resp))
	for _, t := range *resp {
		tags = append(tags, &Tag{
			Name:         t.Tag,
			Reputation:   t.Reputation,
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
	return tags, nil
}

// Delete deletes the current Tag
func (t *Tag) Delete() error {
	req := &api.TagsDeleteRequest{
		Tag: t.Name,
	}
	resp := &api.TagsDeleteResponse{}
	err := globalClient.post("tags/delete", req, resp)
	if err != nil {
		return err
	}
	return nil
}
