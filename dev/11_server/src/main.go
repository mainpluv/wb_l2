package main

import (
	"dev/11_server/handler"
	"dev/11_server/repository"
	"dev/11_server/service"
	"log"
)

func main() {
	repository := repository.Data{}
	eventService := service.NewEventService(&repository)
	eventHandler := handler.NewEventHandler(eventService)
	err := eventHandler.InitHandler()
	if err != nil {
		log.Fatalf("error initializing server", err)
	} else {
		log.Println("server startted")
	}
}
