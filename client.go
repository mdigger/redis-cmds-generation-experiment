package redis

import (
	"context"
	"time"

	"redis/cmd"
)

type RedisResult struct{}

type Client struct {
	cmd cmd.Builder
}

func (c *Client) B() cmd.Builder {
	return c.cmd
}

func (c *Client) Do(ctx context.Context, cmd cmd.Completed) *RedisResult {
	return nil
}

func (c *Client) DoMulti(ctx context.Context, multi ...cmd.Completed) []RedisResult {
	return nil
}

func (c *Client) DoCache(ctx context.Context, cmd cmd.Cacheable, ttl time.Duration) *RedisResult {
	return nil
}
