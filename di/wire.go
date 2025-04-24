//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/dev-rever/cryptoo-pricing/internal/controller"
	"github.com/dev-rever/cryptoo-pricing/internal/db"
	"github.com/dev-rever/cryptoo-pricing/internal/middleware"
	"github.com/dev-rever/cryptoo-pricing/internal/router"
	"github.com/dev-rever/cryptoo-pricing/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5"
)

var MiddlewareSet = wire.NewSet(
	middleware.ProvideJWTMiddleware,
)

type Application struct {
	Router *gin.Engine
	DB     *pgx.Conn
}

func InitApplication(ctx context.Context) (*Application, error) {
	wire.Build(
		db.ProvideDB,
		repository.ProvideUserRepo,
		controller.ProvideUserCtrl,
		MiddlewareSet,
		router.ProvideRouter,
		wire.Struct(new(Application), "*"),
	)
	return &Application{}, nil
}
