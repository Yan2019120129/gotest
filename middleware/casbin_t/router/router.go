package router

import (
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"

	"gotest/middleware/casbin_t/api"
	"gotest/middleware/casbin_t/service"
)

func Register(app *fiber.App, handler *api.Handler, auth *service.AuthService, enforcer *casbin.Enforcer) {
	app.Get("/health", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) })
	app.Post("/api/auth/login", handler.Login)

	protected := app.Group("/api", authenticate(auth), authorize(enforcer))
	protected.Get("/tests", handler.ListTests)
	protected.Post("/tests", handler.CreateTest)
	protected.Get("/tests/:id", handler.GetTest)
	protected.Put("/tests/:id", handler.UpdateTest)
	protected.Delete("/tests/:id", handler.DeleteTest)
}

func authenticate(auth *service.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		parts := strings.Fields(c.Get(fiber.HeaderAuthorization))
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Bearer token is required"})
		}
		claims, err := auth.ParseToken(parts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid or expired token"})
		}
		api.SetUsername(c, claims.Username)
		c.Locals("casbin_t_role", claims.Role)
		return c.Next()
	}
}

func authorize(enforcer *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, _ := c.Locals("casbin_t_role").(string)
		allowed, err := enforcer.Enforce(role, c.Path(), c.Method())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "authorization check failed"})
		}
		if !allowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "permission denied"})
		}
		return c.Next()
	}
}
