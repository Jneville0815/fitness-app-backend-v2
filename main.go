package main

import (
	"fmt"
	"github.com/Jneville0815/fitness-app-backend-v2/pkg/middleware"
	"github.com/Jneville0815/fitness-app-backend-v2/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	app := fiber.New()

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	err = app.Listen("127.0.0.1:9991")
	if err != nil {
		fmt.Println(err)
		return
	}
}
