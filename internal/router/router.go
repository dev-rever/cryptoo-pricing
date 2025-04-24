package router

import (
	"github.com/dev-rever/cryptoo-pricing/internal/controller"
	"github.com/gin-gonic/gin"
)

const (
	root     = "/"
	register = "/user/register"
	profile  = "/user/profile"
)

func ProvideRouter(userController *controller.UserController, jwt gin.HandlerFunc) *gin.Engine {
	router := gin.Default()

	authRouter := router.Group("", jwt)

	// GET
	router.GET(root, userController.Root)
	authRouter.GET(profile, userController.Profile)

	// POST
	router.POST(register, userController.Register)

	return router
}
