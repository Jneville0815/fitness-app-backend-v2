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

	route.Put("/user/workout", middleware.JWTProtected(), controllers.UpdateWorkout)
	route.Put("/user/workout/note", middleware.JWTProtected(), controllers.UpdateNote)
	route.Put("/user/workout/current-day", middleware.JWTProtected(), controllers.UpdateCurrentDay)

	route.Get("/user/workout", middleware.JWTProtected(), controllers.GetWorkout)
	route.Get("/user/workout/note", middleware.JWTProtected(), controllers.GetNote)
	route.Get("/user/workout/current-day", middleware.JWTProtected(), controllers.GetCurrentDay)
	route.Get("/user/quotes", middleware.JWTProtected(), controllers.GetQuotes)

	route.Delete("/user/quotes/:quote_id", middleware.JWTProtected(), controllers.DeleteQuote)
}
