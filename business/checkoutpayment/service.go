package checkoutpayment

import (
	"AltaStore/business"
	"AltaStore/business/user"
	"AltaStore/config"
	"time"

	midtrans "github.com/midtrans/midtrans-go"
	snap "github.com/midtrans/midtrans-go/snap"
)

type InserPaymentSpec struct {
	OrderId           string `validate:"required"`
	MerchantId        string
	StatusCode        string `validate:"required"`
	TransactionStatus string `validate:"required"`
	FraudStatus       string
}

type service struct {
	userService user.Service
	repository  Repository
}

func NewService(userService user.Service, repository Repository) Service {
	return &service{userService, repository}
}

func (s *service) GenerateSnapPayment(customerId string, checkoutId string, amount int64) (*snap.Response, error) {
	// 1. Initiate Snap client
	var sc snap.Client
	var key = config.GetConfig().MidTransServerKey
	sc.New(key, midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	user, err := s.userService.FindUserByID(customerId)
	if err != nil {
		var res snap.Response
		return &res, business.ErrNotFound
	}
	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  checkoutId,
			GrossAmt: amount,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.FirstName,
			LName: user.LastName,
			Email: user.Email,
			Phone: user.HandPhone,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, midtransErr := sc.CreateTransaction(req)
	if midtransErr != nil {
		return snapResp, business.ErrNotFound
	}
	return snapResp, nil
}

func (s *service) InsertPayment(p *InserPaymentSpec) (*InserPaymentSpec, error) {

	//TODO: Check inv apakah ada

	data := InsertPayment(
		p.OrderId,
		p.StatusCode,
		p.MerchantId,
		p.TransactionStatus,
		p.FraudStatus,
		"midtrans",
		time.Now())

	_, err := s.repository.InsertPayment(&data)
	if err != nil {
		return p, business.ErrInternalServer
	}
	return p, nil
}
