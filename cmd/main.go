package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	di_auth_apigw "github.com/shaan/socialMediaApiGateway/pkg/auth_svc/infrastructure/di"
	config_apigw "github.com/shaan/socialMediaApiGateway/pkg/config"
)

func main() {
	config, err := config_apigw.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	err = di_auth_apigw.InitAuthClient(app, config)
	if err != nil {
		log.Fatal(err)
	}

	err = app.Listen(config.Port)
	if err != nil {
		fmt.Printf("Error starting Server: %v\n", err)
	}
}
