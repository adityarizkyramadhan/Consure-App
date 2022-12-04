package transaction

import (
	"Consure-App/middleware"
	transactionUc "Consure-App/usecase/transaction"

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
}

func (ctrl *TransactionController) Create(ctx *gin.Context) {

}
