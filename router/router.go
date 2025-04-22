package router

import (
	"github.com/dev-rever/cryptoo-pricing/controller"
	"github.com/gin-gonic/gin"
)

const (
	Root     = "/"
	Register = "/user/register"
)

func ProvideRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	// GET
	router.GET(Root, userController.Root)

	// POST
	router.POST(Register, userController.Register)

	return router
}
