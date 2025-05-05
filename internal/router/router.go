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
	userRegister = "/user/register"
	userLogin    = "user/login"
	userProfile  = "/user/profile"

	// crypto
	cryptoQueries = "crypto/queries"
	cryptoPrice   = "crypto/price"
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
	je := e.gin.Group("", e.jwt)

	// GET
	e.gin.GET(root, e.userCtrl.Root)
	je.GET(userProfile, e.userCtrl.Profile)
	je.GET(cryptoQueries, e.cryptoCtrl.SearchCoin)
	je.GET(cryptoPrice, e.cryptoCtrl.GetCryptoPriceByIDs)

	// POST
	e.gin.POST(userRegister, e.userCtrl.Register)
	e.gin.POST(userLogin, e.mredis.RateLimitMiddleware(5, time.Minute), e.userCtrl.Login)
	if err := e.gin.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
