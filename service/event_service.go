package service

import (
	"tour_destination/model"
	"tour_destination/repository"
)

type EventService struct {
	repo *repository.EventRepoDB
}

func NewEventService(repo *repository.EventRepoDB) *EventService {
	return &EventService{repo}
}

func (es *EventService) GetEventData(page int, sort, date string) (*[]model.Events, error) {
	limit := 10
	
	if sort == "" {
		sort = "ASC" 
	} else if sort == "highToLow" {
		sort = "DESC"
	} else if sort == "lowToHigh" {
		sort = "ASC"
	}

	events, err := es.repo.GetEvent(page, limit, date, sort)
	if err != nil {
		return nil, err
	}

	return events, nil
}


