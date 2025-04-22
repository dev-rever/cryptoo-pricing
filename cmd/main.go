package main

import (
	"context"
	"log"

	"github.com/dev-rever/cryptoo-pricing/config"
	"github.com/dev-rever/cryptoo-pricing/di"
)

func main() {
	config.LoadEnv()
	app, _ := di.InitApplication(context.Background())
	defer app.DB.Close(context.Background())

	if err := app.Router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
