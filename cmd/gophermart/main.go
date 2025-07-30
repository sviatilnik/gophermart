package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sviatilnik/gophermart/internal/config"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	conf := getConfig()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	logger := getLogger()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.RequestURI)
	})

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

func getConfig() config.Config {
	return config.NewConfig(
		config.NewDefaultProvider(),
		config.NewFlagProvider(),
		config.NewEnvProvider(config.NewOSEnvGetter()),
	)
}

func getLogger() *zap.SugaredLogger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return logger.Sugar()
}
