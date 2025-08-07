package main

import (
	"flag"

	"github.com/cedrickewi/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The listen addr of the api server")
	flag.Parse()
	app := fiber.New()
	apiV1 := app.Group("api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
