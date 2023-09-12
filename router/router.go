package router

import (
	"gomodApp/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(app * fiber.App){
	app.Get("/",controller.Guerydata()) 
	app.Post("/insert",controller.InsertData()) 
}