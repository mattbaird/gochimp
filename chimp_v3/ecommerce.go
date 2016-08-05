package gochimp

import "fmt"

const (
	store_path  = "/ecommerce/stores/%s"
	stores_path = "/ecommerce/stores"

	cart_path  = "/ecommerce/store/%s/cart/%s"
	carts_path = "/ecommerce/store/%s/cart"

	product_path  = "/ecommerce/stores/%s/products/%s"
	products_path = "/ecommerce/stores/%s/products"

	variant_path   = "/ecommerce/stores/%s/products/%s/variants/%s"
	variants_paths = "/ecommerce/stores/%s/products/%s/variants"
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

	ID           string     `json:"id"`
	Customer     Customer   `json:"customer"`
	CampaignID   string     `json:"campaign_id"`
	CheckoutURL  string     `json:"checkout_url"`
	CurrencyCode string     `json:"currency_code"`
	OrderTotal   float64    `json:"order_total"`
	TaxTotal     float64    `json:"tax_total"`
	Lines        []LineItem `json:"lines"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
	Links        []Link     `json:"_links"`
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

// ------------------------------------------------------------------------------------------------
// Orders
// ------------------------------------------------------------------------------------------------

type OrderList struct {
	APIError

	Orders     []Order `json:"cart"`
	TotalItems int     `json:"total_items"`
	Links      []Link  `json:"_links"`
}

type Order struct {
	APIError

	ID           string     `json:"id"`
	Customer     Customer   `json:"customer"`
	CampaignID   string     `json:"campaign_id"`
	CheckoutURL  string     `json:"checkout_url"`
	CurrencyCode string     `json:"currency_code"`
	OrderTotal   float64    `json:"order_total"`
	TaxTotal     float64    `json:"tax_total"`
	Lines        []LineItem `json:"lines"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
	Links        []Link     `json:"_links"`
}

func (store Store) GetOrders(params *ExtendedQueryParams) (*OrderList, error) {
	response := new(OrderList)

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

func (store Store) GetOrder(id string, params *BasicQueryParams) (*Order, error) {
	response := new(Order)

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

// ------------------------------------------------------------------------------------------------
// Products
// ------------------------------------------------------------------------------------------------
type Product struct {
	APIError

	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Handle      string    `json:"handle"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Vendor      string    `json:"vendor"`
	ImageURL    string    `json:"image_url"`
	Variants    []Variant `json:"variants"`
	PublishedAt string    `json:"published_at_foreign"`
}

type ProductList struct {
	APIError

	StoreID    string    `json:"store_id"`
	Products   []Product `json:"products"`
	TotalItems int       `json:"total_items"`
	Links      []Link    `json:"_links"`
}

func (store Store) GetProducts(params *ExtendedQueryParams) (*ProductList, error) {
	response := new(ProductList)

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

func (store Store) GetProduct(id string, params *BasicQueryParams) (*Product, error) {
	response := new(Product)

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

// ------------------------------------------------------------------------------------------------
// Variants
// ------------------------------------------------------------------------------------------------
type Variant struct {
	APIError

	ID                string  `json:"id"`
	Title             string  `json:"title"`
	Url               string  `json:"url"`
	SKU               string  `json:"sku"`
	Price             float64 `json:"price"`
	InventoryQuantity int     `json:"inventory_quantity"`
	ImageUrl          string  `json:"image_url"`
	Backorders        string  `json:"backorders"`
	Visibility        string  `json:"visibility"`
}

type VariantList struct {
	APIError

	StoreID    string    `json:"store_id"`
	Variants   []Variant `json:"variants"`
	TotalItems int       `json:"total_items"`
	Links      []Link    `json:"_links"`
}
