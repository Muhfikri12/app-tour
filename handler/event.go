package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tour_destination/model"
	"tour_destination/service"
	"tour_destination/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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
	

	events, totalData, totalPage, err := eh.service.GetEvent(page, date, sort)
	if err != nil {
		eh.service.Repo.Log.Error("event handler: ", zap.Error(err))
		utils.Response(w, http.StatusNotFound, "Data Not Found", nil)
		return
	}

	utils.SuccessWithPage(w, http.StatusOK, page, totalData, 6,totalPage, "Successfully", events)
}

func (eh *EventHandler) EventHandlerByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		eh.service.Repo.Log.Error("event handler: ", zap.Error(err))
		utils.Response(w, http.StatusInternalServerError, "invalid product id", nil)
		return
	}

	events, err := eh.service.GetEventByID(id)
	if err != nil {
		eh.service.Repo.Log.Error("event handler: ", zap.Error(err))
		utils.Response(w, http.StatusNotFound, fmt.Sprintf("id dengan %d tidak tersedia", id), nil)
		return
	}

	utils.Response(w, http.StatusOK, "Successfully Get Data", events)
	
}

func (eh *EventHandler) CreateHandlerTransaction(w http.ResponseWriter, r *http.Request) {
	trxs := model.Transaction{}
	validate := validator.New()

	if err := json.NewDecoder(r.Body).Decode(&trxs); err != nil {
		eh.service.Repo.Log.Error("event handler: ", zap.Error(err))
		utils.Response(w, http.StatusInternalServerError, "Error: " + err.Error(), nil)
		return
	}

	err := validate.Struct(trxs)
	if err != nil {
		errors := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("Field %s failed validation with tag %s", err.Field(), err.Tag()))
		}

		utils.Response(w, http.StatusUnprocessableEntity, "validation failed", errors)
		return
	}

	if err := eh.service.CreateBooking(&trxs); err != nil {
		eh.service.Repo.Log.Error("event handler: ", zap.Error(err))
		utils.Response(w, http.StatusUnprocessableEntity, "failed to create transactions", nil)
		return
	}
	
	utils.Response(w, http.StatusCreated, "Successfully", trxs)
}

func (eh *EventHandler) EventPlans(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		eh.service.Repo.Log.Error("event handler: id parameter missing")
		utils.Response(w, http.StatusBadRequest, "id parameter is required", nil)
		return
	}
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		eh.service.Repo.Log.Error("event handler: invalid id parameter", zap.Error(err))
		utils.Response(w, http.StatusBadRequest, "invalid id parameter", nil)
		return
	}

	plans, err := eh.service.GetEventPlanById(id)
	if err != nil {
		eh.service.Repo.Log.Error("event handler: ", zap.Error(err))
		utils.Response(w, http.StatusNotFound, fmt.Sprintf("id dengan %d tidak tersedia", id), nil)
		return
	}

	utils.Response(w, http.StatusOK, "Successfully Get Data", plans)
}

func (eh *EventHandler) EventLocations(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		eh.service.Repo.Log.Error("event handler: id parameter missing")
		utils.Response(w, http.StatusBadRequest, "id parameter is required", nil)
		return
	}
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		eh.service.Repo.Log.Error("event handler: invalid id parameter", zap.Error(err))
		utils.Response(w, http.StatusBadRequest, "invalid id parameter", nil)
		return
	}

	locations, err := eh.service.GetLocationById(id)
	if err != nil {
		eh.service.Repo.Log.Error("event handler: ", zap.Error(err))
		utils.Response(w, http.StatusNotFound, fmt.Sprintf("id dengan %d tidak tersedia", id), nil)
		return
	}

	utils.Response(w, http.StatusOK, "Successfully Get Data", locations)
}