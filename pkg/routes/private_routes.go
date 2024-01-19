package routes

import (
	"github.com/Jneville0815/fitness-app-backend-v2/app/controllers"
	"github.com/Jneville0815/fitness-app-backend-v2/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/user/logout", middleware.JWTProtected(), controllers.UserLogout)
	route.Post("/user/workout", middleware.JWTProtected(), controllers.CreateWorkout)
	route.Post("/user/quotes", middleware.JWTProtected(), controllers.CreateQuote)

	route.Get("/user/workout", middleware.JWTProtected(), controllers.GetWorkout)
}
