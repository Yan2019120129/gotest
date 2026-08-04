package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"gotest/middleware/casbin_t/service"
)

type Handler struct {
	auth  *service.AuthService
	tests *service.TestService
}

func NewHandler(auth *service.AuthService, tests *service.TestService) *Handler {
	return &Handler{auth: auth, tests: tests}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil || input.Username == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "username and password are required"})
	}
	token, err := h.auth.Login(input.Username, input.Password)
	if errors.Is(err, service.ErrInvalidCredentials) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "login failed"})
	}
	return c.JSON(fiber.Map{"access_token": token, "token_type": "Bearer"})
}

func (h *Handler) CreateTest(c *fiber.Ctx) error {
	input, err := parseTestInput(c)
	if err != nil {
		return err
	}
	test, err := h.tests.Create(input, Username(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "create test failed"})
	}
	return c.Status(fiber.StatusCreated).JSON(test)
}

func (h *Handler) ListTests(c *fiber.Ctx) error {
	tests, err := h.tests.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "list tests failed"})
	}
	return c.JSON(tests)
}

func (h *Handler) GetTest(c *fiber.Ctx) error {
	test, err := h.tests.Get(id(c))
	return sendTest(c, test, err)
}

func (h *Handler) UpdateTest(c *fiber.Ctx) error {
	input, err := parseTestInput(c)
	if err != nil {
		return err
	}
	test, err := h.tests.Update(id(c), input)
	return sendTest(c, test, err)
}

func (h *Handler) DeleteTest(c *fiber.Ctx) error {
	if err := h.tests.Delete(id(c)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "test not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "delete test failed"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func parseTestInput(c *fiber.Ctx) (service.TestInput, error) {
	var input service.TestInput
	if err := c.BodyParser(&input); err != nil || input.Name == "" {
		return input, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "name is required"})
	}
	return input, nil
}

func id(c *fiber.Ctx) uint {
	value, _ := strconv.ParseUint(c.Params("id"), 10, 64)
	return uint(value)
}

func sendTest(c *fiber.Ctx, test interface{}, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "test not found"})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "test operation failed"})
	}
	return c.JSON(test)
}
