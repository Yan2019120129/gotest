package api

import "github.com/gofiber/fiber/v2"

const usernameKey = "casbin_t_username"

func SetUsername(c *fiber.Ctx, username string) { c.Locals(usernameKey, username) }

func Username(c *fiber.Ctx) string {
	username, _ := c.Locals(usernameKey).(string)
	return username
}
