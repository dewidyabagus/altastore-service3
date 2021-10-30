package checkoutpayment_test

import (
	"AltaStore/business"
	"AltaStore/business/checkoutpayment"
	checkoutpaymentMock "AltaStore/business/checkoutpayment/mocks"
	"AltaStore/business/user"
	userMock "AltaStore/business/user/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	userid     = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	email      = "email@test.com"
	firstname  = "firstname"
	lastname   = "lastname"
	password   = "password"
	statusCode = "statusCode"

	checkoutid              = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	merchantId              = "merchantId"
	transactionstatus       = "transactionstatus"
	fraudstatus             = "fraudstatus"
	grossprice        int64 = 100000000
	snappayment             = "snappayment"
)

var (
	userService               userMock.Service
	checkoutPaymentRepository checkoutpaymentMock.Repository
	checkoutPaymentService    checkoutpayment.Service

	userData          user.User
	checkoutPayment   checkoutpayment.CheckoutPayment
	insertPaymentSpec checkoutpayment.InserPaymentSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	checkoutPayment = checkoutpayment.CheckoutPayment{
		CheckOutID:        checkoutid,
		MerchantId:        merchantId,
		FraudStatus:       fraudstatus,
		TransactionStatus: transactionstatus,
		StatusCode:        statusCode,
	}

	insertPaymentSpec = checkoutpayment.InserPaymentSpec{
		OrderId:           checkoutid,
		MerchantId:        merchantId,
		FraudStatus:       fraudstatus,
		TransactionStatus: transactionstatus,
		StatusCode:        statusCode,
	}
	checkoutPaymentService = checkoutpayment.NewService(&userService, &checkoutPaymentRepository)
}

func TestGenerateSnapPayment(t *testing.T) {
	t.Run("Expect User Not Found", func(t *testing.T) {
		userService.On("FindUserByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotHavePermission).Once()

		user, err := userService.FindUserByID(userid)

		assert.Nil(t, user)
		assert.NotNil(t, err, business.ErrNotHavePermission)

	})
	// t.Run("Expect Generate Checkout Payment Success", func(t *testing.T) {
	// 	userService.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()

	// 	_, err := checkoutPaymentService.GenerateSnapPayment(userid, checkoutid, grossprice)

	// 	assert.Nil(t, err)
	// })
	// t.Run("Expect Generate Checkout Payment Fail", func(t *testing.T) {
	// 	userService.On("FindUserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
	// 	_, err := checkoutPaymentService.GenerateSnapPayment(userid, checkoutid, grossprice)

	// 	assert.NotNil(t, err)
	// })
}

func TestInserPayment(t *testing.T) {
	t.Run("Expect Insert Checkout Payment Success", func(t *testing.T) {
		checkoutPaymentRepository.On("InsertPayment", mock.AnythingOfType("*checkoutpayment.CheckoutPayment")).Return(&checkoutPayment, nil).Once()

		_, err := checkoutPaymentService.InsertPayment(&insertPaymentSpec)

		assert.Nil(t, err)
	})
	t.Run("Expect Insert Checkout Payment Fail", func(t *testing.T) {
		userService.On("FinduserByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
		checkoutPaymentRepository.On("InsertPayment", mock.AnythingOfType("*checkoutpayment.CheckoutPayment")).Return(nil, business.ErrInternalServer).Once()

		_, err := checkoutPaymentService.InsertPayment(&insertPaymentSpec)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}
