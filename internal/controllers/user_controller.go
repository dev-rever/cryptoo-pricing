package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dev-rever/cryptoo-pricing/internal/middleware/jwt"
	"github.com/dev-rever/cryptoo-pricing/repositories"
	"github.com/gin-gonic/gin"

	model "github.com/dev-rever/cryptoo-pricing/model/dto"
	api "github.com/dev-rever/cryptoo-pricing/utils/apiutils"
	logger "github.com/dev-rever/cryptoo-pricing/utils/logutils"
	putils "github.com/dev-rever/cryptoo-pricing/utils/pwdutils"
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
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.ResponseError(api.InternalErrorCode, err.Error()))
		return
	}

	// check if the user exists
	exists, err := u.userRepo.CheckUserExists(ctx, req.Account, req.Email)
	if err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.DBErrorCode, "database error"))
		return
	}
	if exists {
		err := errors.New("account or email already exists")
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusConflict, api.ResponseError(api.DBConflictErrorCode, err.Error()))
		return
	}

	// hash password
	hashedPassword, err := putils.HashPassword(req.Password)
	if err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.InternalErrorCode, err.Error()))
		return
	}

	// store user to db
	id, err := u.userRepo.InsertUser(ctx, req.Account, string(hashedPassword), req.Email)
	if err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.DBErrorCode, err.Error()))
		return
	}

	accToken, err := jwt.GenerateJWT(id)
	if err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.InternalErrorCode, err.Error()))
		return
	}

	payload := model.RegisterResponse{
		Account: req.Account,
		Email:   req.Email,
		Token:   accToken,
	}

	msg := "user registered successfully"
	logger.LogSuccess(msg, fmt.Sprintf("\naccount: %s", req.Account), fmt.Sprintf("\nemail: %s", req.Email))
	ctx.JSON(http.StatusCreated, api.ResponseOK(msg, &payload))
}

func (u *User) Login(ctx *gin.Context) {
	var req model.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.ResponseError(api.InternalErrorCode, err.Error()))
		return
	}

	dbPwd, err := u.userRepo.QueryUserPwdByAccount(ctx, req.Account)
	if err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.DBErrorCode, err.Error()))
		return
	}

	id, err := u.userRepo.QueryUserIDByAccount(ctx, req.Account)
	if err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.DBErrorCode, err.Error()))
		return
	}

	accToken, err := jwt.GenerateJWT(id)
	if err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.AuthorizedErrorCode, err.Error()))
		return
	}

	if err := putils.ComparePassword(dbPwd, req.Password); err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.ResponseError(api.AuthorizedErrorCode, err.Error()))
		return
	}

	payload := model.LoginResponse{
		Token: accToken,
	}
	msg := "user login successfully"
	logger.LogSuccess(msg, fmt.Sprintf(" account: %s", req.Account))
	ctx.JSON(http.StatusOK, api.ResponseOK(msg, &payload))
}

func (u *User) Profile(ctx *gin.Context) {
	if uidRaw, exist := ctx.Get("uid"); !exist {
		err := errors.New("unauthorized")
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, api.ResponseError(api.AuthorizedErrorCode, err.Error()))
		return
	} else {
		uid := uidRaw.(uint)
		payload, err := u.userRepo.QueryUserByID(ctx, uid)
		if err != nil {
			logger.LogError(err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.DBErrorCode, err.Error()))
			return
		}

		msg := "get user profile successfully"
		logger.LogSuccess(msg, fmt.Sprintf("\naccount: %s", payload.Account), fmt.Sprintf("\nemail: %s", payload.Email))
		ctx.JSON(http.StatusOK, api.ResponseOK(msg, &payload))
	}
}
