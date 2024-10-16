package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/http-server/handlers/url/save"
	"url-shortener/internal/lib/logger/sl"

	//mwLogger "url-shortener/internal/http-server/middleware/logger"

	mysql "url-shortener/internal/storage/MySQL"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	log.Error("Error messages are enabled")

	//TO DO: init stirage: sqlite

	storage, err := mysql.NewStorage()
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	// err = storage.SaveURL("https://google.com", "google")
	// if err != nil {
	// 	log.Error("failed to find", sl.Err(err))
	// }
	// err = storage.SaveURL("https://smgoogle.com", "smgoogle")
	// if err != nil {
	// 	log.Error("failed to find", sl.Err(err))
	// }

	// str, err := storage.GetURL("google")
	// if err != nil {
	// 	log.Error("failed to find", sl.Err(err))
	// }
	// fmt.Printf("\n%#v\n", str)
	// err = storage.DeleteURL("google")
	// if err != nil {
	// 	log.Error("failed to find", sl.Err(err))
	// }
	//TO DO: init router: chi, "chi render"

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	//router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(log, storage))

	log.Info("starting server", slog.String("adress", conf.Adress))

	srv := &http.Server{
		Addr:         conf.Adress,
		Handler:      router,
		ReadTimeout:  conf.HTTPServer.Timeout,
		WriteTimeout: conf.HTTPServer.Timeout,
		IdleTimeout:  conf.HTTPServer.IdleTimeout,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
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
