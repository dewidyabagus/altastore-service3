package api

import (
	"AltaStore/api/middleware"

	"AltaStore/api/v1/checkout"
	"AltaStore/api/v1/checkoutpayment"
	"AltaStore/api/v1/shopping"

	"net/http"

	echo "github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo,
	shopping *shopping.Controller,
	checkout *checkout.Controller,
	paymentController *checkoutpayment.Controller,
) {
	if shopping == nil || checkout == nil || paymentController == nil {
		panic("Invalid parameter")
	}

	// Add logger
	e.Use(middleware.MiddlewareLogger)

	// Custome response
	e.HTTPErrorHandler = func(e error, c echo.Context) {
		type Response struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		var response Response
		response.Code = http.StatusInternalServerError // defaul 500
		response.Message = "Internal Server Error"

		if he, ok := e.(*echo.HTTPError); ok {
			response.Code = he.Code
			response.Message = http.StatusText(he.Code)
		}

		c.Logger().Error(e)

		_ = c.JSON(response.Code, response)
	}

	// ROUTING LIST
	// Shopping Cart
	shopCart := e.Group("v1/shoppingcarts")
	shopCart.Use(middleware.JWTMiddleware())

	shopCart.GET("/carts", shopping.GetShoppingCartByUserId)
	shopCart.POST("/carts", shopping.NewShoppingCart)
	shopCart.GET("/carts/:id", shopping.GetShopCartDetailById)
	shopCart.POST("/carts/:id", shopping.NewItemInShopCart)
	shopCart.PUT("/carts/:id", shopping.ModifyItemInShopCart)
	shopCart.DELETE("/carts/:id/products/:productid", shopping.DeleteItemInShopCart)

	// Checkout
	c_out := e.Group("v1/checkouts")
	c_out.Use(middleware.JWTMiddleware())
	c_out.POST("", checkout.NewCheckoutShoppingCart)
	c_out.GET("", checkout.GetAllCheckout)
	c_out.GET("/:id", checkout.GetCheckoutById)

	// Payment
	payment := e.Group("v1/payments")
	payment.Use(middleware.JWTMiddleware())
	payment.PUT("/:id", paymentController.InsertPaymentById)

	paymentCallback := e.Group("v1/payments/notif")
	paymentCallback.GET("", paymentController.InsertPaymentFromMidtrans)
}
