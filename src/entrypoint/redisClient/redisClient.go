package redisClient

import (
	"context"
	"strconv"

	"errors"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Server   string `koanf:"server"`
	Port     int    `koanf:"port"`
	Password string `koanf:"password"`
	Database int    `koanf:"database"`
	client   *redis.Client
}

var ctx = context.Background()

func (rc *RedisConfig) Enabled() bool {
	return rc != nil && len(rc.Server) > 0
}

func (rc *RedisConfig) Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     rc.Server + ":" + strconv.Itoa(rc.Port),
		Password: rc.Password,
		DB:       rc.Database,
	})
	rc.client = client
}

func (rc *RedisConfig) SetBlock(value string, entries int) {
	if entries > 0 {
		rc.client.Set(ctx, Prefix("Block_Value"), value, 0).Err()
		rc.client.Set(ctx, Prefix("Block_Count"), strconv.Itoa(entries), 0).Err()
	}
}

func (rc *RedisConfig) GetBlock() (string, int, error) {
	val, valErr := rc.client.Get(ctx, "Block_Value").Result()
	entries, entriesErr := rc.client.Get(ctx, "Block_Count").Int()
	if valErr == nil && entriesErr == nil {
		return val, entries, nil
	} else {
		return "", 0, errors.New("Coulden't get blocklist from redis")
	}
}

func (rc *RedisConfig) SetAllow(value string, entries int) {
	if entries > 0 {
		rc.client.Set(ctx, Prefix("Allow_Value"), value, 0).Err()
		rc.client.Set(ctx, Prefix("Allow_Count"), entries, 0).Err()
	}
}

func (rc *RedisConfig) GetAllow() (string, int, error) {
	val, valErr := rc.client.Get(ctx, "Allow_Value").Result()
	entries, entriesErr := rc.client.Get(ctx, "Allow_Count").Int()
	if valErr == nil && entriesErr == nil {
		return val, entries, nil
	} else {
		return "", 0, errors.New("Coulden't get blocklist from redis")
	}
}

func Prefix(key string) string {
	return "AdBlockLists-" + key
}
