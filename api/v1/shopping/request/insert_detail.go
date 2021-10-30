package request

import "AltaStore/business/shopping"

type DetailItemInShopCart struct {
	ProductId string `json:"productid"`
	Price     int    `json:"price"`
	Qty       int    `json:"qty"`
}

func (d *DetailItemInShopCart) ToDetailItemInShopCart() *shopping.DetailItemInShopCart {
	return &shopping.DetailItemInShopCart{
		ProductId: d.ProductId,
		Price:     d.Price,
		Qty:       d.Qty,
	}
}
