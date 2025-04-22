package controller

import (
	"net/http"

	"github.com/dev-rever/cryptoo-pricing/model"
	"github.com/dev-rever/cryptoo-pricing/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	. "github.com/dev-rever/cryptoo-pricing/utils"
)

type UserController struct {
	userRepo *repository.UserRepo
}

func ProvideUserCtrl(repo *repository.UserRepo) *UserController {
	return &UserController{userRepo: repo}
}

func (uc *UserController) Root(ctx *gin.Context) {
	msg := "this is root page"
	ctx.JSON(http.StatusOK, msg)
}

func (uc *UserController) Register(ctx *gin.Context) {
	var req model.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		l := Log(ParserErrorCode, err.Error())
		ctx.JSON(http.StatusBadRequest, ResponseError(l.Code, l.Msg))
		return
	}

	// check if the user exists
	exists, err := uc.userRepo.CheckUserExists(ctx, req.Account, req.Email)
	if err != nil {
		l := Log(DBErrorCode, "database error")
		ctx.JSON(http.StatusInternalServerError, ResponseError(l.Code, l.Msg))
		return
	}
	if exists {
		l := Log(DBErrorCode, "account or email already exists")
		ctx.JSON(http.StatusConflict, ResponseError(l.Code, l.Msg))
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		l := Log(OtherErrorCode, "could not hash password")
		ctx.JSON(http.StatusInternalServerError, ResponseError(l.Code, "something error"))
		return
	}

	// store user to db
	_, err = uc.userRepo.CreateUser(ctx, req.Account, string(hashedPassword), req.Email)
	if err != nil {
		l := Log(DBErrorCode, "could not create user")
		ctx.JSON(http.StatusInternalServerError, ResponseError(l.Code, l.Msg))
		return
	}

	l := Log(SuccessCode, "user registered successfully")
	ctx.JSON(http.StatusCreated, ResponseOK(l.Msg, nil))
}
