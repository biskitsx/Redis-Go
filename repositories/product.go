package repositories

import (
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type ProductRepository interface {
	GetAll() ([]product, error)
}

func mockData(db *gorm.DB) error {
	var count int64

	db.Model(&product{}).Count(&count)
	if count > 0 {
		return nil
	}
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	products := []product{}
	for i := 0; i < 5000; i++ {
		products = append(products, product{
			Name:     fmt.Sprintf("Product%v", i+1),
			Quantity: random.Intn(100),
		})
	}
	return db.Create(&products).Error
}
