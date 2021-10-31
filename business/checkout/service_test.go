package checkout_test

import (
	"AltaStore/business"
	"AltaStore/business/checkout"
	checkoutMock "AltaStore/business/checkout/mocks"
	checkoutpaymentMock "AltaStore/business/checkoutpayment/mocks"
	"AltaStore/modules/shoppingdetail"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id             = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	userid         = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	shoppingcartid = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	description    = "description"
	productid      = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	productname    = "productname"
	qty            = 10
	price          = 100000000
)

var (
	checkoutPaymentService   checkoutpaymentMock.Service
	checkoutRepository       checkoutMock.Repository
	checkoutDetailRepository checkoutMock.RepoShoppingDetail
	checkoutService          checkout.Service
	checkoutData             checkout.Checkout
	checkoutDatas            []checkout.Checkout

	shoppCartDetail    shoppingdetail.ShoppingCartDetail
	detailWithProduct  shoppingdetail.ShopCartDetailItemWithProductName
	detailWithProducts []shoppingdetail.ShopCartDetailItemWithProductName
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	checkoutData = checkout.Checkout{
		ID:             id,
		ShoppingCardId: shoppingcartid,
		Description:    description,
	}

	checkoutDatas = append(checkoutDatas, checkoutData)

	shoppCartDetail = shoppingdetail.ShoppingCartDetail{
		ID:        id,
		ProductId: productid,
		Price:     price,
		Qty:       qty,
	}
	detailWithProduct = shoppingdetail.ShopCartDetailItemWithProductName{
		ShoppingCartDetail: shoppCartDetail,
		ProductName:        productname,
	}
	detailWithProducts = append(detailWithProducts, detailWithProduct)

	checkoutService = checkout.NewService(&checkoutPaymentService, &checkoutRepository, &checkoutDetailRepository)
}

func TestNewCheckoutShoppingCart(t *testing.T) {
	t.Run("Expect Get Detail Shopping Cart Failed", func(t *testing.T) {
		checkoutRepository.On("GetCheckoutByShoppingCartId", mock.AnythingOfType("string")).Return(true, business.ErrInternalServer).Once()

		_, err := checkoutService.NewCheckoutShoppingCart(userid, &checkoutData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Get Detail Shopping Cart Status true", func(t *testing.T) {
		checkoutRepository.On("GetCheckoutByShoppingCartId", mock.AnythingOfType("string")).Return(true, nil).Once()

		_, err := checkoutService.NewCheckoutShoppingCart(userid, &checkoutData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrDataExists)
	})
	t.Run("Expect Get shopping cart detail failed", func(t *testing.T) {
		checkoutRepository.On("GetCheckoutByShoppingCartId", mock.AnythingOfType("string")).Return(false, nil).Once()
		checkoutDetailRepository.On("GetShopCartDetailById", mock.AnythingOfType("string")).Return(nil, business.ErrInternalServer).Once()

		_, err := checkoutService.NewCheckoutShoppingCart(userid, &checkoutData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Insert New Checkout Failed", func(t *testing.T) {
		checkoutRepository.On("GetCheckoutByShoppingCartId", mock.AnythingOfType("string")).Return(false, nil).Once()
		checkoutDetailRepository.On("GetShopCartDetailById", mock.AnythingOfType("string")).Return(&detailWithProducts, nil).Once()
		checkoutRepository.On("NewCheckoutShoppingCart", mock.AnythingOfType("*checkout.Checkout")).Return(business.ErrInternalServer).Once()

		_, err := checkoutService.NewCheckoutShoppingCart(userid, &checkoutData)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)

	})
	t.Run("Expect Generate Snap Payment Failed", func(t *testing.T) {
		checkoutRepository.On("GetCheckoutByShoppingCartId", mock.AnythingOfType("string")).Return(false, nil).Once()
		checkoutDetailRepository.On("GetShopCartDetailById", mock.AnythingOfType("string")).Return(&detailWithProducts, nil).Once()
		checkoutRepository.On("NewCheckoutShoppingCart", mock.AnythingOfType("*checkout.Checkout")).Return(nil).Once()
		checkoutPaymentService.On("GenerateSnapPayment",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("int64"),
		).Return(nil, business.ErrInternalServer).Once()

		_, err := checkoutService.NewCheckoutShoppingCart(userid, &checkoutData)

		assert.NotNil(t, err)
	})
	t.Run("Expect Insert New Checkout Success", func(t *testing.T) {
		checkoutRepository.On("GetCheckoutByShoppingCartId", mock.AnythingOfType("string")).Return(false, nil).Once()
		checkoutDetailRepository.On("GetShopCartDetailById", mock.AnythingOfType("string")).Return(&detailWithProducts, nil).Once()
		checkoutRepository.On("NewCheckoutShoppingCart", mock.AnythingOfType("*checkout.Checkout")).Return(nil).Once()
		checkoutPaymentService.On("GenerateSnapPayment",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("int64"),
		).Return(nil, nil).Once()
		_, err := checkoutService.NewCheckoutShoppingCart(userid, &checkoutData)

		assert.Nil(t, err)
	})
}

func TestGetAllCheckout(t *testing.T) {
	t.Run("Expect found the data Checkout", func(t *testing.T) {
		checkoutRepository.On("GetAllCheckout", mock.AnythingOfType("string")).Return(&checkoutDatas, nil).Once()

		data, err := checkoutService.GetAllCheckout()

		assert.Nil(t, err)
		assert.NotNil(t, data)

		assert.Equal(t, id, (*data)[0].ID)
		assert.Equal(t, shoppingcartid, (*data)[0].ShoppingCardId)
		assert.Equal(t, description, (*data)[0].Description)
	})

	t.Run("Expect data nil", func(t *testing.T) {
		checkoutRepository.On("GetAllCheckout", mock.AnythingOfType("string")).Return(nil, nil).Once()
		data, err := checkoutService.GetAllCheckout()

		assert.Nil(t, err)
		assert.Nil(t, data)
	})
}

func TestGetCheckoutById(t *testing.T) {
	t.Run("Expect Checkout not found", func(t *testing.T) {
		checkoutRepository.On("GetCheckoutById", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		_, err := checkoutService.GetCheckoutById(id)
		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)

	})

	t.Run("Expect Shopping Cart Detail not found", func(t *testing.T) {
		checkoutRepository.On("GetCheckoutById", mock.AnythingOfType("string")).Return(&checkoutData, nil).Once()
		checkoutDetailRepository.On("GetShopCartDetailById", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		_, err := checkoutService.GetCheckoutById(id)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)
	})

	t.Run("Expect Found Checkout", func(t *testing.T) {
		checkoutRepository.On("GetCheckoutById", mock.AnythingOfType("string")).Return(&checkoutData, nil).Once()
		checkoutDetailRepository.On("GetShopCartDetailById", mock.AnythingOfType("string")).Return(&detailWithProducts, nil).Once()

		data, err := checkoutService.GetCheckoutById(id)

		assert.NotNil(t, data)
		assert.Nil(t, err)
		//assert.Equal(t, id, data.ID)
		assert.Equal(t, shoppingcartid, data.ShoppingCardId)
		assert.Equal(t, description, data.Description)
	})
}
