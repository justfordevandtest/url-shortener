package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	// Shortener config
	BaseURL        string `env:"BASE_URL" envDefault:"http://localhost:8080/api/v1"`
	CacheThreshold int    `env:"CACHE_THRESHOLD" envDefault:"10"`

	// Blacklist config
	Blacklist string `env:"BLACKLIST" envDefault:"bad(.+),invalid(.+)"`

	// JWT config
	JWTRealm  string `env:"JWT_REALM" envDefault:"rabbit finance test"`
	JWTSecret string `env:"JWT_SECRET" envDefault:"secret"`

	// Redis config
	RedisCacheAddr    string `env:"REDIS_CACHE_ADDR" envDefault:"localhost:6379"`

	// MongoDB config
	MongoDBEndpoint    string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName        string `env:"MONGODB_NAME" envDefault:"ShortURL"`
	MongoDBURLCollName string `env:"MONGODB_URL_COLLECTION_NAME" envDefault:"url"`
}

func Get() *Config {
	appConfig := &Config{}
	_ = env.Parse(appConfig)
	return appConfig
}
