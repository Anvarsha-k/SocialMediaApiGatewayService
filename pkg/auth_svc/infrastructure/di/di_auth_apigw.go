package di_auth_apigw

import (
	"log"

	"github.com/gofiber/fiber/v2"
	handler_auth_apigw "github.com/shaan/socialMediaApiGateway/pkg/auth_svc/handler"
	client_apigw_auth "github.com/shaan/socialMediaApiGateway/pkg/auth_svc/infrastructure/client"
	router_auth_apigw "github.com/shaan/socialMediaApiGateway/pkg/auth_svc/infrastructure/router"
	config_apigw "github.com/shaan/socialMediaApiGateway/pkg/config"
)

func InitAuthClient(app *fiber.App, config *config_apigw.Config) error {
	//Client Initialisation
	client, err := client_apigw_auth.InitAuthClient(config)
	if err != nil {
		log.Fatal(err)
	}
	//Handler initialisation
	userHandler := handler_auth_apigw.NewUserHandler(client)

	router_auth_apigw.AuthUserRoutes(app,userHandler)

	return nil
}
