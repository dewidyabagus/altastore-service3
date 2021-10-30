package checkout

import (
	"AltaStore/api/common"
	"AltaStore/api/middleware"

	"AltaStore/api/v1/checkout/request"
	"AltaStore/api/v1/checkout/response"

	"AltaStore/business/checkout"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	service checkout.Service
}

func NewController(service checkout.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) NewCheckoutShoppingCart(ctx echo.Context) error {
	var err error

	checkoutShopCart := new(request.NewCheckoutShoppingCart)

	if err = ctx.Bind(checkoutShopCart); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	userId, err := middleware.ExtractTokenUser(ctx)
	if err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	if _, err := uuid.Parse(userId); err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}

	snap, err := c.service.NewCheckoutShoppingCart(userId, checkoutShopCart.ToBusinessCheckout())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(snap))
}

func (c *Controller) GetAllCheckout(ctx echo.Context) error {
	listCheckout, err := c.service.GetAllCheckout()
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.AllCheckout(listCheckout)))
}

func (c *Controller) GetCheckoutById(ctx echo.Context) error {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	checkoutDetails, err := c.service.GetCheckoutById(id)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(checkoutDetails))
}
