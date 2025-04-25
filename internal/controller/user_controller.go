package controller

import (
	"errors"
	"net/http"

	"github.com/dev-rever/cryptoo-pricing/internal/middleware"
	model "github.com/dev-rever/cryptoo-pricing/model/dto"
	"github.com/dev-rever/cryptoo-pricing/repository"
	"github.com/gin-gonic/gin"

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
		LogError(err)
		ctx.JSON(http.StatusBadRequest, ResponseError(InternalErrorCode, err.Error()))
		return
	}

	// check if the user exists
	exists, err := uc.userRepo.CheckUserExists(ctx, req.Account, req.Email)
	if err != nil {
		LogError(err)
		ctx.JSON(http.StatusInternalServerError, ResponseError(DBErrorCode, "database error"))
		return
	}
	if exists {
		err := errors.New("account or email already exists")
		LogError(err)
		ctx.JSON(http.StatusConflict, ResponseError(DBConflictErrorCode, err.Error()))
		return
	}

	// hash password
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		LogError(err)
		ctx.JSON(http.StatusInternalServerError, ResponseError(InternalErrorCode, err.Error()))
		return
	}

	// store user to db
	id, err := uc.userRepo.InsertUser(ctx, req.Account, string(hashedPassword), req.Email)
	if err != nil {
		LogError(err)
		ctx.JSON(http.StatusInternalServerError, ResponseError(DBErrorCode, err.Error()))
		return
	}

	accToken, err := middleware.GenerateJWT(id)
	if err != nil {
		LogError(err)
		ctx.JSON(http.StatusInternalServerError, ResponseError(InternalErrorCode, err.Error()))
		return
	}

	payload := model.RegisterResponse{
		Account: req.Account,
		Email:   req.Email,
		Token:   accToken,
	}

	msg := "user registered successfully"
	LogSuc(msg)
	ctx.JSON(http.StatusCreated, ResponseOK(msg, &payload))
}

func (uc *UserController) Login(ctx *gin.Context) {
	var req model.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		LogError(err)
		ctx.JSON(http.StatusBadRequest, ResponseError(InternalErrorCode, err.Error()))
		return
	}

	dbPwd, err := uc.userRepo.QueryUserPwdByAccount(ctx, req.Account)
	if err != nil {
		LogError(err)
		ctx.JSON(http.StatusInternalServerError, ResponseError(DBErrorCode, err.Error()))
		return
	}

	id, err := uc.userRepo.QueryUserIDByAccount(ctx, req.Account)
	if err != nil {
		LogError(err)
		ctx.JSON(http.StatusInternalServerError, ResponseError(DBErrorCode, err.Error()))
		return
	}

	accToken, err := middleware.GenerateJWT(id)
	if err != nil {
		LogError(err)
		ctx.JSON(http.StatusInternalServerError, ResponseError(AuthorizedErrorCode, err.Error()))
		return
	}

	if err := ComparePassword(dbPwd, req.Password); err != nil {
		LogError(err)
		ctx.JSON(http.StatusUnauthorized, ResponseError(AuthorizedErrorCode, err.Error()))
		return
	}

	payload := model.LoginResponse{
		Token: accToken,
	}
	msg := "user login successfully"
	LogSuc(msg)
	ctx.JSON(http.StatusOK, ResponseOK(msg, &payload))
}

func (uc *UserController) Profile(ctx *gin.Context) {
	if uidRaw, exist := ctx.Get("uid"); !exist {
		err := errors.New("unauthorized")
		LogError(err)
		ctx.JSON(http.StatusUnauthorized, ResponseError(AuthorizedErrorCode, err.Error()))
		return
	} else {
		uid := uidRaw.(uint)
		payload, err := uc.userRepo.QueryUserByID(ctx, uid)
		if err != nil {
			LogError(err)
			ctx.JSON(http.StatusInternalServerError, ResponseError(DBErrorCode, err.Error()))
			return
		}

		msg := "get user profile successfully"
		LogSuc(msg)
		ctx.JSON(http.StatusOK, ResponseOK(msg, &payload))
	}
}
