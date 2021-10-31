package shopping

import (
	"AltaStore/api/common"
	"AltaStore/api/middleware"

	"AltaStore/api/v1/shopping/request"
	"AltaStore/api/v1/shopping/response"

	"AltaStore/business/shopping"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	service shopping.Service
}

func NewController(service shopping.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) GetShoppingCartByUserId(ctx echo.Context) error {
	userId := ctx.Param("id")
	if _, err := uuid.Parse(userId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	shoppCart, err := c.service.GetShoppingCartByUserId(userId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	response := response.GetOneResponse(shoppCart)

	return ctx.JSON(common.SuccessResponseWithData(response))
}

func (c *Controller) NewShoppingCart(ctx echo.Context) error {
	var err error

	userId, err := middleware.ExtractTokenUser(ctx)
	if err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}

	if _, err := uuid.Parse(userId); err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}

	result, err := c.service.NewShoppingCart(userId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	response := response.GetOneResponse(result)

	return ctx.JSON(common.SuccessResponseWithData(response))
}

func (c *Controller) GetShopCartDetailById(ctx echo.Context) error {
	cartId := ctx.Param("id")

	if _, err := uuid.Parse(cartId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	itemDetail, err := c.service.GetShopCartDetailById(cartId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithData(response.ToResponseDetails(itemDetail)))
}

func (c *Controller) NewItemInShopCart(ctx echo.Context) error {
	var item = new(request.DetailItemInShopCart)

	cartId := ctx.Param("id")
	if _, err := uuid.Parse(cartId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := ctx.Bind(item); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err := c.service.NewItemInShopCart(cartId, item.ToDetailItemInShopCart())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) ModifyItemInShopCart(ctx echo.Context) error {
	var err error
	var item = new(request.DetailItemInShopCart)

	cartId := ctx.Param("id")
	if _, err := uuid.Parse(cartId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	if err := ctx.Bind(item); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err = c.service.ModifyItemInShopCart(cartId, item.ToDetailItemInShopCart())
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}

func (c *Controller) DeleteItemInShopCart(ctx echo.Context) error {
	var err1, err2 error

	cartId := ctx.Param("id")
	productId := ctx.Param("productid")

	_, err1 = uuid.Parse(cartId)
	_, err2 = uuid.Parse(productId)

	if err1 != nil || err2 != nil {
		return ctx.JSON(common.BadRequestResponse())
	}

	err := c.service.DeleteItemInShopCart(cartId, productId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}

	return ctx.JSON(common.SuccessResponseWithoutData())
}
