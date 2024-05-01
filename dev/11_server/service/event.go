package service

import (
	"dev/11_server/model"
	"dev/11_server/repository"
	"time"
)

type EventServiceInterface interface {
	Create(user_id uint64, e model.Event) (*model.Event, error)
	Update(user_id uint64, e model.Event) error
	Delete(user_id uint64, event_id uint64) error
	AllForDay(user_id uint64, day time.Time) ([]model.Event, error)
	AllForWeek(user_id uint64, day time.Time) ([]model.Event, error)
	AllForMonth(user_id uint64, day time.Time) ([]model.Event, error)
}

type EventService struct {
	Repo repository.Repository
}

func NewEventService(repo repository.Repository) *EventService {
	return &EventService{Repo: repo}
}

func (s *EventService) CreateEvent(userID uint64, event model.Event) (model.Event, error) {
	createdEvent, err := s.Repo.Create(userID, event)
	if err != nil {
		return model.Event{}, err
	}
	return createdEvent, nil
}

func (s *EventService) UpdateEvent(userID uint64, event model.Event) error {
	err := s.Repo.Update(userID, event)
	if err != nil {
		return err
	}
	return nil
}

func (s *EventService) DeleteEvent(userID, eventID uint64) error {
	err := s.Repo.Delete(userID, eventID)
	if err != nil {
		return err
	}
	return nil
}

func (s *EventService) GetEventsForDay(userID uint64, day time.Time) ([]model.Event, error) {
	events, err := s.Repo.GetForDay(userID, day)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) GetEventsForWeek(userID uint64, week time.Time) ([]model.Event, error) {
	events, err := s.Repo.GetForWeek(userID, week)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) GetEventsForMonth(userID uint64, month time.Time) ([]model.Event, error) {
	events, err := s.Repo.GetForMonth(userID, month)
	if err != nil {
		return nil, err
	}
	return events, nil
}
