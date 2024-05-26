package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	api_tansaksi "transaksi-service/api/api_transaksi"
	account_service "transaksi-service/service/transaksi"
	postgres_store "transaksi-service/store/postgres_store/store"
	"transaksi-service/utils/config"
	"transaksi-service/utils/errs"

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

	// init data access layer
	postgresStore := postgres_store.NewPostgresStore(logger, conn)

	// init accountService layer
	transaksiService := account_service.NewService(config, logger, postgresStore)

	// init presentation layer
	apiTransaksi := api_tansaksi.NewApi(transaksiService)

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

	app.Post("/tabung", apiTransaksi.Tabung)
	app.Post("/tarik", apiTransaksi.Tarik)

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
