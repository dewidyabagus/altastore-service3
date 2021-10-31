package response

import (
	"AltaStore/business/shopping"
	"time"
)

type ShoppData struct {
	ID         string    `json:"id"`
	IsCheckOut bool      `json:"ischeckout"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type DetailItem struct {
	ProductId   string    `json:"product_id"`
	ProductName string    `json:"product_name"`
	Price       int       `json:"price"`
	Qty         int       `json:"qty"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResponseDetailItems struct {
	ID        string       `json:"id"`
	CreatedBy string       `json:"created_by"`
	UpdatedAt time.Time    `json:"updated_at"`
	Details   []DetailItem `json:"details"`
}

func ToDetailItem(item *shopping.ItemInCart) *DetailItem {
	return &DetailItem{
		ProductId:   item.ProductId,
		ProductName: item.ProductName,
		Price:       item.Price,
		Qty:         item.Qty,
		UpdatedAt:   item.UpdatedAt,
	}
}

func ToResponseDetails(cart *shopping.ShopCartDetail) *ResponseDetailItems {
	var items ResponseDetailItems

	items.ID = cart.ID
	items.CreatedBy = cart.CreatedBy
	items.UpdatedAt = cart.UpdatedAt

	for _, item := range cart.Details {
		items.Details = append(items.Details, *ToDetailItem(&item))
	}

	if items.Details == nil {
		items.Details = []DetailItem{}
	}

	return &items
}
