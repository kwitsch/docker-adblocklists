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
	p := rc.Port
	if p <= 0 {
		p = 6379
	}
	d := rc.Database
	if d < 0 {
		d = 0
	}
	client := redis.NewClient(&redis.Options{
		Addr:     rc.Server + ":" + strconv.Itoa(p),
		Password: rc.Password,
		DB:       d,
	})
	rc.client = client
}

func (rc *RedisConfig) SetBlock(value string, entries int) {
	if entries > 0 {
		rc.setSVal("Block_Value", value)
		rc.setIVal("Block_Count", entries)
	}
}

func (rc *RedisConfig) GetBlock() (string, int, error) {
	val, valErr := rc.getSVal("Block_Value")
	entries, entriesErr := rc.getIVal("Block_Count")
	if valErr == nil && entriesErr == nil {
		return val, entries, nil
	} else {
		return "", 0, errors.New("Coulden't get blocklist from redis")
	}
}

func (rc *RedisConfig) SetAllow(value string, entries int) {
	if entries > 0 {
		rc.setSVal("Allow_Value", value)
		rc.setIVal("Allow_Count", entries)
	}
}

func (rc *RedisConfig) GetAllow() (string, int, error) {
	val, valErr := rc.getSVal("Allow_Value")
	entries, entriesErr := rc.getIVal("Allow_Count")
	if valErr == nil && entriesErr == nil {
		return val, entries, nil
	} else {
		return "", 0, errors.New("Coulden't get blocklist from redis")
	}
}

func (rc *RedisConfig) setSVal(key, value string) error {
	return rc.client.Set(ctx, Prefix(key), value, 0).Err()
}
func (rc *RedisConfig) getSVal(key string) (string, error) {
	return rc.client.Get(ctx, Prefix(key)).Result()
}
func (rc *RedisConfig) setIVal(key string, value int) error {
	return rc.client.Set(ctx, Prefix(key), value, 0).Err()
}
func (rc *RedisConfig) getIVal(key string) (int, error) {
	return rc.client.Get(ctx, Prefix(key)).Int()
}

func Prefix(key string) string {
	return "AdBlockLists-" + key
}
