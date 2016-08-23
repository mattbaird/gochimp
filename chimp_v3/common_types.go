package gochimp

import (
	"fmt"
	"strings"
)

// APIError is what the what the api returns on error
type APIError struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func (err APIError) String() string {
	return fmt.Sprintf("%d : %s : %s : %s", err.Status, err.Type, err.Title, err.Detail)
}
func (err APIError) Error() string {
	return err.String()
}

// HasError checks if this call had an error
func (err APIError) HasError() bool {
	return err.Type != ""
}

// QueryParams defines the different params
type QueryParams interface {
	Params() map[string]string
}

// ExtendedQueryParams includes a count and offset
type ExtendedQueryParams struct {
	BasicQueryParams

	Count  int
	Offset int
}

func (q ExtendedQueryParams) Params() map[string]string {
	m := q.BasicQueryParams.Params()
	m["count"] = fmt.Sprintf("%d", q.Count)
	m["offset"] = fmt.Sprintf("%d", q.Offset)
	return m
}

// BasicQueryParams basic filter queries
type BasicQueryParams struct {
	Fields        []string
	ExcludeFields []string
}

func (q BasicQueryParams) Params() map[string]string {
	return map[string]string{
		"fields":         strings.Join(q.Fields, ","),
		"exclude_fields": strings.Join(q.ExcludeFields, ","),
	}
}

type withLinks struct {
	Link []Link `json:"_link"`
}

type baseList struct {
	TotalItems int    `json:"total_items"`
	Links      []Link `json:"_links"`
}

// Link refereneces another object
type Link struct {
	Rel          string `json:"re"`
	Href         string `json:"href"`
	Method       string `json:"method"`
	TargetSchema string `json:"targetSchema"`
	Schema       string `json:"schema"`
}

// Address represents what it says
type Address struct {
	Address1     string  `json:"address1"`
	Address2     string  `json:"address2"`
	City         string  `json:"city"`
	Province     string  `json:"province"`
	ProvinceCode string  `json:"province_code"`
	PostalCode   string  `json:"postal_code"`
	Country      string  `json:"country"`
	CountryCode  string  `json:"country_code"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
}

// Customer defines a mailchimp customer
type Customer struct {
	ID           string  `json:"id"`
	EmailAddress string  `json:"email_address"`
	OptInStatus  bool    `json:"opt_in_status"`
	Company      string  `json:"company"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	OrdersCount  int     `json:"orders_count"`
	TotalSpent   int     `json:"total_spent"` // float
	Address      Address `json:"address"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	Links        []Link  `json:"_links"`
}

// LineItem defines a mailchimp cart or order line item
type LineItem struct {
	ProductID           string `json:"product_id"`
	ProductTitle        string `json:"product_title"`
	ProductVariantID    string `json:"product_variant_id"`
	ProductVariantTitle string `json:"product_variant_title"`
	Quantity            int    `json:"quantity"`
	Price               int    `json:"price"` // float?
}

// Contact defines a single contact
type Contact struct {
	Company     string `json:"customer"`
	Address1    string `json:"address1"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	State       string `json:"state"`
	Zip         string `json:"zip"`
	Country     string `json:"country"`
	PhoneNumber string `json:"phone"`
}
