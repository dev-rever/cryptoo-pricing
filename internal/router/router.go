package router

import (
	"log"
	"time"

	"github.com/dev-rever/cryptoo-pricing/internal/controllers"
	"github.com/dev-rever/cryptoo-pricing/internal/middleware/mredis"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	gin        *gin.Engine
	mredis     *mredis.Wrap
	jwt        func(*gin.Context)
	userCtrl   *controllers.User
	cryptoCtrl *controllers.Crypto
}

const (
	root = "/"

	// user
	userRegister = "user/register"
	userLogin    = "user/login"
	userProfile  = "user/profile"

	// crypto
	cryptoQueries       = "crypto/queries"
	cryptoPrice         = "crypto/price"
	cryptoSupCurrencies = "crypto/supcurrencies"
)

func ProvideRouter(
	userCtrl *controllers.User,
	cryptoCtrl *controllers.Crypto,
	mredis *mredis.Wrap,
	jwt gin.HandlerFunc,
) *Engine {
	return &Engine{
		gin:        gin.Default(),
		mredis:     mredis,
		jwt:        jwt,
		userCtrl:   userCtrl,
		cryptoCtrl: cryptoCtrl,
	}
}

func (e *Engine) Init() {
	jwt := e.gin.Group("", e.jwt)

	// GET
	e.gin.GET(root, e.userCtrl.Home)
	jwt.GET(userProfile, e.userCtrl.Profile)
	jwt.GET(cryptoQueries, e.cryptoCtrl.SearchCoin)
	jwt.GET(cryptoSupCurrencies, e.cryptoCtrl.SupportCurrencies)
	jwt.GET(cryptoPrice, e.cryptoCtrl.GetCryptoPriceByIDs)

	// POST
	e.gin.POST(userRegister, e.userCtrl.Register)
	e.gin.POST(userLogin, e.mredis.RateLimitMiddleware(5, time.Minute), e.userCtrl.Login)
	if err := e.gin.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
