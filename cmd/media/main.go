package main

import (
	"log"

	"github.com/anuar45/media/internal/api"
	"github.com/anuar45/media/internal/service/media"
)

func main() {
	config, err := media.ConfigFromFlags()
	if err != nil {
		log.Fatalln(err)
	}

	svc := media.NewService(config)
	apiServer := api.NewServer(svc, config)

	log.Println("Starting media server on port", config.ListenAddr)
	log.Fatalln(apiServer.Run())
}
