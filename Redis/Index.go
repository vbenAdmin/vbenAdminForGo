package Redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"goVben/Utils/Config"
	"time"
)

var rdb *redis.Client
var ctx context.Context

func Init() error {
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     Config.Conf.Redis.Url,
		Password: Config.Conf.Redis.Pwd, // no password set
		DB:       Config.Conf.Redis.DB,  // use default DB
	})
	return rdb.Ping(ctx).Err()
}

func Set(key string, value any, expAt time.Duration) error {
	return rdb.Set(ctx, key, value, expAt).Err()
}

func Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}
