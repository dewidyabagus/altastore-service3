package shopping

import (
	"AltaStore/business"
	"AltaStore/util/validator"
	"time"

	"github.com/google/uuid"
)

type DetailItemInShopCart struct {
	ProductId string `validate:"required"`
	Price     int    `validate:"required"`
	Qty       int    `validate:"required"`
}

type service struct {
	repository     Repository
	repoCartDetail RepositoryCartDetail
}

func NewService(repository Repository, repoCartDetail RepositoryCartDetail) Service {
	return &service{repository, repoCartDetail}
}

func (s *service) GetShoppingCartByUserId(userid string) (*ShoppCart, error) {
	return s.repository.GetShoppingCartByUserId(userid)
}

func (s *service) NewShoppingCart(userid string) (*ShoppCart, error) {
	return s.repository.NewShoppingCart(uuid.NewString(), userid, time.Now())
}

func (s *service) GetShopCartDetailById(cartId string, userid string) (*ShopCartDetail, error) {
	shopCartUser, err := s.repository.GetShoppingCartByUserId(userid)
	if err != nil {
		return nil, err
	}

	shopCart, err := s.repository.GetShoppingCartById(cartId)
	if err != nil {
		return nil, err
	}

	if shopCartUser.ID != shopCart.ID {
		return nil, business.ErrNotFound
	}

	items, err := s.repoCartDetail.GetShopCartDetailById(cartId)
	if err != nil {
		return nil, err
	}

	cnvItems := toDetailItemInCart(items)

	return getShopCartDetailFormat(shopCart, cnvItems), nil
}

func (s *service) NewItemInShopCart(cartId string, item *DetailItemInShopCart, userid string) error {
	err := validator.GetValidator().Struct(item)
	if err != nil {
		return business.ErrInvalidSpec
	}
	shopCartUser, err := s.repository.GetShoppingCartByUserId(userid)
	if err != nil {
		return err
	}

	if shopCartUser.ID != cartId {
		return business.ErrNotHavePermission
	}

	return s.repoCartDetail.NewItemInShopCart(cartId, insertItemFormat(item))

}

func (s *service) ModifyItemInShopCart(cartId string, item *DetailItemInShopCart, userid string) error {
	err := validator.GetValidator().Struct(item)
	if err != nil {
		return business.ErrInvalidSpec
	}

	shopCartUser, err := s.repository.GetShoppingCartByUserId(userid)
	if err != nil {
		return err
	}

	if shopCartUser.ID != cartId {
		return business.ErrNotFound
	}

	return s.repoCartDetail.ModifyItemInShopCart(cartId, updateItemFormat(item))
}

func (s *service) DeleteItemInShopCart(cartId string, productid string, userid string) error {

	shopCartUser, err := s.repository.GetShoppingCartByUserId(userid)
	if err != nil {
		return err
	}

	if shopCartUser.ID != cartId {
		return business.ErrNotFound
	}

	return s.repoCartDetail.DeleteItemInShopCart(cartId, productid)
}
