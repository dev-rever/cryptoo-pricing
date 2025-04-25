package router

import (
	"time"

	"github.com/dev-rever/cryptoo-pricing/internal/controller"
	"github.com/dev-rever/cryptoo-pricing/internal/middleware/redisutil"
	"github.com/gin-gonic/gin"
)

const (
	root = "/"
	// user
	userRegister = "/user/register"
	userLogin    = "user/login"
	userProfile  = "/user/profile"
)

func ProvideRouter(userController *controller.UserController, redis *redisutil.MRedis, jwt gin.HandlerFunc) *gin.Engine {
	router := gin.Default()

	authRouter := router.Group("", jwt)

	// GET
	router.GET(root, userController.Root)
	authRouter.GET(userProfile, userController.Profile)

	// POST
	router.POST(userRegister, userController.Register)
	router.POST(userLogin, redis.RateLimitMiddleware(5, time.Minute), userController.Login)

	return router
}
