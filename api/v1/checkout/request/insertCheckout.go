package request

import (
	"AltaStore/business/checkout"
)

type NewCheckoutShoppingCart struct {
	ShoppingCartId string `json:"shoppingcartid"`
	Description    string `json:"description"`
}

func (n *NewCheckoutShoppingCart) ToBusinessCheckout() *checkout.Checkout {
	var checkout checkout.Checkout

	checkout.ShoppingCardId = n.ShoppingCartId
	checkout.Description = n.Description

	return &checkout
}
