package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	config2 "github.com/sviatilnik/gophermart/internal/infrastructure/config"
	"github.com/sviatilnik/gophermart/internal/infrastructure/http/handlers"
	middleware2 "github.com/sviatilnik/gophermart/internal/infrastructure/http/middleware"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// TODO Сделать отключение приложения красивее
	conf := getConfig()

	db, er := sql.Open("pgx", conf.DatabaseDSN)
	if er != nil {
		panic(er)
	}

	//#region Миграции

	// TODO Перенести миграции
	driver, er := postgres.WithInstance(db, &postgres.Config{})
	if er != nil {
		panic(er)
	}

	// https://github.com/golang-migrate/migrate/blob/master/source/file/README.md
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/infrastructure/migrations",
		"postgres", driver)

	if err != nil {
		panic(err)
	}
	m.Up()
	//#endregion Миграции

	r := chi.NewRouter()
	r.Use(middleware2.GZIPCompress)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	logger := getLogger()

	r.Get("/", handlers.GetUser())

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:    conf.Host,
		Handler: r,
	}

	go func() {
		logger.Info(fmt.Sprintf("start server on %s", server.Addr))

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(err.Error())
		}
	}()

	<-quitChan

	logger.Info("shutting down server...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Fatal(err.Error())
	}

}

func getConfig() config2.Config {
	return config2.NewConfig(
		config2.NewDefaultProvider(),
		config2.NewFlagProvider(),
		config2.NewEnvProvider(config2.NewOSEnvGetter()),
	)
}

func getLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return logger.Sugar()
}
