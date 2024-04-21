package main

import (
	"log"

	server "github.com/makiwebdeveloper/30dayschallenge-server"
	"github.com/makiwebdeveloper/30dayschallenge-server/pkg/handler"
	"github.com/makiwebdeveloper/30dayschallenge-server/pkg/repository"
	"github.com/makiwebdeveloper/30dayschallenge-server/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
