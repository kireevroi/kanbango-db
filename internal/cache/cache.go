package cache

import (
	"time"
	"log"
	"github.com/go-redis/redis"
)

type Cache struct {
	*redis.Client
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Connect(url string) error {
	c.Client = redis.NewClient(&redis.Options{
		Addr: url,
		Password: "",
		DB: 0,
	})
	pong, err := c.Ping().Result();
	if pong == "PONG" {
		log.Println("Connected to Redis DB")
	}

	return err
}

func (c *Cache) NewSession(key, value string) {
	m, _ := time.ParseDuration("60m")
	c.Set(key, value, m)
}

func (c *Cache) GetSession(key string) (string, error) {
	val, err := c.Get(key).Result()
	return val, err
}

func (c *Cache) DeleteSession(key string) (int, error) {
	val, err := c.Del(key).Result()
	return int(val), err
}

