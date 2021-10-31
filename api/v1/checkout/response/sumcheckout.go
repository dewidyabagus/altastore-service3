package response

import (
	"AltaStore/business/checkout"
	"time"
)

type CheckoutResponse struct {
	ID             string    `json:"id"`
	ShoppingCardId string    `json:"shoppingcartid"`
	Description    string    `json:"description"`
	PaymentStatus  string    `json:"paymentstatus"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
}

func toCheckoutResponse(checkout checkout.Checkout) CheckoutResponse {
	return CheckoutResponse{
		ID:             checkout.ID,
		ShoppingCardId: checkout.ShoppingCartId,
		Description:    checkout.Description,
		PaymentStatus:  checkout.PaymentStatus,
		CreatedBy:      checkout.CreatedBy,
		CreatedAt:      checkout.CreatedAt,
	}
}

func AllCheckout(checkout *[]checkout.Checkout) *[]CheckoutResponse {
	var response []CheckoutResponse

	for _, val := range *checkout {
		response = append(response, toCheckoutResponse(val))
	}

	if response == nil {
		response = []CheckoutResponse{}
	}

	return &response
}

type DetailItem struct {
	ProductId   string    `json:"product_id"`
	ProductName string    `json:"product_name"`
	Price       int       `json:"price"`
	Qty         int       `json:"qty"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResponseDetailItems struct {
	ID             string       `json:"id"`
	ShoppingCartId string       `json:"shopping_cart_id"`
	Description    string       `json:"description"`
	CreatedBy      string       `json:"created_by"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	Details        []DetailItem `json:"details"`
}

func ToDetailItem(item *checkout.ItemInCart) *DetailItem {
	return &DetailItem{
		ProductId:   item.ProductId,
		ProductName: item.ProductName,
		Price:       item.Price,
		Qty:         item.Qty,
		UpdatedAt:   item.UpdatedAt,
	}
}

func ToResponseDetailItems(checkout *checkout.CheckItemDetails) *ResponseDetailItems {
	var items ResponseDetailItems

	items.ID = checkout.ID
	items.ShoppingCartId = checkout.ShoppingCardId
	items.Description = checkout.Description
	items.CreatedBy = checkout.CreatedBy
	items.CreatedAt = checkout.CreatedAt
	items.UpdatedAt = checkout.UpdatedAt

	for _, item := range checkout.Details {
		items.Details = append(items.Details, *ToDetailItem(&item))
	}

	if items.Details == nil {
		items.Details = []DetailItem{}
	}

	return &items
}
