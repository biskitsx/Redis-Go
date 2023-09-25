package main

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"www.github.com/biskitsx/Redis-Go/handlers"
	"www.github.com/biskitsx/Redis-Go/repositories"
	"www.github.com/biskitsx/Redis-Go/services"
)

func initServer() *fiber.App {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		time.Sleep(time.Millisecond * 10)
		return c.SendString("hello world")
	})
	return app
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:password@tcp(localhost:3306)/mariaja")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
func main() {

	app := initServer()

	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient

	productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogService(productRepo)
	productHandler := handlers.NewCatalogHandlerRedis(productService, redisClient)

	app.Get("/products", productHandler.GetProducts)
	app.Listen(":8080")

}
