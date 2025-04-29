package main

import (
	"context"

	"github.com/dev-rever/cryptoo-pricing/config"
	"github.com/dev-rever/cryptoo-pricing/di"
	"github.com/dev-rever/cryptoo-pricing/internal/validator"
)

func main() {
	config.LoadEnv()
	validator.InitValidators()

	app, _ := di.InitApplication(context.Background())
	app.Router.Init()

	defer app.DB.Close(context.Background())
}
