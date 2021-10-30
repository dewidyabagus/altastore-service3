package migration

import (
	"AltaStore/business/checkoutpayment"

	"AltaStore/modules/checkout"
	"AltaStore/modules/shopping"
	"AltaStore/modules/shoppingdetail"

	"gorm.io/gorm"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&shopping.ShoppingCart{},
		&shoppingdetail.ShoppingCartDetail{},
		&checkout.Checkout{},
		&checkoutpayment.CheckoutPayment{},
	)
}
