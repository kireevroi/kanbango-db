package cache

import (
	"time"

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
	_, err := c.Ping().Result(); 
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
