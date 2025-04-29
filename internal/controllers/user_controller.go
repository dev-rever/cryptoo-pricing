package controllers

import (
	"errors"
	"net/http"

	"github.com/dev-rever/cryptoo-pricing/internal/middleware/jwt"
	"github.com/dev-rever/cryptoo-pricing/repositories"
	"github.com/gin-gonic/gin"

	model "github.com/dev-rever/cryptoo-pricing/model/dto"
	. "github.com/dev-rever/cryptoo-pricing/utils"
)

type User struct {
	userRepo *repositories.User
}

func ProvideUserCtrl(repo *repositories.User) *User {
	return &User{userRepo: repo}
}

func (u *User) Root(ctx *gin.Context) {
	msg := "this is root page"
	ctx.JSON(http.StatusOK, msg)
}

func (u *User) Register(ctx *gin.Context) {
	var req model.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseError(InternalErrorCode, err.Error()))
		return
	}

	// check if the user exists
	exists, err := u.userRepo.CheckUserExists(ctx, req.Account, req.Email)
	if err != nil {
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError(DBErrorCode, "database error"))
		return
	}
	if exists {
		err := errors.New("account or email already exists")
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusConflict, ResponseError(DBConflictErrorCode, err.Error()))
		return
	}

	// hash password
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError(InternalErrorCode, err.Error()))
		return
	}

	// store user to db
	id, err := u.userRepo.InsertUser(ctx, req.Account, string(hashedPassword), req.Email)
	if err != nil {
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError(DBErrorCode, err.Error()))
		return
	}

	accToken, err := jwt.GenerateJWT(id)
	if err != nil {
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError(InternalErrorCode, err.Error()))
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

func (u *User) Login(ctx *gin.Context) {
	var req model.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseError(InternalErrorCode, err.Error()))
		return
	}

	dbPwd, err := u.userRepo.QueryUserPwdByAccount(ctx, req.Account)
	if err != nil {
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError(DBErrorCode, err.Error()))
		return
	}

	id, err := u.userRepo.QueryUserIDByAccount(ctx, req.Account)
	if err != nil {
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError(DBErrorCode, err.Error()))
		return
	}

	accToken, err := jwt.GenerateJWT(id)
	if err != nil {
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError(AuthorizedErrorCode, err.Error()))
		return
	}

	if err := ComparePassword(dbPwd, req.Password); err != nil {
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, ResponseError(AuthorizedErrorCode, err.Error()))
		return
	}

	payload := model.LoginResponse{
		Token: accToken,
	}
	msg := "user login successfully"
	LogSuc(msg)
	ctx.JSON(http.StatusOK, ResponseOK(msg, &payload))
}

func (u *User) Profile(ctx *gin.Context) {
	if uidRaw, exist := ctx.Get("uid"); !exist {
		err := errors.New("unauthorized")
		LogError(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, ResponseError(AuthorizedErrorCode, err.Error()))
		return
	} else {
		uid := uidRaw.(uint)
		payload, err := u.userRepo.QueryUserByID(ctx, uid)
		if err != nil {
			LogError(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, ResponseError(DBErrorCode, err.Error()))
			return
		}

		msg := "get user profile successfully"
		LogSuc(msg)
		ctx.JSON(http.StatusOK, ResponseOK(msg, &payload))
	}
}
