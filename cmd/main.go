package main

import (
	"context"
	"log"

	"github.com/dev-rever/cryptoo-pricing/config"
	"github.com/dev-rever/cryptoo-pricing/di"
	"github.com/dev-rever/cryptoo-pricing/internal/validator"
)

func main() {
	config.LoadEnv()
	validator.InitValidators()

	app, _ := di.InitApplication(context.Background())
	defer app.DB.Close(context.Background())

	if err := app.Router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
