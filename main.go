package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	redis "github.com/redis/go-redis/v9"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	rdb := redis.NewClient(&redis.Options{
		Addr:     "moonveil-brocoli-redis.bnfvce.0001.apse1.cache.amazonaws.com:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("ping...")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("redis/set", func(c *gin.Context) {
		fmt.Println("redis set...")
		err := rdb.Set(context.TODO(), "test_key", uuid.NewString(), 0).Err()
		if err != nil {
			fmt.Println("error:", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("redis/get", func(c *gin.Context) {
		fmt.Println("redis get...")
		s, err := rdb.Get(context.TODO(), "test_key").Result()
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Println("get:", s)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":80")
}
