package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/hanoy/messenger/internal/config"
	"github.com/hanoy/messenger/internal/handler"
	"github.com/hanoy/messenger/internal/repository"
	"github.com/hanoy/messenger/internal/service"
	"github.com/hanoy/messenger/pkg/auth"
	"github.com/hanoy/messenger/pkg/db/postgres"
)

func Run(configPath string) {
	config, err := config.GetConfig(configPath)
    if err != nil {
        log.Fatalf("cannot load config: %v", err)
    }

	dbpool, err := postgres.CreateConnectionPool(context.Background(), config.DB.URL)
    if err != nil {
		log.Fatalf("unable to establish connection with database: %v", err)
    }

    repo := repository.NewRepositories(dbpool)
	services := service.NewServices(repo)
    tokenProvider := auth.NewProvider(config)
	handler := handler.NewHandler(services, tokenProvider)

	server := http.Server{
		Handler:      handler.Init(),
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}

    log.Printf("Starting server at: %v\n", server.Addr)
    if err := server.ListenAndServe(); err != nil {
        log.Fatalf("error while listening: %v", err)
    }
}
