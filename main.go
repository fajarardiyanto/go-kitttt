package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"lat-gokit/users"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	DbHost = "localhost"
	DbPort = "5432"
	DbUser = "frontend-hq"
	DbName = "gokitexample"
	DbPassword = "1qaz2wsx"
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

	var db *gorm.DB
	DB_URI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	{
		var err error

		db, err = gorm.Open("postgres", DB_URI)
		if err != nil {
			fmt.Println("qqqq", err.Error())
			level.Error(logger).Log("Exit", err.Error())
			os.Exit(-1)
		}
	}

	flag.Parse()
	ctx := context.Background()
	var srv users.Service
	{
		repository := users.NewRepo(db, logger)

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