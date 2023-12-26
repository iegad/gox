package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewRedisClient(c *RedisConfig) (*redis.Client, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Username: c.Username,
		Password: c.Password,
	})
	err := r.Ping(context.TODO()).Err()
	if err != nil {
		r.Close()
		return nil, err
	}

	return r, nil
}
