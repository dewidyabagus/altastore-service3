package checkout

import (
	"AltaStore/business"
	"AltaStore/business/checkoutpayment"
	"AltaStore/util/validator"

	snap "github.com/midtrans/midtrans-go/snap"
)

type service struct {
	checkoutpaymentService checkoutpayment.Service
	repository             Repository
	repoShoppingDetail     RepoShoppingDetail
}

func NewService(
	checkoutpaymentService checkoutpayment.Service,
	repository Repository,
	repoShoppingDetail RepoShoppingDetail,

) Service {
	return &service{
		checkoutpaymentService,
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

	status, err := s.repository.GetCheckoutByShoppingCartId(checkout.ShoppingCardId)
	if err != nil {
		return nil, err
	}

	if status {
		return nil, business.ErrDataExists
	}

	dets, err := s.repoShoppingDetail.GetShopCartDetailById(newCheckout.ShoppingCardId)
	if err != nil {
		return nil, err
	}

	err = s.repository.NewCheckoutShoppingCart(newCheckout)
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
	data, err := s.repository.GetAllCheckout()
	if err != nil {
		return nil, err
	}
	for _, checkout := range *data {
		payment, err := s.checkoutpaymentService.GetPaymentByCheckoutId(checkout.ID)
		if err != nil {
			return nil, err
		}
		checkout.PaymentStatus = string(payment.TransactionStatus)
	}

	return data, nil
}

func (s *service) GetCheckoutById(id string) (*CheckItemDetails, error) {
	dtCheckout, err := s.repository.GetCheckoutById(id)

	if err != nil {
		return nil, err
	}

	items, err := s.repoShoppingDetail.GetShopCartDetailById(dtCheckout.ShoppingCardId)
	if err != nil {
		return nil, err
	}
	details := toDetailItemInCart(items)

	payment, err := s.checkoutpaymentService.GetPaymentByCheckoutId(dtCheckout.ID)
	if err != nil {
		return nil, err
	}

	return getCheckItemsDetails(dtCheckout, details, string(payment.TransactionStatus)), nil
}
