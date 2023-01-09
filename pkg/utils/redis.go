package utils

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/akorol1998/go-auth-service/pkg/config"
	"github.com/go-redis/redis"
)

type RedisHandler struct {
	client *redis.Client
}

func RedisInit(c config.Config) RedisHandler {
	addr := fmt.Sprintf("%v:%v", c.RedisHost, c.RedisPort)
	pass := c.RedisPassword
	db, _ := strconv.Atoi(c.RedisDb)

	log.Printf("Redis connection string - addr: %v, pass: %v, db: %v", addr, pass, db)
	cl := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	return RedisHandler{client: cl}
}

func (r *RedisHandler) Get(key string) (string, error) {
	cmd := r.client.Get(key)
	val, err := cmd.Result()
	return val, err
}

func (r *RedisHandler) Set(k string, val interface{}, exp time.Duration) error {
	log.Println("Setting redis value - key:", k)
	status := r.client.Set(k, val, exp)
	_, err := status.Result()
	return err
}

func MakeRedisKey(tknType TokenType, identifier string) string {
	return fmt.Sprintf("user:token_%v:%v", tknType, identifier)
}
