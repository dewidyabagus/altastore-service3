package shopping

import (
	"AltaStore/modules/shoppingdetail"
	"time"
)

type Service interface {
	// Mengambil shopping cart aktif untuk berbelanja
	GetShoppingCartByUserId(userid string) (*ShoppCart, error)

	// Membuat keranjang belanjaan baru, ketika keranjang belanjaan ada yang aktif akan dikembalikan error
	NewShoppingCart(userid string) (*ShoppCart, error)

	// Mengambil detail item pada shopping cart
	GetShopCartDetailById(id string, userid string) (*ShopCartDetail, error)

	// Menambahkan item produk pada shopping cart
	NewItemInShopCart(cartId string, item *DetailItemInShopCart, userid string) error

	// Merubah item produk pada shopping cart
	ModifyItemInShopCart(cartId string, item *DetailItemInShopCart, userid string) error

	// Menghapus item produk pada shopping cart
	DeleteItemInShopCart(cartId string, productid string, userid string) error
}

type Repository interface {
	// Mengambil shopping cart aktif untuk berbelanja berdasarkan userid
	GetShoppingCartByUserId(userid string) (*ShoppCart, error)

	// Mengambil shopping cart berdasarkan id
	GetShoppingCartById(id string) (*ShoppCart, error)

	// Membuat keranjang belanjaan baru, ketika keranjang belanjaan ada yang aktif akan dikembalikan error
	NewShoppingCart(id string, userid string, createdAt time.Time) (*ShoppCart, error)
}

type RepositoryCartDetail interface {
	GetShopCartDetailById(id string) (*[]shoppingdetail.ShopCartDetailItemWithProductName, error)

	// Menambahkan item produk pada shopping cart
	NewItemInShopCart(cartId string, item *shoppingdetail.InsertItemInCartSpec) error

	// Merubah item produk pada shopping cart
	ModifyItemInShopCart(cartId string, item *shoppingdetail.UpdateItemInCartSpec) error

	// Menghapus item produk pada shopping cart
	DeleteItemInShopCart(cartId string, productid string) error
}
