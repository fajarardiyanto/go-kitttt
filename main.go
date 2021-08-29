package main

import (
	"context"
	"flag"
	"fmt"
	"lat-gokit/users"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	httpAddr := flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "users",
			"time: ", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	err := godotenv.Load()
	if err != nil {
		level.Error(logger).Log("Error getting env, not comming through ", err)
	}

	DB_URI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DbHost"), os.Getenv("DbPort"), os.Getenv("DbUser"), os.Getenv("DbName"), os.Getenv("DbPassword"))
	var db *gorm.DB
	{
		var err error

		db, err = gorm.Open("postgres", DB_URI)
		if err != nil {
			level.Error(logger).Log("Exit", err.Error())
			os.Exit(-1)
		}
	}

	flag.Parse()
	ctx := context.Background()
	var srv users.Service
	{
		repository := users.NewRepo(db, logger)
		users.LoadSeed(db)

		srv = users.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := users.MakeEndpoints(srv)

	go func() {
		fmt.Println("listing on port", *httpAddr)
		handler := users.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
