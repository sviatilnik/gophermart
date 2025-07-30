package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sviatilnik/gophermart/internal/config"
	"go.uber.org/zap"
	"net/http"
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

	err := http.ListenAndServe(conf.Host, r)
	if err != nil {
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
