package config

import (
	"adblocklists/redisClient"
	"strings"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/env"
	"github.com/kwitsch/go-dockerutils/net"
)

type Config struct {
	Resolver *net.DockerResolver      `koanf:"resolver"`
	Redis    *redisClient.RedisConfig `koanf:"redis"`
	Refresh  time.Duration            `koanf:"refresh"`
	Block    List                     `koanf:"block"`
	Allow    List                     `koanf:"allow"`
}

type List struct {
	Lists   map[int]string `koanf:"lists"`
	Entries map[int]string `koanf:"entries"`
}

const prefix = "ABL_"

func Get() *Config {
	var k = koanf.New(".")
	k.Load(env.Provider(prefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, prefix)), "_", ".", -1)
	}), nil)

	var res Config
	k.UnmarshalWithConf("", &res, koanf.UnmarshalConf{Tag: "koanf"})
	return &res
}
