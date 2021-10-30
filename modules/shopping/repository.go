package shopping

import (
	"AltaStore/business"
	"AltaStore/business/shopping"

	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type ShoppingCart struct {
	ID         string    `gorm:"id;type:uuid;primaryKey"`
	IsCheckOut bool      `gorm:"is_check_out;type:boolean;default:false"`
	CreatedBy  string    `gorm:"created_by;type:uuid"`
	CreatedAt  time.Time `gorm:"created_at;type:timestamp"`
	UpdatedAt  time.Time `gorm:"updated_at;type:timestamp"`
	DeletedAt  time.Time `gorm:"deleted_at;type:timestamp"`
}

func (s *ShoppingCart) toShoppCart() *shopping.ShoppCart {
	return &shopping.ShoppCart{
		ID:         s.ID,
		IsCheckOut: s.IsCheckOut,
		CreatedBy:  s.CreatedBy,
		CreatedAt:  s.CreatedAt,
		UpdatedAt:  s.UpdatedAt,
	}
}

func (s *ShoppingCart) newShoppingCart(id string, userid string, createdAt time.Time) {
	s.ID = id
	s.IsCheckOut = false
	s.CreatedBy = userid
	s.CreatedAt = createdAt
	s.UpdatedAt = createdAt
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetShoppingCartByUserId(userid string) (*shopping.ShoppCart, error) {
	var shopCart ShoppingCart

	err := r.DB.First(&shopCart, "is_check_out = false and created_by = ?", userid).Error
	if err != nil {
		return nil, err
	}

	return shopCart.toShoppCart(), nil
}

func (r *Repository) GetShoppingCartById(id string) (*shopping.ShoppCart, error) {
	var shopCart ShoppingCart

	err := r.DB.First(&shopCart, "is_check_out = false and id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return shopCart.toShoppCart(), nil
}

func (r *Repository) NewShoppingCart(id string, userid string, createdAt time.Time) (*shopping.ShoppCart, error) {
	var shopCart ShoppingCart

	err := r.DB.First(&shopCart, "is_check_out = false and created_by = ?", userid).Error

	// Pengecekan jika masih terdapat keranjang aktif maka dikembalikan bad request
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, business.ErrDataExists
	}

	shopCart.newShoppingCart(id, userid, createdAt)

	if err := r.DB.Create(&shopCart).Error; err != nil {
		return nil, err
	}

	return shopCart.toShoppCart(), nil
}
