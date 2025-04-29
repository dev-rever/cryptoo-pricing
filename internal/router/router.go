package router

import (
	"log"
	"time"

	"github.com/dev-rever/cryptoo-pricing/internal/controllers"
	"github.com/dev-rever/cryptoo-pricing/internal/middleware/mredis"
	"github.com/gin-gonic/gin"
)

type Engine struct {
	gin      *gin.Engine
	mredis   *mredis.Wrap
	jwt      func(*gin.Context)
	userCtrl *controllers.User
}

const (
	root = "/"
	// user
	userRegister = "/user/register"
	userLogin    = "user/login"
	userProfile  = "/user/profile"
)

// func ProvideRouter(userController *controllers.User, mredis *mredis.Wrap, jwt gin.HandlerFunc) *gin.Engine {
// 	router := gin.Default()

// 	authRouter := router.Group("", jwt)

// 	// GET
// 	router.GET(root, userController.Root)
// 	authRouter.GET(userProfile, userController.Profile)

// 	// POST
// 	router.POST(userRegister, userController.Register)
// 	router.POST(userLogin, mredis.RateLimitMiddleware(5, time.Minute), userController.Login)

// 	return router
// }

func ProvideRouter(userCtrl *controllers.User, mredis *mredis.Wrap, jwt gin.HandlerFunc) *Engine {
	return &Engine{
		gin:      gin.Default(),
		mredis:   mredis,
		jwt:      jwt,
		userCtrl: userCtrl,
	}
}

func (e *Engine) Init() {
	je := e.gin.Group("", e.jwt)

	// GET
	e.gin.GET(root, e.userCtrl.Root)
	je.GET(userProfile, e.userCtrl.Profile)

	// POST
	e.gin.POST(userRegister, e.userCtrl.Register)
	e.gin.POST(userLogin, e.mredis.RateLimitMiddleware(5, time.Minute), e.userCtrl.Login)
	if err := e.gin.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
