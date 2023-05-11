package app

import (
	"context"
	"net/http"
	"time"

	"github.com/hanoy/messenger/internal/config"
	"github.com/hanoy/messenger/internal/handler"
	"github.com/hanoy/messenger/internal/repository"
	"github.com/hanoy/messenger/internal/service"
	"github.com/hanoy/messenger/pkg/db/postgres"
)

func Run() {
	config := config.GetConfig()
	dbpool := postgres.CreateConnectionPool(context.TODO(), config.DB.URL)
    repo := repository.NewRepositories(dbpool)
	services := service.NewServices(repo)
	handler := handler.NewHandler(services)

	server := http.Server{
		Handler:      handler.Init(),
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}

	server.ListenAndServe()
}
