package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/sav1nbrave4code/APG3/configs"
	v1 "github.com/sav1nbrave4code/APG3/internal/controller/http/v1"
	"github.com/sav1nbrave4code/APG3/internal/repository/postgres"
	"github.com/sav1nbrave4code/APG3/internal/service/data_manager"
	"github.com/sav1nbrave4code/APG3/internal/service/function_manager"
	"github.com/sav1nbrave4code/APG3/internal/service/table_manager"
	"github.com/sav1nbrave4code/APG3/pkg/db/postgres_db"
	"github.com/sav1nbrave4code/APG3/pkg/http_server"
	"github.com/sav1nbrave4code/APG3/pkg/logger"
)

func main() {
	logger.Init()
	s := gocron.NewScheduler(time.UTC)
	s.Cron("0 0 * * *").Do(logger.Init)
	s.StartAsync()

	cfg, err := configs.New()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	log.Info().Msg("Completed read configs")

	postgresDb, err := postgres_db.NewPgx(postgres_db.DSN(cfg.Postgres.DSN))
	if err != nil {
		log.Fatal().Msgf("PostgresDb error: %s", err)
	}
	log.Info().Msg("Completed init postgresDb")

	defer postgresDb.Close()

	postgresRepo := postgres.New(postgresDb)

	tableManagerService := table_manager.New(postgresRepo)
	dataManagerService := data_manager.New(postgresRepo)
	functionManagerService := function_manager.New(postgresRepo)

	mux := chi.NewRouter()

	handler := v1.New(mux, tableManagerService, dataManagerService, functionManagerService)
	handler.Run()

	httpServer := http_server.New(mux, http_server.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info().Msgf("Signal: %s", s.String())
	case err = <-httpServer.Notify():
		log.Error().Msgf("HttpServer.Notify: %v", err)
	}

	if err := httpServer.Shutdown(); err != nil {
		log.Error().Msgf("HttpServer.Shutdown: %v", err)
	}
}
