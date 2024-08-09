package utils

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func Cache() *redis.Client {
	DB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Printf("%s:Error loading .env file", "main.go")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       DB,
	})

	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Error connecting to Redis: %v", err)
		return nil
	}

	log.Printf("Connected to Redis successfully")

	return rdb
}

func Set(c *redis.Client, ctx context.Context, key string, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	res := c.Set(ctx, key, p, 0)
	return res.Err()
}

func Get(c *redis.Client, ctx context.Context, key string, dest interface{}) error {
	p, err := c.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(p), dest)
}
