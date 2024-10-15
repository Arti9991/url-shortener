package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shortener/internal/config"
	sl "url-shortener/internal/lib/sl/logger"
	mysql "url-shortener/internal/storage/MySQL"
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

	storage, err := mysql.NewStorage()
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	err = storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("failed to find", sl.Err(err))
	}
	err = storage.SaveURL("https://smgoogle.com", "smgoogle")
	if err != nil {
		log.Error("failed to find", sl.Err(err))
	}

	str, err := storage.GetURL("google")
	if err != nil {
		log.Error("failed to find", sl.Err(err))
	}
	fmt.Printf("\n%#v\n", str)

	// err = storage.DeleteURL("google")
	// if err != nil {
	// 	log.Error("failed to find", sl.Err(err))
	// }
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
