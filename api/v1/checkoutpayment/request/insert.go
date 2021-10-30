package request

import "AltaStore/business/checkoutpayment"

type InserPaymentRequest struct {
	OrderId           string `json:"order_id"`
	MerchantId        string `json:"merchant_id"`
	StatusCode        string `json:"status_code"`
	TransactionStatus string `json:"transaction_status"`
	FraudStatus       string `json:"fraud_status"`
}

func (u *InserPaymentRequest) ToPaymentSpec() *checkoutpayment.InserPaymentSpec {
	var spec checkoutpayment.InserPaymentSpec

	spec.OrderId = u.OrderId
	spec.MerchantId = u.MerchantId
	spec.StatusCode = u.StatusCode
	spec.TransactionStatus = u.TransactionStatus
	spec.FraudStatus = u.FraudStatus

	return &spec
}
