package repositories

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) ProductRepository {
	db.AutoMigrate(&product{})
	mockData(db)
	return &productRepositoryRedis{db, redisClient}
}

func (r productRepositoryRedis) GetAll() (products []product, err error) {
	// Redis Get
	key := "repository::GetAll"
	productJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		if err = json.Unmarshal([]byte(productJson), &products); err == nil {
			return products, nil
		}
	}

	// Database
	if err = r.db.Order("quantity desc").Limit(30).Find(&products).Error; err != nil {
		return nil, err
	}

	// Redis Set (Caching)
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}
	r.redisClient.Set(context.Background(), key, string(data), time.Second*10)

	return products, nil
}
