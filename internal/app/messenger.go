package app

import (
	"context"
	"fmt"
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
	dbpool := postgres.CreateConnectionPool(context.TODO(),
    "postgres://"+config.DB.User+":"+config.DB.Password+"@"+config.DB.Host+":"+config.DB.Port+"/"+config.DB.DBName)
	repo := repository.NewUsersRepositoryPostgres(dbpool)
	service := service.NewUsersService(repo)
	handler := handler.NewHandler(service)

	server := http.Server{
		Handler:      handler.Init(),
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}

	server.ListenAndServe()
}
