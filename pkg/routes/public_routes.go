package routes

import (
	"github.com/Jneville0815/fitness-app-backend-v2/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("api/v1")

	route.Post("/user/:user_id/token/renew", controllers.RenewTokens)
	route.Post("/user/register", controllers.UserRegister)
	route.Post("/user/login", controllers.UserLogin)
}
