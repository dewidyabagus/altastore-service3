package response

import (
	"AltaStore/business/checkout"
	"time"
)

type CheckoutResponse struct {
	ID             string    `json:"id"`
	ShoppingCardId string    `json:"shoppingcartid"`
	Description    string    `json:"description"`
	CreatedBy      string    `json:"created_by"`
	CreatedAt      time.Time `json:"created_at"`
}

func toCheckoutResponse(checkout checkout.Checkout) CheckoutResponse {
	return CheckoutResponse{
		ID:             checkout.ID,
		ShoppingCardId: checkout.ShoppingCardId,
		Description:    checkout.Description,
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
