package cart

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joshuahenriques/go-ecom/types"
)

var mockProducts = []types.Product{
	{ID: uuid.MustParse("f7342950-c643-461c-87de-03ca3abf354c"), Name: "product 1", Price: 10, Quantity: 100},
	{ID: uuid.MustParse("19b35093-c93b-40f4-8970-8bf27555424c"), Name: "product 2", Price: 20, Quantity: 200},
	{ID: uuid.MustParse("577544aa-fb27-43ac-9f20-44565d93a371"), Name: "product 3", Price: 30, Quantity: 300},
	{ID: uuid.MustParse("9b731eab-c217-4f71-aa50-6fd9356aaf76"), Name: "empty stock", Price: 30, Quantity: 0},
	{ID: uuid.MustParse("0def1069-ffa5-42b0-a538-fbe904c67600"), Name: "almost stock", Price: 30, Quantity: 1},
}

func TestCartServiceHandler(t *testing.T) {
	productStore := &mockProductStore{}
	orderStore := &mockOrderStore{}
	handler := NewHandler(orderStore, productStore, nil)

	t.Run("should fail to checkout if the cart items do not exist", func(t *testing.T) {
		payload := types.CartCheckoutPayload{
			Items: []types.CartItem{
				{ProductID: uuid.MustParse("1ce3516e-050a-495b-9cff-8c45bfc9c2f7"), Quantity: 100},
			},
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/cart/checkout", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/cart/checkout", handler.handleCheckout).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail to checkout if the cart has negative quantities", func(t *testing.T) {
		payload := types.CartCheckoutPayload{
			Items: []types.CartItem{
				{ProductID: uuid.MustParse("f7342950-c643-461c-87de-03ca3abf354c"), Quantity: 0}, // invalid quantity
			},
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/cart/checkout", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/cart/checkout", handler.handleCheckout).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail to checkout if there is no stock for an item", func(t *testing.T) {
		payload := types.CartCheckoutPayload{
			Items: []types.CartItem{
				{ProductID: uuid.MustParse("9b731eab-c217-4f71-aa50-6fd9356aaf76"), Quantity: 2},
			},
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/cart/checkout", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/cart/checkout", handler.handleCheckout).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail to checkout if there is not enough stock", func(t *testing.T) {
		payload := types.CartCheckoutPayload{
			Items: []types.CartItem{
				{ProductID: uuid.MustParse("0def1069-ffa5-42b0-a538-fbe904c67600"), Quantity: 2},
			},
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/cart/checkout", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/cart/checkout", handler.handleCheckout).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should checkout and calculate the price correctly", func(t *testing.T) {
		payload := types.CartCheckoutPayload{
			Items: []types.CartItem{
				{ProductID: uuid.MustParse("f7342950-c643-461c-87de-03ca3abf354c"), Quantity: 10},
				{ProductID: uuid.MustParse("19b35093-c93b-40f4-8970-8bf27555424c"), Quantity: 20},
				{ProductID: uuid.MustParse("0def1069-ffa5-42b0-a538-fbe904c67600"), Quantity: 1},
			},
		}

		marshalled, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/cart/checkout", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/cart/checkout", handler.handleCheckout).Methods(http.MethodPost)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}

		var response map[string]interface{}
		if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}

		if response["totalPrice"] != 530.0 {
			t.Errorf("expected total price to be 530, got %f", response["totalPrice"])
		}
	})
}

type mockProductStore struct{}

func (m *mockProductStore) GetProductByID(productID uuid.UUID) (*types.Product, error) {
	return &types.Product{}, nil
}

func (m *mockProductStore) GetProducts() (*[]types.Product, error) {
	return &[]types.Product{}, nil
}

func (m *mockProductStore) CreateProduct(product types.CreateProductPayload) error {
	return nil
}

func (m *mockProductStore) GetProductsByID(ids []string) (*[]types.Product, error) {
	return &mockProducts, nil
}

func (m *mockProductStore) UpdateProduct(product types.Product) error {
	return nil
}

type mockOrderStore struct{}

func (m *mockOrderStore) CreateOrder(order types.Order) (uuid.UUID, error) {
	return uuid.Nil, nil
}

func (m *mockOrderStore) CreateOrderItem(orderItem types.OrderItem) error {
	return nil
}
