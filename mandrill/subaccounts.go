package mandrill

import (
	"context"
	"time"
)

// SubAccount represents a subaccount in mandrill
type SubAccount struct {
	ID             string
	Name           string
	CustomQuota    int32
	Status         string
	Reputation     int32
	CreatedAt      time.Time
	FirstSendAt    time.Time
	SentWeekly     int32
	SentMonthly    int32
	SentTotal      int32
	SentHourly     int32
	Notes          string
	LastThirtyDays map[string]int32
}

// SubAccountsList lists subaccounts
func (c *Client) SubAccountsList(q string) ([]*SubAccount, error) {
	return nil, nil
}

// SubAccountsListContext lists subaccounts with the provided context
func (c *Client) SubAccountsListContext(ctx context.Context, q string) ([]*SubAccount, error) {
	return nil, nil
}

// SubAccountsAdd adds a subaccount
func (c *Client) SubAccountsAdd(a SubAccount) (*SubAccount, error) {
	return nil, nil
}

// SubAccountsAddContext adds a subaccount with the provided context
func (c *Client) SubAccountsAddContext(ctx context.Context, a SubAccount) (*SubAccount, error) {
	return nil, nil
}

// SubAccountsInfo returns the info for the subaccount
func (c *Client) SubAccountsInfo(id string) (*SubAccount, error) {
	return nil, nil
}

// SubAccountsInfoContext returns the info for the subaccount with the provided context
func (c *Client) SubAccountsInfoContext(ctx context.Context, id string) (*SubAccount, error) {
	return nil, nil
}

// SubAccountsUpdate updates the subaccount
func (c *Client) SubAccountsUpdate(a *SubAccount) error {
	return nil
}

// SubAccountsUpdateContext updates the subaccount with the provided context
func (c *Client) SubAccountsUpdateContext(ctx context.Context, a *SubAccount) error {
	return nil
}

// SubAccountsDelete deletes the subaccount
func (c *Client) SubAccountsDelete(id string) error {
	return nil
}

// SubAccountsDeleteContext deletes the subaccount with the provided context
func (c *Client) SubAccountsDeleteContext(ctx context.Context, id string) error {
	return nil
}

// SubAccountsPause pauses the subaccount
func (c *Client) SubAccountsPause(id string) error {
	return nil
}

// SubAccountsPauseContext pauses the subaccount with the provided context
func (c *Client) SubAccountsPauseContext(ctx context.Context, id string) error {
	return nil
}

// SubAccountsResume resumes the subaccount
func (c *Client) SubAccountsResume(id string) error {
	return nil
}

// SubAccountsResumeContext resumes the subaccount with context
func (c *Client) SubAccountsResumeContext(ctx context.Context, id string) error {
	return nil
}
