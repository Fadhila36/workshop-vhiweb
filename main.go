package main

import (
	"workshop-vhiweb/database"
	"workshop-vhiweb/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// initial database
	database.DatabaseInit()
	
    app := fiber.New()

	// INITIAL ROUTE
	route.Routeinit(app)

	app.Listen(":3000")
}