package transaction

import (
	"Consure-App/dto"
	"Consure-App/middleware"
	"Consure-App/sdk/auth"
	"Consure-App/sdk/response"
	transactionUc "Consure-App/usecase/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TrxUc transactionUc.TransactionUsecase
}

func NewTransactionController(trxUc transactionUc.TransactionUsecase, r *gin.RouterGroup) {
	trxCtrl := &TransactionController{
		TrxUc: trxUc,
	}
	r.POST("", middleware.ValidateJWToken(), trxCtrl.Create)
	r.GET("history", middleware.ValidateJWToken(), trxCtrl.History)
}

func (ctrl *TransactionController) Create(ctx *gin.Context) {
	id := auth.GetIDFromBearer(ctx)
	input := new(transactionUc.InputTransaction)
	if err := ctx.ShouldBind(input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	if err := ctrl.TrxUc.Create(input, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseWhenFail(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, response.ResponseWhenSuccess(http.StatusCreated, "Success", nil))
}

func (ctrl *TransactionController) History(ctx *gin.Context) {
	status := ctx.Query("status")
	id := auth.GetIDFromBearer(ctx)
	data := []*dto.History{}
	if err := ctrl.TrxUc.History(id, status, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseWhenFail(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusOK, "Success", data))

}
