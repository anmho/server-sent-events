package main

import (
	"errors"
	"fmt"
	"github.com/anmho/prism/api"
	"github.com/anmho/prism/scope"
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"log/slog"
	"net/http"
)

const (
	port = 8080
)

type Config struct {
	OpenAIKey string `env:"OPENAI_KEY"`
}

func main() {
	var config Config

	err := env.Parse(&config)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("config: %+v", config)

	mux := api.MakeServer()

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	scope.GetLogger().Info("server is listening", slog.Int("port", port))
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}
