package expert

import (
	"Consure-App/sdk/response"
	"Consure-App/usecase/expert"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpertController struct {
	ExUc expert.ExpertUsecase
}

func NewExpertController(exUc expert.ExpertUsecase, r *gin.RouterGroup) {
	exCtrl := &ExpertController{
		ExUc: exUc,
	}
	r.POST("", exCtrl.SignUp)
}

func (exCtrl *ExpertController) SignUp(ctx *gin.Context) {
	input := new(expert.InputExpert)
	if err := ctx.ShouldBind(input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	if err := exCtrl.ExUc.SignUp(input); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseWhenFail(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, response.ResponseWhenSuccess(http.StatusCreated, "Success", nil))
}
