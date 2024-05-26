package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	request_processor "mutasi-service/message_hub"
	postgres_store "mutasi-service/store/postgres_store/store"
	"mutasi-service/store/redis_store"
	"mutasi-service/utils/config"
	"mutasi-service/utils/errs"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func start() {
	const op errs.Op = "main/start"

	// init logger
	var logger = logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	logger.Formatter = new(logrus.TextFormatter)                     //default
	logger.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	logger.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	logger.Level = logrus.DebugLevel
	logger.Out = os.Stdout

	// load environment variables from .env file
	config, err := config.LoadConfig(".")
	if err != nil {
		logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "LoadConfig",
			"err":   err.Error(),
		}).Error("failed to read config file!")

		os.Exit(1)
	}

	// create db connection
	conn, err := sql.Open(config.PostgresDriver, config.PostgresUrl)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "Open",
			"err":   err.Error(),
		}).Error("failed to connect to the db!")

		os.Exit(1)
	}

	// create redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisServiceAddress,
		Password: config.RedisPassword,
		DB:       0,
	})

	// init data access layer
	postgresStore := postgres_store.NewPostgresStore(logger, conn)
	redisStore := redis_store.NewRedisStore(logger, redisClient)

	// init presentation layer
	redisConsumer := request_processor.NewRequestProcessor(config, logger, postgresStore, redisStore)

	// run request processor
	go redisConsumer.Run(context.Background())

	// create error channel
	errorChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errorChan <- fmt.Errorf("%s", <-c)
	}()

	logger.WithFields(logrus.Fields{
		"op":    op,
		"scope": "Open",
		"err":   (<-errorChan).Error(),
	}).Error("failed to connect to the db!")
}
