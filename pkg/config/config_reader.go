package config

import (
	"os"
	"time"

	"github.com/RexterR/imger/pkg/log"
	"github.com/RexterR/imger/repository/mongo"
	"github.com/RexterR/imger/repository/redis"
	"github.com/RexterR/imger/transport/http"
	goredis "github.com/go-redis/redis"
)

// GetLogConfiguration get logger configurations
func GetLogConfiguration() (log.Configuration, error) {
	config := log.Configuration{
		Level:  getEnv("LOG_LEVEL", "debug"),
		Output: os.Stdout,
	}

	return config, config.Validate()
}

// GetServerConfiguration get server configurations
func GetServerConfiguration() (http.Configuration, error) {
	config := http.Configuration{
		Address:      getEnv("PORT", "4005"),
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 4,
	}

	return config, config.Validate()
}

// GetMongoConfiguration returns the mongo configuration
func GetMongoConfiguration() (mongo.Configuration, error) {
	config := mongo.Configuration{
		Database: getEnv("MONGO_DATABASE", "imger"),
		MongoURL: getEnv("MONGO_URL", "mongodb://localhost:27017"),
	}

	return config, config.Validate()
}

// GetRedisConfiguration returns the redis configuration
func GetRedisConfiguration() (redis.Configuration, error) {
	db, err := goredis.ParseURL(getEnv("REDIS_URL", "redis://localhost:6379"))

	if err != nil {
		return redis.Configuration{}, err
	}

	config := redis.Configuration{
		Address:  db.Addr,
		Password: db.Password,
		Database: db.DB,
	}

	return config, config.Validate()
}

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
