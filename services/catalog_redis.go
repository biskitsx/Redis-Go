package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"www.github.com/biskitsx/Redis-Go/repositories"
)

type catalogServiceRedis struct {
	productRepo repositories.ProductRepository
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productRepo repositories.ProductRepository, redisClient *redis.Client) CatalogService {
	return &catalogServiceRedis{productRepo, redisClient}
}

func (s catalogServiceRedis) GetProducts() (products []Product, err error) {
	key := "service::GetAll"

	// Redis Get
	if productJson, err := s.redisClient.Get(context.Background(), key).Result(); err == nil {
		if err := json.Unmarshal([]byte(productJson), &products); err == nil {
			return products, nil
		}
	}

	// Repository
	productsDB, err := s.productRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, product := range productsDB {
		products = append(products, Product(product))
	}

	// Redis Set

	if data, err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	return products, err
}
