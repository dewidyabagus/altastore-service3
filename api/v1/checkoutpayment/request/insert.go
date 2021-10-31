package request

import "AltaStore/business/checkoutpayment"

type InserPaymentRequestMidtrans struct {
	OrderId           string `json:"order_id"`
	MerchantId        string `json:"merchant_id"`
	StatusCode        string `json:"status_code"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
}

func (u *InserPaymentRequestMidtrans) ToPaymentSpec() *checkoutpayment.InserPaymentSpec {
	var spec checkoutpayment.InserPaymentSpec

	spec.OrderId = u.OrderId
	spec.MerchantId = u.MerchantId
	spec.StatusCode = u.StatusCode
	spec.TransactionStatus = u.TransactionStatus
	spec.FraudStatus = u.FraudStatus

	return &spec
}

type InserPaymentRequestAdmin struct {
	TransactionStatus string `json:"transactionStatus"`
}

func (u *InserPaymentRequestAdmin) ToPaymentSpec(id string) *checkoutpayment.InserPaymentSpec {
	var spec checkoutpayment.InserPaymentSpec

	spec.OrderId = id
	spec.MerchantId = ""
	spec.StatusCode = "200"
	spec.TransactionStatus = u.TransactionStatus
	spec.FraudStatus = ""

	return &spec
}
