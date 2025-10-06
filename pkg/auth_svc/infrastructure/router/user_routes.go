package router_auth_apigw

import (
	"github.com/gofiber/fiber/v2"
	handler_auth_apigw "github.com/shaan/socialMediaApiGateway/pkg/auth_svc/handler"
)

func AuthUserRoutes(app *fiber.App, userHandler *handler_auth_apigw.UserHandler) {
	app.Post("/signup", userHandler.UserSignUp)
	app.Post("/login", userHandler.UserLogin)
	app.Get("/ping", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "pong",
        })
	})
	app.Post("/verify",userHandler.UserOTPVerication)
	app.Post("/forgotpassword",userHandler.ForgotPasswordRequest)
	app.Post("/resetpassword",userHandler.ResetPassword)

	
}
