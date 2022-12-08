package user

import (
	"Consure-App/domain"
	"Consure-App/middleware"
	"Consure-App/sdk/auth"
	"Consure-App/sdk/response"
	userUc "Consure-App/usecase/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUc userUc.UserUsecase
}

func NewUserController(userUc userUc.UserUsecase, r *gin.RouterGroup) {
	userCtrl := &UserController{
		UserUc: userUc,
	}
	r.POST("signin", userCtrl.SignIn)
	r.POST("signup", userCtrl.SignUp)
	r.PUT("profile", middleware.ValidateJWToken(), userCtrl.UpdateFotoProfile)
	r.GET("profile", middleware.ValidateJWToken(), userCtrl.GetProfile)
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ctrl *UserController) SignIn(ctx *gin.Context) {
	input := new(signInInput)
	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	token, err := ctrl.UserUc.SignIn(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusOK, "success", gin.H{
		"token": token,
	}))
}

type signUpInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (ctrl *UserController) SignUp(ctx *gin.Context) {
	input := new(signUpInput)
	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	token, err := ctrl.UserUc.SignUp(input.Username, input.Password, input.Email)
	if err != nil || token == "" {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusCreated, "success", gin.H{
		"token": token,
	}))
}

func (ctrl *UserController) UpdateFotoProfile(ctx *gin.Context) {
	avatar, err := ctx.FormFile("avatar")
	id := auth.GetIDFromBearer(ctx)
	log.Println("data => ", id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	link, err := ctrl.UserUc.UpdateProfile(id, avatar)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusOK, "success", gin.H{
		"link": link,
	}))
}

func (ctrl *UserController) GetProfile(ctx *gin.Context) {
	id := auth.GetIDFromBearer(ctx)
	data := new(domain.User)
	if err := ctrl.UserUc.GetProfile(id, data); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ResponseWhenFail(http.StatusInternalServerError, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusOK, "success", data))
}
