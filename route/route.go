package route

import (
	"database/sql"
	"fmt"
	"tour_destination/database"
	"tour_destination/handler"
	"tour_destination/library"
	"tour_destination/middleware"
	"tour_destination/repository"
	"tour_destination/service"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func RouteInit() (*sql.DB, *chi.Mux, *zap.Logger) {
	// Inisialisasi Router
	r := chi.NewRouter()

	db, err := database.InitDB()
	if err != nil {
		fmt.Println("Error saat inisialisasi database:", err)
		return nil, nil, nil
	}

	logger := library.InitLog()

	eventRepo := repository.NewEventRepo(db, logger)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	mw := middleware.NewMiddleware(logger)

	r.Route("/api", func(api chi.Router) {

		api.Use(mw.MiddlewareLogger)

		api.Route("/events", func(eventRoute chi.Router) {
			eventRoute.Get("/", eventHandler.EventHandler)      
			eventRoute.Get("/{id}", eventHandler.EventHandlerByID) 
			eventRoute.Post("/booking", eventHandler.CreateHandlerTransaction) 
		})
	})

	return db, r, logger
}
