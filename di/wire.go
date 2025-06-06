//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/dev-rever/cryptoo-pricing/internal/controllers"
	"github.com/dev-rever/cryptoo-pricing/internal/db"
	"github.com/dev-rever/cryptoo-pricing/internal/middleware/jwt"
	"github.com/dev-rever/cryptoo-pricing/internal/middleware/mredis"
	"github.com/dev-rever/cryptoo-pricing/internal/router"
	"github.com/dev-rever/cryptoo-pricing/repositories"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5"
)

var MiddlewareSet = wire.NewSet(
	jwt.ProvideJWTMiddleware,
	mredis.ProvideMRedis,
)

var controllerSet = wire.NewSet(
	controllers.ProvideUserCtrl,
	controllers.ProvideCryptoCtrl,
)

var repositorySet = wire.NewSet(
	repositories.ProvideUserRepo,
	repositories.ProvideCryptoRepo,
)

type Application struct {
	Router *router.Engine
	DB     *pgx.Conn
}

func InitApplication(ctx context.Context) (*Application, error) {
	wire.Build(
		db.ProvideDB,
		repositorySet,
		controllerSet,
		MiddlewareSet,
		router.ProvideRouter,
		wire.Struct(new(Application), "*"),
	)
	return &Application{}, nil
}
