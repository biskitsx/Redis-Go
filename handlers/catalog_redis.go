package handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"www.github.com/biskitsx/Redis-Go/services"
)

type catalogHandlerRedis struct {
	catalogService services.CatalogService
	redisClient    *redis.Client
}

func NewCatalogHandlerRedis(catalogService services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return &catalogHandlerRedis{catalogService, redisClient}
}

func (h catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {

	key := "handler::GetAll"

	// Redis GET
	if reponseJson, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		c.Set("Content-Type", "application/json")
		return c.SendString(reponseJson)
	}

	products, err := h.catalogService.GetProducts()
	if err != nil {
		return err
	}
	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}

	// Redis Set
	if data, err := json.Marshal(products); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	return c.JSON(response)
}
