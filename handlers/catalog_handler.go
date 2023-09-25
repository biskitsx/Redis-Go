package handlers

import (
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/Redis-Go/services"
)

type catalogHandler struct {
	catalogService services.CatalogService
}

func NewCatalogHandler(catalogService services.CatalogService) CatalogHandler {
	return &catalogHandler{catalogService}
}

func (h catalogHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.catalogService.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}
	return c.JSON(response)
}
