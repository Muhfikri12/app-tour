package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tour_destination/service"
)

type EventHandler struct {
	service *service.EventService
}

func NewEventHandler(service *service.EventService) *EventHandler {
	return &EventHandler{service}
}

func (eh *EventHandler) EventHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")

	sort := r.URL.Query().Get("sort")
	date := r.URL.Query().Get("date")

	page, _ := strconv.Atoi(pageStr)
	

	if page < 1 {
		page = 1
	}
	

	events, err := eh.service.GetEvent(page, date, sort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
