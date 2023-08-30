package handlers

import (
	"resqiar.com-server/services"

	"github.com/gofiber/fiber/v2"
)

type ParserHandler interface {
	ParseMDtoHTML(c *fiber.Ctx) error
}

type ParserHandlerImpl struct {
	ParserService services.ParserService
}

func (s *ParserHandlerImpl) ParseMDtoHTML(c *fiber.Ctx) error {
	payload := c.Body()
	parsed := s.ParserService.ParseMDByte(payload)
	if parsed == nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	c.Type("application/octet-stream")
	return c.Status(fiber.StatusOK).Send(parsed)
}
