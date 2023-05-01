package app

import (
	"net/http"
	"time"

	"github.com/hanoy/messenger/internal/handler"
	"github.com/hanoy/messenger/internal/repository"
	"github.com/hanoy/messenger/internal/service"
)



func Run() {
    repo := repository.NewUsersRepositoryPostgres()
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
