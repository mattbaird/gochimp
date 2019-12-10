package mandrill

import (
	"context"
	"fmt"

	"github.com/lusis/gochimp/mandrill/api"
)

// MetaData represent a single metadata element in the mandrill api
type MetaData struct {
	Name         string
	State        string
	ViewTemplate string
}

// ListMetaData lists all the metadata from the mandrill api
func (c *Client) ListMetaData() ([]*MetaData, error) {
	return c.ListMetaDataContext(context.TODO())
}

// ListMetaDataContext lists all the metadata from the mandrill api
func (c *Client) ListMetaDataContext(ctx context.Context) ([]*MetaData, error) {
	req := &api.MetaDataListRequest{}
	resp := &api.MetaDataListResponse{}
	if err := c.postContext(ctx, "metadata/list", req, resp); err != nil {
		return nil, err
	}
	md := make([]*MetaData, len(*resp))
	c.debugLog(fmt.Sprintf("length of response: %d", len(md)))
	for i, m := range *resp {
		tempMD := &MetaData{
			Name:         m.Name,
			State:        m.State,
			ViewTemplate: m.ViewTemplate,
		}
		md[i] = tempMD
		c.debugLog(fmt.Sprintf("Length: %d", len(md)))
	}
	return md, nil
}

// DeleteMetaData deletes the named metadata
func (c *Client) DeleteMetaData(md string) error {
	return c.DeleteMetaDataContext(context.TODO(), md)
}

// DeleteMetaDataContext deletes the named metadata
func (c *Client) DeleteMetaDataContext(ctx context.Context, md string) error {
	req := &api.MetaDataDeleteRequest{
		Name: md,
	}
	resp := &api.MetaDataDeleteResponse{}
	return c.postContext(ctx, "metadata/delete", req, resp)
}

// Delete deletes the current instance of MetaData
func (md *MetaData) Delete() error {
	return globalClient.DeleteMetaData(md.Name)
}
