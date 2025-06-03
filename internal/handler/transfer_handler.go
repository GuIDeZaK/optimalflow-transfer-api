// Directory: internal/handler/transfer_handler.go
package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/guide-backend/internal/service"
)

type TransferHandler struct {
	transferService service.TransferService
}

func NewTransferHandler(ts service.TransferService) TransferHandler {
	return TransferHandler{transferService: ts}
}

func (h TransferHandler) Transfer(c *fiber.Ctx) error {
	userIDVal := c.Locals("user_id")
	userID, ok := userIDVal.(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user context")
	}

	var req service.TransferRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	req.FromUserID = userID

	resp, err := h.transferService.Transfer(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(resp)
}
