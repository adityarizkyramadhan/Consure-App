package review

import (
	"Consure-App/domain"
	"Consure-App/dto"
	"Consure-App/middleware"
	"Consure-App/sdk/auth"
	"Consure-App/sdk/response"
	"Consure-App/usecase/review"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewController struct {
	RevUc review.ReviewUsecase
}

func NewController(revUc review.ReviewUsecase, r *gin.RouterGroup) {
	ctrl := &ReviewController{
		RevUc: revUc,
	}
	r.GET("all", ctrl.GetAll)
	r.GET("expert/:id", ctrl.GetByIdExpert)
	r.GET("user", middleware.ValidateJWToken(), ctrl.GetByIdUSer)
	r.POST("expert/:id_expert", middleware.ValidateJWToken(), ctrl.CreateReview)
}

func (ctrl *ReviewController) GetAll(ctx *gin.Context) {
	data := []*domain.Review{}
	if err := ctrl.RevUc.FindAll(&data); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseWhenFail(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusOK, "success", data))
}

func (ctrl *ReviewController) GetByIdExpert(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	data := new(dto.DataExpertWithReview)
	if err := ctrl.RevUc.FindByIdExpert(id, data); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseWhenFail(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusOK, "success", data))
}

func (ctrl *ReviewController) GetByIdUSer(ctx *gin.Context) {
	id := auth.GetIDFromBearer(ctx)
	data := []*domain.Review{}
	if err := ctrl.RevUc.FindByIdUser(id, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseWhenFail(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusOK, "success", data))
}

func (ctrl *ReviewController) CreateReview(ctx *gin.Context) {
	id := auth.GetIDFromBearer(ctx)
	idExpert, err := strconv.Atoi(ctx.Param("id_expert"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	input := new(review.InputReview)
	if err := ctx.ShouldBind(input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	input.IdExpert = idExpert
	if err := ctrl.RevUc.Create(id, input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, response.ResponseWhenSuccess(http.StatusCreated, "success", nil))
}
