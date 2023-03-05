package main

import (
	"github.com/divrhino/divrhino-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	//home
	app.Get("/", handlers.Home)

	app.Post("/addFacts", handlers.CreateFact)
	app.Put("/updateFacts/:id", handlers.UpdateFact)
	app.Delete("/deleteFacts/:id", handlers.DeleteFact)
	app.Get("/getFacts", handlers.GetFacts)
}
