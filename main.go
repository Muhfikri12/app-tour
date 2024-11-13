package main

import (
	"fmt"
	"log"
	"net/http"
	"tour_destination/database"
	"tour_destination/handler"
	"tour_destination/repository"
	"tour_destination/service"
)

func main() {
	
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("error connecting database")
		return
	}

	defer db.Close()
	repo := repository.NewEventRepo(db)
	eventService := service.NewEventService(repo)
	eventHandler := handler.NewEventHandler(eventService)

	serverMux := http.NewServeMux()

	serverMux.HandleFunc("GET /event", eventHandler.EventHandler)
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8000", serverMux))
}