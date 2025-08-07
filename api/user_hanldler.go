package api

import (
	"github.com/cedrickewi/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := []types.User{
		{FirstName: "Nchia"},
		{FirstName: "Nchea"},
		{FirstName: "Nchza"},
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("single user")
}
