package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shortener/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//TO DO: init config: cleanenv

	conf := config.LoadConfig()

	fmt.Printf("%#v\n", conf)

	//TO DO: init logger: slog (just import log/slog)

	log := confLogger(conf.Env)

	log.Info("Starting url-shortner", slog.String("env", conf.Env))
	log.Debug("Debug messages are enabled")

	//TO DO: init stirage: sqlite

	//TO DO: init router: chi, "chi render"

	//TO DO: run server
}

func confLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)

	}
	return log

}
