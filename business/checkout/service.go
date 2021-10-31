package checkout

import (
	"AltaStore/business"
	"AltaStore/business/checkoutpayment"
	"AltaStore/business/shopping"
	"AltaStore/util/validator"
	"fmt"

	snap "github.com/midtrans/midtrans-go/snap"
)

type service struct {
	checkoutpaymentService checkoutpayment.Service
	shoppingService        shopping.Service
	repository             Repository
	repoShoppingDetail     RepoShoppingDetail
}

func NewService(
	checkoutpaymentService checkoutpayment.Service,
	shoppingService shopping.Service,
	repository Repository,
	repoShoppingDetail RepoShoppingDetail,
) Service {
	return &service{
		checkoutpaymentService,
		shoppingService,
		repository,
		repoShoppingDetail,
	}
}

func (s *service) NewCheckoutShoppingCart(userid string, checkout *Checkout) (*snap.Response, error) {
	err := validator.GetValidator().Struct(checkout)
	if err != nil {
		return nil, business.ErrInvalidSpec
	}

	var newCheckout = checkout.toCheckout(userid)

	status, err := s.repository.GetCheckoutByShoppingCartId(checkout.ShoppingCartId)
	if err != nil {
		return nil, err
	}

	if status {
		return nil, business.ErrDataExists
	}

	dets, err := s.repoShoppingDetail.GetShopCartDetailById(newCheckout.ShoppingCartId)
	if err != nil {
		return nil, err
	}

	err = s.repository.NewCheckoutShoppingCart(newCheckout)
	if err != nil {
		return nil, err
	}

	err = s.shoppingService.UpdateShopCartStatusById(checkout.ShoppingCartId, true)
	if err != nil {
		return nil, err
	}

	var sum int64 = 0
	for _, val := range *dets {
		sum += int64(val.Price)
	}

	var payment = checkoutpayment.InserPaymentSpec{
		OrderId:           checkout.ID,
		StatusCode:        "200",
		TransactionStatus: "pending",
	}

	_, err = s.checkoutpaymentService.InsertPayment(&payment, userid)
	if err != nil {
		return nil, err
	}
	return s.checkoutpaymentService.GenerateSnapPayment(
		newCheckout.CreatedBy,
		newCheckout.ID,
		sum)
}

func (s *service) GetAllCheckout() (*[]Checkout, error) {
	return s.repository.GetAllCheckout()
}

func (s *service) GetCheckoutById(id string) (*CheckItemDetails, error) {
	dtCheckout, err := s.repository.GetCheckoutById(id)

	fmt.Printf("%v", dtCheckout)
	if err != nil {
		return nil, err
	}

	items, err := s.repoShoppingDetail.GetShopCartDetailById(dtCheckout.ShoppingCartId)
	if err != nil {
		return nil, err
	}
	details := toDetailItemInCart(items)

	return getCheckItemsDetails(dtCheckout, details), nil
}
