package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RexterR/imger/cache"
	"github.com/RexterR/imger/imger"
	"github.com/RexterR/imger/implementation/image"
	"github.com/RexterR/imger/implementation/profile"
	"github.com/RexterR/imger/pkg/config"
	"github.com/RexterR/imger/pkg/log"
	httprepository "github.com/RexterR/imger/repository/http"
	"github.com/RexterR/imger/repository/memory"
	"github.com/RexterR/imger/repository/mongo"
	"github.com/RexterR/imger/repository/redis"
	httpapi "github.com/RexterR/imger/transport/http"
	"github.com/talento90/go-health"
)

func mongoClient() (*mongo.Client, error) {
	c, err := config.GetMongoConfiguration()

	if err != nil {
		return nil, err
	}

	return mongo.NewClient(c)
}

func redisClient() (*redis.Client, error) {
	c, err := config.GetRedisConfiguration()

	if err != nil {
		return nil, err
	}

	return redis.NewClient(c)
}

func httpServer(l log.Logger, rc *redis.Client, ms *mongo.Client, h health.Health) *http.Server {
	var imgService imger.ImageService
	{
		redisCache := redis.New(rc)
		imgRepository := httprepository.NewImageRepository()
		effectRepo := memory.NewImageRepository(imgRepository)
		imgCache := cache.NewImage(redisCache)

		imgService = image.NewService(imgRepository, effectRepo)
		imgService = image.NewCacheService(imgCache, imgService)
		imgService = image.NewLogService(l, imgService)
	}

	var profileService imger.ProfileService
	{
		profileRepository := mongo.NewProfileRepository(ms)
		profileService = profile.NewService(profileRepository)
		profileService = profile.NewLogService(l, profileService)
	}

	h.RegisterChecker("redis", rc)
	h.RegisterChecker("mongo", ms)

	serverDeps := &httpapi.ServerDependencies{
		Logger:         l,
		ImgService:     imgService,
		ProfileService: profileService,
		Health:         h,
	}

	serverConfig, err := config.GetServerConfiguration()

	if err != nil {
		l.Panic(err)
	}

	srv := httpapi.NewServer(&serverConfig, serverDeps)

	return &srv
}

func main() {
	var gracefulShutdown = make(chan os.Signal)

	signal.Notify(gracefulShutdown, syscall.SIGTERM)
	signal.Notify(gracefulShutdown, syscall.SIGINT)

	logConfig, err := config.GetLogConfiguration()

	if err != nil {
		panic(err)
	}

	logger, err := log.NewLogger(logConfig)

	if err != nil {
		panic(err)
	}

	redisClient, err := redisClient()

	if err != nil {
		logger.Panic("Error connecting to Redis", err)
	}

	mongoClient, err := mongoClient()

	if err != nil {
		logger.Panic("Error connecting to Mongo", err)
	}

	h := health.New("imgart", health.Options{CheckersTimeout: time.Second})

	server := httpServer(logger, redisClient, mongoClient, h)

	go func() {
		<-gracefulShutdown
		exitCode := 0
		h.Shutdown()

		logger.Info("Starting graceful shutdown")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		err := server.Shutdown(ctx)

		if err != nil {
			exitCode = 1
			logger.Error("Error closing server:", err)
		}

		if err := mongoClient.Disconnect(ctx); err != nil {
			exitCode = 1
			logger.Error("Error closing mongo:", err)
		}

		if err := redisClient.Close(); err != nil {
			exitCode = 1
			logger.Error("Error closing redis:", err)
		}

		logger.Info("Graceful shutdown completed")

		os.Exit(exitCode)
	}()

	logger.Info("Server listening on port: ", server.Addr)

	if err := http.ListenAndServe(":"+server.Addr, server.Handler); err != nil {
		logger.Error("Error starting server", err)
	}
}
