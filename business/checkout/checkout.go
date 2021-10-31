package checkout

import (
	"AltaStore/modules/shoppingdetail"
	"time"

	"github.com/google/uuid"
)

type Checkout struct {
	ID             string
	ShoppingCardId string `validate:"required"`
	Description    string `validate:"required"`
	PaymentStatus  string
	CreatedBy      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (c *Checkout) toCheckout(userid string) *Checkout {

	c.ID = uuid.NewString()
	c.CreatedBy = userid
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	return c
}

type ItemInCart struct {
	ID          string
	ProductId   string
	ProductName string
	Price       int
	Qty         int
	UpdatedAt   time.Time
}

type CheckItemDetails struct {
	ID             string
	ShoppingCardId string
	Description    string
	PaymentStatus  string
	CreatedBy      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Details        []ItemInCart
}

func toItem(item shoppingdetail.ShopCartDetailItemWithProductName) ItemInCart {
	return ItemInCart{
		ID:          item.ID,
		ProductId:   item.ProductId,
		ProductName: item.ProductName,
		Price:       item.Price,
		Qty:         item.Qty,
		UpdatedAt:   item.UpdatedAt,
	}
}

func toDetailItemInCart(items *[]shoppingdetail.ShopCartDetailItemWithProductName) *[]ItemInCart {
	var details []ItemInCart

	for _, item := range *items {
		details = append(details, toItem(item))
	}

	if details == nil {
		details = []ItemInCart{}
	}

	return &details
}

func getCheckItemsDetails(cinfo *Checkout, items *[]ItemInCart) *CheckItemDetails {
	return &CheckItemDetails{
		ID:             cinfo.ID,
		ShoppingCardId: cinfo.ShoppingCardId,
		Description:    cinfo.Description,
		PaymentStatus:  cinfo.PaymentStatus,
		CreatedBy:      cinfo.CreatedBy,
		CreatedAt:      cinfo.CreatedAt,
		UpdatedAt:      cinfo.UpdatedAt,
		Details:        *items,
	}
}
