package routes

import (
	"campus-api/controller"
	"campus-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Use(middleware.CorsMiddleware())
	// app.Post("/api/sendmail", controller.SendMail)
	app.Post("/api/rsvp", controller.CreateReservation)
	app.Post("/api/admreg", controller.RegisterAdmin)
	app.Post("/api/admlog", controller.LoginAdmin)

	app.Use(middleware.IsAuthenticate)
	app.Get("/api/getid/:id", controller.GetId)
	app.Get("/api/getuser", controller.GetUser)

}
