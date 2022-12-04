package expert

import (
	"Consure-App/domain"
	"Consure-App/sdk/response"
	"Consure-App/usecase/expert"
	"fmt"
	"net/http"
	"strconv"

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
	r.GET("all", exCtrl.FindAll)
	r.GET("single/:id", exCtrl.FindById)
	r.GET("search", exCtrl.FindByTag)
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

func (exCtrl *ExpertController) FindAll(ctx *gin.Context) {
	data := []*domain.Expert{}
	if err := exCtrl.ExUc.FindAll(&data); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseWhenFail(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusOK, "Success", data))
}

func (exCtrl *ExpertController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := new(domain.Expert)
	if err := exCtrl.ExUc.FindById(id, data); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseWhenFail(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusOK, "Success", data))
}

func (exCtrl *ExpertController) FindByTag(ctx *gin.Context) {
	tag := ctx.Query("tag")
	fmt.Println(tag)
	data := []*domain.Expert{}
	if err := exCtrl.ExUc.FindByTag(tag, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseWhenFail(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusOK, "Success", data))
}
