package checkoutpayment

import (
	"time"
)

type PaymentStatus string

const (
	Pending     PaymentStatus = "pending"
	Capture     PaymentStatus = "captured"
	Settlement  PaymentStatus = "settlement"
	DenyPayment PaymentStatus = "deny"
	Cancel      PaymentStatus = "cancel"
	Expire      PaymentStatus = "expire"
	Failure     PaymentStatus = "failure"
)

type FraudStatus string

const (
	None      FraudStatus = ""
	Accept    FraudStatus = "accept"
	Deny      FraudStatus = "deny"
	Challenge FraudStatus = "challenge"
)

type CheckoutPayment struct {
	CheckOutID        string
	MerchantId        string
	StatusCode        string
	TransactionStatus PaymentStatus
	FraudStatus       FraudStatus
	CreatedAt         time.Time
	CreatedBy         string
	UpdatedAt         time.Time
	UpdatedBy         string
	DeletedAt         time.Time
	DeletedBy         string
}

func InsertPayment(
	checkoutId string,
	statusCode string,
	merchantId string,
	transactionstatus string,
	fraudstatus string,
	creator string,
	createAt time.Time,
) CheckoutPayment {
	return CheckoutPayment{
		CheckOutID:        checkoutId,
		MerchantId:        merchantId,
		StatusCode:        statusCode,
		TransactionStatus: PaymentStatus(transactionstatus),
		FraudStatus:       FraudStatus(fraudstatus),
		CreatedAt:         createAt,
		CreatedBy:         creator,
	}
}

// func (oldData *CheckoutPayment) ModifyPayment(
// 	merchantId string,
// 	transactionstatus string,
// 	fraudstatus string,
// 	modifier string,
// 	updatedAt time.Time,
// ) CheckoutPayment {
// 	return CheckoutPayment{
// 		CheckOutID:        oldData.CheckOutID,
// 		MerchantId:        merchantId,
// 		TransactionStatus: PaymentStatus(transactionstatus),
// 		FraudStatus:       FraudStatus(fraudstatus),
// 		CreatedAt:         oldData.CreatedAt,
// 		CreatedBy:         oldData.CreatedBy,
// 		UpdatedAt:         updatedAt,
// 		UpdatedBy:         modifier,
// 	}
// }
