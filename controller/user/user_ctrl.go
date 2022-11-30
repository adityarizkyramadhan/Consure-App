package user

import (
	"Consure-App/sdk/response"
	userUc "Consure-App/usecase/user"
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
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ctrl *UserController) SignIn(ctx *gin.Context) {
	input := new(signInInput)
	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	token, err := ctrl.UserUc.SignIn(input.Username, input.Password)
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
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ResponseWhenFail(http.StatusBadRequest, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.ResponseWhenSuccess(http.StatusOK, "success", gin.H{
		"token": token,
	}))
}
