package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	redisAdder    = "localhost:6379"
	redisPassword = ""
	redisDb       = 0
	serverPort    = ":8000"
	tokenPrefix   = "token: "
)

var (
	redisClient *redis.Client
)

func loginHandler(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return c.JSON(400, gin.H{"msg": "error"})
	}
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAdder,
		Password: redisPassword,
		DB:       redisDb,
	})

	r := gin.New()
	r.POST("/login", loginHandler())

}
