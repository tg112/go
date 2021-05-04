package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tg112/go/go-admin/controllers"
	"github.com/tg112/go/go-admin/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)

	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.User)
}
