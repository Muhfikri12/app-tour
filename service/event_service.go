package service

import (
	"math"
	"sort"
	"time"
	"tour_destination/model"
	"tour_destination/repository"

	"go.uber.org/zap"
)

type EventService struct {
	Repo *repository.EventRepoDB
	
}

func NewEventService(Repo *repository.EventRepoDB) *EventService {
	return &EventService{Repo}
}

func (s *EventService) GetEvent(page int, date, sortParam string) (*[]model.Events, int, int, error) {
	limit := 6

	var totalData int

	events, totalData, err := s.Repo.GetEvent(page, limit, date)

	if err != nil {
		s.Repo.Log.Error("event service: ", zap.Error(err))
		return nil, 0, 0, err
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

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

	
	return &filteredEvents, totalData, totalPage, nil
}

func (s *EventService) GetEventByID(id int) (*model.Events, error) {
	events, err := s.Repo.GetEventByID(id)
	if err != nil {
		s.Repo.Log.Error("event service GetEventByID: ", zap.Error(err))
		return nil, err
	}

	return events, nil
}

func (s *EventService) CreateBooking(trx *model.Transaction) error {
	
	if err := s.Repo.CreateTransaction(trx); err != nil {
		s.Repo.Log.Error("event service CreateBooking: ", zap.Error(err))
		return err
	}

	return nil
}

func (s *EventService) GetEventPlanById(id int) (*[]model.EventPlan, error) {
	plans, err := s.Repo.GetEventPlanById(id)
	if err != nil {
		s.Repo.Log.Error("event service : ", zap.Error(err))
		return nil, err
	}

	return plans, nil
}

func (s *EventService) GetLocationById(id int) (*[]model.Location, error) {
	locations, err := s.Repo.GetEventLocationById(id)
	if err != nil {
		s.Repo.Log.Error("event service : ", zap.Error(err))
		return nil, err
	}

	return locations, nil
}