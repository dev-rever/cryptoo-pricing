package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/dev-rever/cryptoo-pricing/model"
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

// query params
// blur: boolean
// coinName: string
func (c *Crypto) SearchCoin(ctx *gin.Context) {
	var blurSearch = false
	if blurValue, exist := ctx.GetQuery("blur"); exist == true {
		result, err := strconv.ParseBool(blurValue)
		if err == nil {
			blurSearch = result
		} else {
			err = errors.New("Query \"blur\" must be of boolean type")
			logger.LogError(err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, api.ResponseError(api.InternalErrorCode, err.Error()))
		}
	}

	cname, exist := ctx.GetQuery("coinName")
	if !exist {
		err := errors.New("Missing \"coinName\" query params")
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, api.ResponseError(api.InternalErrorCode, err.Error()))
		return
	}

	result, err := c.cryptoRepo.SearchQueries(cname)
	if err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.InternalErrorCode, err.Error()))
		return
	}

	if !blurSearch {
		for _, coin := range result.Coins {
			if strings.EqualFold(coin.Name, cname) {
				result = &model.CryptoSearchQueries{
					Coins: []model.Coin{coin},
				}
				break
			}
		}
	}

	logger.LogAsJSON(result)
	ctx.JSON(http.StatusOK, api.ResponseOK(api.SuccessText, result))
}

func (c *Crypto) SupportCurrencies(ctx *gin.Context) {

}

func (c *Crypto) GetCryptoPriceByIDs(ctx *gin.Context) {
	var currencies []string
	var coinIDs []string
	if currStr, exist := ctx.GetQuery("currencies"); exist {
		currencies = strings.Split(currStr, ",")
	}
	if coinStr, exist := ctx.GetQuery("coinIDs"); exist {
		coinIDs = strings.Split(coinStr, ",")
	}

	result, err := c.cryptoRepo.CoinPriceByIDs(currencies, coinIDs)
	if err != nil {
		logger.LogError(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, api.ResponseError(api.InternalErrorCode, err.Error()))
		return
	}

	logger.LogAsJSON(result)
	ctx.JSON(http.StatusOK, api.ResponseOK(api.SuccessText, result))
}
