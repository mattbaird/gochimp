package mandrill

import (
	"context"

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
	Stats        map[string]Stats
}

// ListTags lists all tags
func (c *Client) ListTags() ([]*Tag, error) {
	return c.ListTagsContext(context.TODO())
}

// ListTagsContext lists all tags
func (c *Client) ListTagsContext(ctx context.Context) ([]*Tag, error) {
	req := &api.TagsListRequest{}
	resp := &api.TagsListResponse{}
	err := c.postContext(ctx, "tags/list", req, resp)
	if err != nil {
		return nil, err
	}
	tags := []*Tag{}
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

// TagInfo return details about the named tag
func (c *Client) TagInfo(t string) (*Tag, error) {
	return c.TagInfoContext(context.TODO(), t)
}

// TagInfoContext returns details about the named tag with the provided context
func (c *Client) TagInfoContext(ctx context.Context, t string) (*Tag, error) {
	req := &api.TagsInfoRequest{
		Tag: t,
	}
	resp := &api.TagsInfoResponse{}
	err := c.postContext(ctx, "tags/info", req, resp)
	if err != nil {
		return nil, err
	}
	tag := &Tag{
		Name:        resp.Tag,
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
		tag.Stats[k] = statsResponseToStats(v)
	}
	return tag, nil
}

// DeleteTag deletes the provided tag by name
func (c *Client) DeleteTag(t string) error {
	return c.DeleteTagContext(context.TODO(), t)
}

// DeleteTagContext context deletes the provided tag by name with provided context
func (c *Client) DeleteTagContext(ctx context.Context, t string) error {
	req := &api.TagsDeleteRequest{
		Tag: t,
	}
	resp := &api.TagsDeleteResponse{}
	err := c.postContext(ctx, "tags/delete", req, resp)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes the current Tag
func (t *Tag) Delete() error {
	return globalClient.DeleteTag(t.Name)
}
