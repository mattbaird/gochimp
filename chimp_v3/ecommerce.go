package gochimp

import (
	"errors"
	"fmt"
)

const (
	store_path  = "/ecommerce/stores/%s"
	stores_path = "/ecommerce/stores"

	cart_path  = "/ecommerce/store/%s/cart/%s"
	carts_path = "/ecommerce/store/%s/cart"

	order_path  = "/ecommerce/store/%s/order/%s"
	orders_path = "/ecommerce/store/%s/order"

	product_path  = "/ecommerce/stores/%s/products/%s"
	products_path = "/ecommerce/stores/%s/products"

	variant_path  = "/ecommerce/stores/%s/products/%s/variants/%s"
	variants_path = "/ecommerce/stores/%s/products/%s/variants"
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

func (store Store) CanMakeRequest() error {
	if store.ID == "" {
		return errors.New("No ID provided on store")
	}

	return nil
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
	response := new(Store)

	endpoint := fmt.Sprintf(store_path, id)
	err := api.Request("GET", endpoint, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (api ChimpAPI) CreateStore(req *Store) (*Store, error) {
	res := new(Store)
	res.api = &api

	return res, api.Request("POST", stores_path, nil, req, res)
}

func (api ChimpAPI) UpdateStore(req *Store) (*Store, error) {
	res := new(Store)
	res.api = &api

	endpoint := fmt.Sprintf(store_path, req.ID)
	return res, api.Request("PATCH", endpoint, nil, req, res)
}

func (api ChimpAPI) DeleteStore(id string) (bool, error) {
	endpoint := fmt.Sprintf(store_path, id)
	return api.RequestOk("DELETE", endpoint)
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

	// Required
	Customer     Customer   `json:"customer"`
	CurrencyCode string     `json:"currency_code"`
	OrderTotal   float64    `json:"order_total"`
	Lines        []LineItem `json:"lines"`

	// Optional
	ID          string  `json:"id"`
	CampaignID  string  `json:"campaign_id"`
	CheckoutURL string  `json:"checkout_url"`
	TaxTotal    float64 `json:"tax_total"`

	// Response only
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Links     []Link `json:"_links"`
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

func (store Store) CreateCart(req *Cart) (*Cart, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(carts_path, store.ID)
	res := new(Cart)

	return res, store.api.Request("POST", endpoint, nil, req, res)
}

func (store Store) UpdateCart(req *Cart) (*Cart, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(cart_path, store.ID, req.ID)
	res := new(Cart)

	return res, store.api.Request("PATCH", endpoint, nil, req, res)
}

func (store Store) DeleteCart(id string) (bool, error) {
	if err := store.CanMakeRequest(); err != nil {
		return false, err
	}

	endpoint := fmt.Sprintf(cart_path, store.ID, id)
	return store.api.RequestOk("DELETE", endpoint)
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

	// Required
	ID           string     `json:"id"`
	Customer     Customer   `json:"customer"`
	Lines        []LineItem `json:"lines"`
	CurrencyCode string     `json:"currency_code"`
	OrderTotal   float64    `json:"order_total"`

	// Optional
	TaxTotal           float64 `json:"tax_total"`
	ShippingTotal      float64 `json:"shipping_total"`
	TrackingCode       string  `json:"tracking_code"`
	ProcessedAtForeign string  `json:processed_at_foreign`
	CancelledAtForeign string  `json:cancelled_at_foreign`
	UpdatedAtForeign   string  `json:updated_at_foreign`
	CampaignID         string  `json:"campaign_id"`
	FinancialStatus    string  `json:"financial_status"`
	FulfillmentStatus  string  `json:"fulfillment_status"`

	BillingAddress  Address `json:"billing_address"`
	ShippingAddress Address `json:"shipping_address"`

	// Response only
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Links     []Link `json:"_links"`
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

	endpoint := fmt.Sprintf(order_path, store.ID, id)
	err := store.api.Request("GET", endpoint, params, nil, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (store Store) CreateOrder(req *Order) (*Order, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(orders_path, store.ID)
	res := new(Order)

	return res, store.api.Request("POST", endpoint, nil, req, res)
}

func (store Store) UpdateOrder(req *Order) (*Order, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(order_path, store.ID, req.ID)
	res := new(Order)

	return res, store.api.Request("PATCH", endpoint, nil, req, res)
}

func (store Store) DeleteOrder(id string) (bool, error) {
	if err := store.CanMakeRequest(); err != nil {
		return false, err
	}

	endpoint := fmt.Sprintf(order_path, store.ID, id)
	return store.api.RequestOk("DELETE", endpoint)
}

// ------------------------------------------------------------------------------------------------
// Products
// ------------------------------------------------------------------------------------------------
type Product struct {
	APIError

	api *ChimpAPI

	StoreID     string
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title"`
	Handle      string    `json:"handle"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Vendor      string    `json:"vendor"`
	ImageURL    string    `json:"image_url"`
	Variants    []Variant `json:"variants"`
	PublishedAt string    `json:"published_at_foreign"`
	Links       []Link    `json:"_links"`
}

func (product Product) CanMakeRequest() error {
	if product.ID == "" || product.StoreID == "" {
		return errors.New("No ID provided on product")
	}

	return nil
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

func (store Store) CreateProduct(req *Product) (*Product, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(products_path, store.ID)
	res := new(Product)
	res.api = store.api

	return res, store.api.Request("POST", endpoint, nil, req, res)
}

func (store Store) UpdateProduct(req *Product) (*Product, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(product_path, store.ID, req.ID)
	res := new(Product)
	res.api = store.api

	return res, store.api.Request("PATCH", endpoint, nil, req, res)
}

func (store Store) DeleteProduct(id string) (bool, error) {
	if err := store.CanMakeRequest(); err != nil {
		return false, err
	}

	endpoint := fmt.Sprintf(product_path, store.ID, id)
	return store.api.RequestOk("DELETE", endpoint)
}

// ------------------------------------------------------------------------------------------------
// Variants
// ------------------------------------------------------------------------------------------------
type Variant struct {
	APIError

	api *ChimpAPI

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

func (product Product) CreateVariant(req *Variant) (*Variant, error) {
	if err := product.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(variants_path, product.StoreID, product.ID)
	res := new(Variant)
	res.api = product.api

	return res, product.api.Request("POST", endpoint, nil, req, res)
}

func (product Product) UpdateVariant(req *Variant) (*Variant, error) {
	if err := product.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(variant_path, product.StoreID, product.ID, req.ID)
	res := new(Variant)
	res.api = product.api

	return res, product.api.Request("PATCH", endpoint, nil, req, res)
}

func (product Product) DeleteVariant(id string) (bool, error) {
	if err := product.CanMakeRequest(); err != nil {
		return false, err
	}

	endpoint := fmt.Sprintf(variant_path, product.StoreID, product.ID, id)
	return product.api.RequestOk("DELETE", endpoint)
}
