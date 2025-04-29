package controllers

import (
	"errors"
	"net/http"

	"github.com/dev-rever/cryptoo-pricing/repositories"
	"github.com/gin-gonic/gin"

	api "github.com/dev-rever/cryptoo-pricing/utils/apiutils"
	logger "github.com/dev-rever/cryptoo-pricing/utils/logutils"
)

type Crypto struct {
	cryptoRepo *repositories.CryptoRepo
}

func ProvideCryptoCtrl(repo *repositories.CryptoRepo) *Crypto {
	return &Crypto{cryptoRepo: repo}
}

func (c *Crypto) SearchQueries(ctx *gin.Context) {
	coin, exist := ctx.GetQuery("coinName")
	if !exist {
		err := errors.New("Missing \"coinName\" query params")
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.ResponseError(api.InternalErrorCode, err.Error()))
		return
	}

	result, err := c.cryptoRepo.SearchQueries(coin)
	if err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.InternalErrorCode, err.Error()))
		return
	}

	logger.LogAsJSON(result)
	ctx.JSON(http.StatusOK, api.ResponseOK(api.SuccessText, result))
}
