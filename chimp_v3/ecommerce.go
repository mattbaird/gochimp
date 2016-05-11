package gochimp

import "fmt"

const (
	store_path  = "/ecommerce/stores/%s"
	stores_path = "/ecommerce/stores"

	cart_path  = "/ecommerce/store/%s/cart/%s"
	carts_path = "/ecommerce/store/%s/cart"
)

// ------------------------------------------------------------------------------------------------
// Stores
// ------------------------------------------------------------------------------------------------

type Store struct {
	APIError

	api *ChimpAPI

	ID            string  `json:"id"`
	ListID        string  `json:"list_id"`
	Name          string  `json:"name"`
	Platform      string  `json:"platform"`
	Domain        string  `json:"domain"`
	EmailAddress  string  `json:"email_address"`
	CurrencyCode  string  `json:"currency_code"`
	MoneyFormat   string  `json:"money_format"`
	PrimaryLocale string  `json:"primary_locale"`
	Timezone      string  `json:"timezone"`
	Phone         string  `json:"phone"`
	Address       Address `json:"address"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	Links         []Link  `json:"_links"`
}

type StoreList struct {
	APIError

	Stores     []Store `json:"stores"`
	TotalItems int     `json:"total_items"`
	Links      []Link  `json:"_link"`
}

func (api ChimpAPI) GetStores(params *ExtendedQueryParams) (*StoreList, error) {
	response := new(StoreList)
	err := api.Request("GET", stores_path, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (api ChimpAPI) GetStore(id string, params QueryParams) (*Store, error) {

	return nil, nil
}

// ------------------------------------------------------------------------------------------------
// Carts
// ------------------------------------------------------------------------------------------------

type CartList struct {
	APIError

	Carts      []Cart `json:"cart"`
	TotalItems int    `json:"total_items"`
	Links      []Link `json:"_links"`
}

type Cart struct {
	APIError

	ID           string         `json:"id"`
	Customer     Customer       `json:"customer"`
	CampaignID   string         `json:"campaign_id"`
	CheckoutURL  string         `json:"checkout_url"`
	CurrencyCode string         `json:"currency_code"`
	OrderTotal   int            `json:"order_total"` // Float?
	TaxTotal     int            `json:"tax_total"`   // Float?
	Lines        []CartLineItem `json:"lines"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
	Links        []Link         `json:"_links"`
}

type CartLineItem struct {
	ProductID           string `json:"product_id"`
	ProductTitle        string `json:"product_title"`
	ProductVariantID    string `json:"product_variant_id"`
	ProductVariantTitle string `json:"product_variant_title"`
	Quantity            int    `json:"quantity"`
	Price               int    `json:"price"` // float?
}

func (store Store) GetCarts(params *ExtendedQueryParams) (*CartList, error) {
	response := new(CartList)

	if store.HasError() {
		return nil, fmt.Errorf("The store has an error, can't process request")
	}
	endpoint := fmt.Sprintf(carts_path, store.ID)
	err := store.api.Request("GET", endpoint, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (store Store) GetCart(id string, params *BasicQueryParams) (*Cart, error) {
	response := new(Cart)

	if store.HasError() {
		return nil, fmt.Errorf("The store has an error, can't process request")
	}

	endpoint := fmt.Sprintf(cart_path, store.ID, id)
	err := store.api.Request("GET", endpoint, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
