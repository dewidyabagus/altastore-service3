package shopping_test

import (
	"AltaStore/business"
	"AltaStore/business/shopping"
	shoppingMock "AltaStore/business/shopping/mocks"
	"AltaStore/modules/shoppingdetail"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	userid          = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	cartid          = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	wrongcartid     = "f9c8c2bf-d525-420e-86e5-4caf03cd8028"
	id              = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	code            = "code"
	name            = "name"
	description     = "description"
	productid       = "f9c8c2bf-d525-420e-86e5-4caf03cd8027"
	productname     = "productname"
	ischeckout      = false
	qty         int = 10
	price       int = 100000000
)

var (
	shoppingRepository       shoppingMock.Repository
	shoppingDetailRepository shoppingMock.RepositoryCartDetail
	shoppingService          shopping.Service

	shoppCart            shopping.ShoppCart
	shoppCartDetail      shoppingdetail.ShoppingCartDetail
	detailItemInShopCart shopping.DetailItemInShopCart
	detailWithProduct    shoppingdetail.ShopCartDetailItemWithProductName
	detailWithProducts   []shoppingdetail.ShopCartDetailItemWithProductName
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	shoppCart = shopping.ShoppCart{
		ID:         id,
		IsCheckOut: ischeckout,
	}
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

	detailItemInShopCart = shopping.DetailItemInShopCart{
		ProductId: productid,
		Price:     price,
		Qty:       qty,
	}
	shoppingService = shopping.NewService(&shoppingRepository, &shoppingDetailRepository)
}

func TestNewShoppingCart(t *testing.T) {
	t.Run("Expect Insert Shopping Cart Success", func(t *testing.T) {
		shoppingRepository.On("NewShoppingCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("time.Time"),
		).Return(&shoppCart, nil).Once()

		cart, err := shoppingService.NewShoppingCart(userid)

		assert.Nil(t, err)
		assert.Equal(t, cart.ID, shoppCart.ID)
		assert.Equal(t, cart.IsCheckOut, shoppCart.IsCheckOut)
	})
	t.Run("Expect Insert Shopping Cart Fail", func(t *testing.T) {
		shoppingRepository.On("NewShoppingCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("time.Time"),
		).Return(nil, business.ErrInternalServer).Once()

		_, err := shoppingService.NewShoppingCart(userid)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestNewItemInShopCart(t *testing.T) {
	t.Run("Expect Get Detail Shopping Cart by user id", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(nil, business.ErrInternalServer).Once()

		err := shoppingService.NewItemInShopCart(cartid, &detailItemInShopCart, userid)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Insert New Item In Shopping Cart Fail", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(&shoppCart, nil).Once()
		shoppingDetailRepository.On("NewItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*shoppingdetail.InsertItemInCartSpec"),
		).Return(business.ErrInternalServer).Once()

		err := shoppingService.NewItemInShopCart(cartid, &detailItemInShopCart, userid)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Insert New Item In Shopping Cart Success", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(&shoppCart, nil).Once()
		shoppingDetailRepository.On("NewItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*shoppingdetail.InsertItemInCartSpec"),
		).Return(nil).Once()

		err := shoppingService.NewItemInShopCart(cartid, &detailItemInShopCart, userid)

		assert.Nil(t, err)
	})

}

func TestModifyItemInShopCart(t *testing.T) {
	t.Run("Expect Get Detail Shopping Cart by user id", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(nil, business.ErrInternalServer).Once()

		err := shoppingService.NewItemInShopCart(cartid, &detailItemInShopCart, userid)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Modify Item In Shopping Cart Fail", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(&shoppCart, nil).Once()
		shoppingDetailRepository.On("ModifyItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*shoppingdetail.UpdateItemInCartSpec"),
		).Return(business.ErrInternalServer).Once()

		err := shoppingService.ModifyItemInShopCart(cartid, &detailItemInShopCart, userid)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Modify Item In Shopping Cart Success", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(&shoppCart, nil).Once()
		shoppingDetailRepository.On("ModifyItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("*shoppingdetail.UpdateItemInCartSpec"),
		).Return(nil).Once()

		err := shoppingService.ModifyItemInShopCart(cartid, &detailItemInShopCart, userid)

		assert.Nil(t, err)
	})
}

