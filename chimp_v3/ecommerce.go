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

	// Required
	ID           string `json:"id"`
	ListID       string `json:"list_id"`
	CurrencyCode string `json:"currency_code"`
	Name         string `json:"name"`

	// Optional
	Platform      string  `json:"platform,omitempty"`
	Domain        string  `json:"domain,omitempty"`
	EmailAddress  string  `json:"email_address,omitempty"`
	MoneyFormat   string  `json:"money_format,omitempty"`
	PrimaryLocale string  `json:"primary_locale,omitempty"`
	Timezone      string  `json:"timezone,omitempty"`
	Phone         string  `json:"phone,omitempty"`
	Address       Address `json:"address,omitempty"`

	// Response
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Links     []Link `json:"_links,omitempty"`
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
	Links      []Link  `json:"_links"`
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
	res := new(Store)
	res.api = &api

	endpoint := fmt.Sprintf(store_path, id)
	err := api.Request("GET", endpoint, params, nil, res)
	if err != nil {
		return nil, err
	}

	return res, nil
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
	ID          string  `json:"id,omitempty"`
	CampaignID  string  `json:"campaign_id,omitempty"`
	CheckoutURL string  `json:"checkout_url,omitempty"`
	TaxTotal    float64 `json:"tax_total,omitempty"`

	// Response only
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Links     []Link `json:"_links,omitempty"`
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
	Links      []Link  `json:"_links,omitempty"`
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
	TaxTotal           float64 `json:"tax_total,omitempty"`
	ShippingTotal      float64 `json:"shipping_total,omitempty"`
	TrackingCode       string  `json:"tracking_code,omitempty"`
	ProcessedAtForeign string  `json:processed_at_foreign`
	CancelledAtForeign string  `json:cancelled_at_foreign`
	UpdatedAtForeign   string  `json:updated_at_foreign`
	CampaignID         string  `json:"campaign_id,omitempty"`
	FinancialStatus    string  `json:"financial_status,omitempty"`
	FulfillmentStatus  string  `json:"fulfillment_status,omitempty"`

	BillingAddress  Address `json:"billing_address,omitempty"`
	ShippingAddress Address `json:"shipping_address,omitempty"`

	// Response only
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	Links     []Link `json:"_links,omitempty"`
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

	api     *ChimpAPI
	StoreID string `json:"-"`

	// Required
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Variants []Variant `json:"variants"`

	// Optional
	Handle      string `json:"handle,omitempty"`
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Vendor      string `json:"vendor,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	PublishedAt string `json:"published_at_foreign,omitempty"`

	// Response only
	Links []Link `json:"_links,omitempty"`
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
	if store.HasError() {
		return nil, fmt.Errorf("The store has an error, can't process request")
	}

	res := new(Product)
	res.api = store.api
	res.StoreID = store.ID

	endpoint := fmt.Sprintf(cart_path, store.ID, id)
	err := store.api.Request("GET", endpoint, params, nil, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (store Store) CreateProduct(req *Product) (*Product, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(products_path, store.ID)
	res := new(Product)
	res.api = store.api
	res.StoreID = store.ID

	return res, store.api.Request("POST", endpoint, nil, req, res)
}

func (store Store) UpdateProduct(req *Product) (*Product, error) {
	if err := store.CanMakeRequest(); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf(product_path, store.ID, req.ID)
	res := new(Product)
	res.api = store.api
	res.StoreID = store.ID

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

	api       *ChimpAPI
	StoreID   string `json:"-"`
	ProductID string `json:"-"`

	// Required
	ID    string `json:"id"`
	Title string `json:"title"`

	// Optional
	Url               string  `json:"url,omitempty"`
	SKU               string  `json:"sku,omitempty"`
	Price             float64 `json:"price,omitempty"`
	InventoryQuantity int     `json:"inventory_quantity,omitempty"`
	ImageUrl          string  `json:"image_url,omitempty"`
	Backorders        string  `json:"backorders,omitempty"`
	Visibility        string  `json:"visibility,omitempty"`
}

type VariantList struct {
	APIError

	StoreID    string    `json:"store_id"`
	Variants   []Variant `json:"variants"`
	TotalItems int       `json:"total_items"`
	Links      []Link    `json:"_links,omitempty"`
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
