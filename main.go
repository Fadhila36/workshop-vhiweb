package main

import (
	"workshop-vhiweb/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

	// INITIAL ROUTE
	route.Routeinit(app)

	app.Listen(":3000")
}