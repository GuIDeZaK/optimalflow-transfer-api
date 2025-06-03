package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guide-backend/internal/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{userService: userService}
}

// func (h UserHandler) GetUserInfoByID(c *fiber.Ctx) error {
// 	id, err := c.ParamsInt("id")
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, "wrong id")
// 	}
// 	resp, err := h.userService.GetUserInfoByID(id)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, "no information")
// 	}

// 	return c.JSON(resp)
// }

func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	var req service.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON body")
	}

	resp, err := h.userService.CreateUser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h UserHandler) Login(c *fiber.Ctx) error {
	var req service.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request")
	}

	resp, err := h.userService.Login(req)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.JSON(resp)
}

func (h UserHandler) ListAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.ListAllUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(users)
}

func (h UserHandler) GetUserByID(c *fiber.Ctx) error {
	idParam, err := c.ParamsInt("id")
	if err != nil || idParam <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	resp, err := h.userService.GetUserByID(uint(idParam))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.JSON(resp)
}
