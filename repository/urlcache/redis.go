package url

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"shorturl/entity"
	"time"
)

type RedisCache struct {
	ctx    context.Context
	Client *redis.Client
}

func New(addr string) (cache *RedisCache, err error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	ctx := context.Background()

	return &RedisCache{
		ctx: ctx,
		Client: cli,
	}, nil
}

func (c *RedisCache) Read(ID string) (ent *entity.URL) {
	res, err := c.Client.Get(c.ctx, ID).Result()
	if err != nil {
		return nil
	}

	var url entity.URL
	err = json.Unmarshal([]byte(res), &url)
	if err != nil {
		return nil
	}

	return &url
}

func (c *RedisCache) Write(ent *entity.URL) (err error) {
	bytes, err := json.Marshal(ent)
	return c.Client.Set(c.ctx, ent.ID, string(bytes), 24 * time.Hour).Err()
}
