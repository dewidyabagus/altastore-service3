package checkoutpayment

import (
	"AltaStore/api/common"
	"AltaStore/api/middleware"
	"AltaStore/api/v1/checkoutpayment/request"
	"AltaStore/business/checkoutpayment"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service checkoutpayment.Service
}

func NewController(service checkoutpayment.Service) *Controller {
	return &Controller{service}
}

// // MidtransTransactionCallbackHandler handles incoming notification about payment status from midtrans.
// func (c *Controller) Call(ctx echo.Context) error {
// 	userid := ctx.QueryParam("userid")

// 	if _, err := uuid.Parse(userid); err != nil {
// 		return ctx.JSON(common.BadRequestResponse())
// 	}
// 	snap, err := c.service.GenerateSnapPayment(userid, uuid.New().String(), 10000)
// 	if err != nil {
// 		return ctx.JSON(common.BadRequestResponse())
// 	}
// 	return ctx.JSON(
// 		common.SuccessResponseWithData(snap),
// 	)
// }

// MidtransTransactionCallbackHandler handles incoming notification about payment status from midtrans.
func (c *Controller) InsertPaymentFromMidtrans(ctx echo.Context) error {
	merchantId := ctx.QueryParam("merchant_id")
	orderId := ctx.QueryParam("order_id")
	statusCode := ctx.QueryParam("status_code")
	transactionStatus := ctx.QueryParam("transaction_status")
	fraudStatus := ctx.QueryParam("fraud_status")

	if _, err := uuid.Parse(orderId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}
	var newData request.InserPaymentRequestMidtrans
	newData.OrderId = orderId
	newData.FraudStatus = fraudStatus
	newData.TransactionStatus = transactionStatus
	newData.MerchantId = merchantId
	newData.StatusCode = statusCode
	saveData, err := c.service.InsertPayment(newData.ToPaymentSpec(), "Midtrans")
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}
	return ctx.JSON(
		common.SuccessResponseWithData(saveData),
	)
}
func (c *Controller) InsertPaymentById(ctx echo.Context) error {
	id := ctx.Param("id")

	payment := new(request.InserPaymentRequestAdmin)
	adminId, err := middleware.ExtractTokenUser(ctx)
	if err != nil {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	isAdmin, err := middleware.ExtractTokenRule(ctx)
	if err != nil || !isAdmin {
		return ctx.JSON(common.UnAuthorizedResponse())
	}
	if _, err := uuid.Parse(adminId); err != nil {
		return ctx.JSON(common.BadRequestResponse())
	}
	saveData, err := c.service.InsertPayment(payment.ToPaymentSpec(id), adminId)
	if err != nil {
		return ctx.JSON(common.NewBusinessErrorResponse(err))
	}
	return ctx.JSON(
		common.SuccessResponseWithData(saveData),
	)
}

// // MidtransTransactionCallbackHandler handles incoming notification about payment status from midtrans.
// func (c *Controller) MidtransTransactionCallbackHandler(ctx echo.Context) error {
// 	_ = func(w http.ResponseWriter, r *http.Request) error {
// 		decoder := json.NewDecoder(r.Body)
// 		var notification coreapi.TransactionStatusResponse
// 		err := decoder.Decode(&notification)
// 		if err != nil {
// 			return ctx.JSON(common.NotFoundResponse())
// 		}
// 		if err != nil {
// 			return ctx.JSON(common.NotFoundResponse())
// 		}
// 		err = c.service.ProcessMidtransCallback(&notification)
// 		if err != nil {
// 			return ctx.JSON(common.NotFoundResponse())
// 		}
// 		return ctx.JSON(
// 			common.SuccessResponseWithoutData(),
// 		)
// 	}
// 	return nil
// }
