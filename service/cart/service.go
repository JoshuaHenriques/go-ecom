package cart

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/joshuahenriques/go-ecom/types"
)

func getCartItemsIDs(items []types.CartItem) ([]string, error) {
	productIDs := make([]string, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity value: %s", item.ProductID)
		}

		productIDs[i] = item.ProductID.String()
	}

	return productIDs, nil
}

func (h *Handler) createOrder(products []types.Product, items []types.CartItem, userID uuid.UUID) (uuid.UUID, float64, error) {
	productMap := make(map[uuid.UUID]types.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return uuid.Nil, 0, nil
	}

	totalPrice := calculateTotalPrice(items, productMap)

	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)
	}

	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "689 Hickory Drive, Blackwood, NJ 08012", // todo: extend address
	})
	if err != nil {
		return uuid.Nil, 0, err
	}

	for _, item := range items {
		h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     productMap[item.ProductID].Price,
		})
	}

	return orderID, totalPrice, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, products map[uuid.UUID]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []types.CartItem, products map[uuid.UUID]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
