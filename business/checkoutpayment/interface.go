package checkoutpayment

import (
	snap "github.com/midtrans/midtrans-go/snap"
)

type Service interface {
	// GenerateSnapPayment from midtrans
	GenerateSnapPayment(customerId string, checkoutId string, amount int64) (*snap.Response, error)

	InsertPayment(payment *InserPaymentSpec) (*InserPaymentSpec, error)

	//ProcessMidtransCallback(notification *coreapi.TransactionStatusResponse) error
}

type Repository interface {
	InsertPayment(payment *CheckoutPayment) (*CheckoutPayment, error)

	//UpdatePayment(payment *CheckoutPayment) error
}
