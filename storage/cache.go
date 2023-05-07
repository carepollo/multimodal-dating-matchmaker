package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

// a custom wrapper for common actions
type Cache struct {
	client  *redis.Client   // the redis official client
	context context.Context // mongodb has its own context, redis also its own required context
}

func NewRedis() *Cache {
	cache := &Cache{
		client: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_CONNECTION"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
			Username: "default",
		}),
		context: context.TODO(),
	}
	fmt.Println("Connected to cache succesfully")
	return cache
}

// returns the value requested while might return error if not found
func (cache *Cache) Save(key string) (interface{}, error) {
	return nil, nil
}