func TestDeleteItemInShopCart(t *testing.T) {
	t.Run("Expect Get Detail Shopping Cart by user id", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(nil, business.ErrInternalServer).Once()

		err := shoppingService.NewItemInShopCart(cartid, &detailItemInShopCart, userid)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Delete Item In Shopping Cart Fail", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(&shoppCart, nil).Once()
		shoppingDetailRepository.On("DeleteItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(business.ErrInternalServer).Once()

		err := shoppingService.DeleteItemInShopCart(cartid, productid, userid)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Delete Item In Shopping Cart Success", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(&shoppCart, nil).Once()
		shoppingDetailRepository.On("DeleteItemInShopCart",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(nil).Once()

		err := shoppingService.DeleteItemInShopCart(cartid, productid, userid)

		assert.Nil(t, err)
	})
}

func TestGetShoppingCartByUserId(t *testing.T) {
	t.Run("Expect found the Shopping Cart", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId",
			mock.AnythingOfType("string"),
		).Return(&shoppCart, nil).Once()

		shopping, err := shoppingService.GetShoppingCartByUserId(userid)

		assert.Nil(t, err)
		assert.NotNil(t, shopping)

		assert.Equal(t, id, shopping.ID)
		assert.Equal(t, ischeckout, shopping.IsCheckOut)

	})

	t.Run("Expect Shopping Cart not found", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId",
			mock.AnythingOfType("string"),
		).Return(nil, business.ErrInternalServer).Once()

		shopping, err := shoppingService.GetShoppingCartByUserId(userid)

		assert.NotNil(t, err)
		assert.Nil(t, shopping)

		assert.Equal(t, err, business.ErrInternalServer)
	})
}

func TestGetShopCartDetailById(t *testing.T) {
	t.Run("Expect Get Detail Shopping Cart by user id", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(nil, business.ErrInternalServer).Once()

		err := shoppingService.NewItemInShopCart(cartid, &detailItemInShopCart, userid)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Shopping Cart Not Found", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(&shoppCart, nil).Once()
		shoppingRepository.On("GetShoppingCartById",
			mock.AnythingOfType("string"),
		).Return(nil, business.ErrNotFound).Once()

		_, err := shoppingService.GetShopCartDetailById(id, userid)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)

	})
	t.Run("Expect Detail Shopping Cart Not Found", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(&shoppCart, nil).Once()
		shoppingRepository.On("GetShoppingCartById",
			mock.AnythingOfType("string"),
		).Return(&shoppCart, nil).Once()
		shoppingDetailRepository.On("GetShopCartDetailById",
			mock.AnythingOfType("string"),
		).Return(nil, business.ErrNotFound).Once()

		_, err := shoppingService.GetShopCartDetailById(id, userid)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrNotFound)

	})
	t.Run("Expect found the Detail Shopping Cart", func(t *testing.T) {
		shoppingRepository.On("GetShoppingCartByUserId", mock.AnythingOfType("string")).Return(&shoppCart, nil).Once()
		shoppingRepository.On("GetShoppingCartById",
			mock.AnythingOfType("string"),
		).Return(&shoppCart, nil).Once()
		shoppingDetailRepository.On("GetShopCartDetailById",
			mock.AnythingOfType("string"),
		).Return(&detailWithProducts, nil).Once()

		shopping, err := shoppingService.GetShopCartDetailById(id, userid)

		assert.Nil(t, err)
		assert.NotNil(t, shopping)

		assert.Equal(t, id, shopping.ID)
	})
}

func TestUpdateShopCartStatusById(t *testing.T) {
	t.Run("Expect Update status shopping cart failed", func(t *testing.T) {
		shoppingRepository.On("UpdateShopCartStatusById",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("bool"),
		).Return(business.ErrInternalServer).Once()

		err := shoppingService.UpdateShopCartStatusById(id, ischeckout)

		assert.NotNil(t, err)
		assert.Equal(t, err, business.ErrInternalServer)
	})
	t.Run("Expect Update status shopping cart succes", func(t *testing.T) {
		shoppingRepository.On("UpdateShopCartStatusById",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("bool"),
		).Return(nil).Once()

		err := shoppingService.UpdateShopCartStatusById(id, ischeckout)

		assert.Nil(t, err)
	})
}
