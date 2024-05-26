package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"transaksi-service/api/api_auth"
	api_tansaksi "transaksi-service/api/api_transaksi"
	"transaksi-service/service/auth_service"
	transaksi "transaksi-service/service/transaksi"
	postgres_store "transaksi-service/store/postgres_store/store"
	"transaksi-service/store/redis_store"
	"transaksi-service/utils/config"
	"transaksi-service/utils/errs"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// read args
	host := "0.0.0.0"
	port := "3000"
	if flag.NArg() >= 2 {
		host = flag.Arg(1)
		port = flag.Arg(2)
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

	// init accountService layer
	transaksiService := transaksi.NewService(config, logger, postgresStore, redisStore)
	authService := auth_service.NewService(config, logger, postgresStore)

	// init presentation layer
	apiTransaksi := api_tansaksi.NewApi(transaksiService)
	apiAuth := api_auth.NewApi(authService)

	// init fiber app
	app := fiber.New()

	// CORS middleware configuration
	corsConfig := cors.Config{
		AllowOrigins: "http://0.0.0.0:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}

	app.Use(cors.New(corsConfig))

	// endpoints
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("PONG")
	})

	app.Post("/tabung", apiAuth.Tabung, apiTransaksi.Tabung)
	app.Post("/tarik", apiAuth.Tarik, apiTransaksi.Tarik)

	// start the server
	err = app.Listen(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "Listen",
			"err":   err.Error(),
		}).Error(fmt.Sprintf("failed to listen at%s:%s", host, port))

		os.Exit(1)
	}
}
