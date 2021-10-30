package checkout

import (
	"AltaStore/modules/shoppingdetail"

	snap "github.com/midtrans/midtrans-go/snap"
)

type Service interface {
	NewCheckoutShoppingCart(userid string, checkout *Checkout) (*snap.Response, error)
	GetAllCheckout() (*[]Checkout, error)
	GetCheckoutById(id string) (*CheckItemDetails, error)
}

type Repository interface {
	NewCheckoutShoppingCart(checkout *Checkout) error
	GetAllCheckout() (*[]Checkout, error)
	GetCheckoutById(id string) (*Checkout, error)
	GetCheckoutByShoppingCartId(cartid string) (bool, error)
}

type RepoShoppingDetail interface {
	GetShopCartDetailById(id string) (*[]shoppingdetail.ShopCartDetailItemWithProductName, error)
}
