package route

import (
	"workshop-vhiweb/handler"

	"github.com/gofiber/fiber/v2"
)

func Routeinit(r *fiber.App) {
	r.Get("/user", handler.UserHandlerRead)
	
}