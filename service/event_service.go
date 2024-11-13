package service

import (
	"sort"
	"time"
	"tour_destination/model"
	"tour_destination/repository"
)

type EventService struct {
	repo *repository.EventRepoDB
}

func NewEventService(repo *repository.EventRepoDB) *EventService {
	return &EventService{repo}
}

func (s *EventService) GetEvent(page int, date, sortParam string) (*[]model.Events, error) {
	limit := 10

	events, err := s.repo.GetEvent(page, limit, date)
	if err != nil {
		return nil, err
	}

	today := time.Now().Format("2006-01-02")

	filteredEvents := []model.Events{}
	for _, event := range *events {
		if event.Date >= today {
			filteredEvents = append(filteredEvents, event)
		}
	}

	switch sortParam {
	case "highToLow":
		sort.Slice(filteredEvents, func(i, j int) bool {
			return filteredEvents[i].Price > filteredEvents[j].Price
		})
	case "lowToHigh":
		sort.Slice(filteredEvents, func(i, j int) bool {
			return filteredEvents[i].Price < filteredEvents[j].Price
		})
	}

	return &filteredEvents, nil
}
