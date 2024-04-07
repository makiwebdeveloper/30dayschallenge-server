package main

import (
	"log"

	server "github.com/makiwebdeveloper/30dayschallenge-server"
	"github.com/makiwebdeveloper/30dayschallenge-server/pkg/handler"
)

func main() {
	handler := handler.NewHandler()

	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
