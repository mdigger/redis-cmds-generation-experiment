package redis_test

import (
	"context"
	"testing"

	"redis"
)

func TestCommands(t *testing.T) {
	var (
		client = new(redis.Client)
		ctx    = context.Background()
	)

	client.Do(ctx, client.B().Mset().KeyValue().Build())
}
