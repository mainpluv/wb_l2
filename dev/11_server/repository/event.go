package repository

import (
	"dev/11_server/model"
	"time"
)

type Repository interface {
	Create(user_id uint64, e model.Event) (model.Event, error)
	Update(user_id uint64, e model.Event) error
	Delete(user_id uint64, event_id uint64) error
	GetForDay(user_id uint64, day time.Time) ([]model.Event, error)
	GetForWeek(user_id uint64, week time.Time) ([]model.Event, error)
	GetForMonth(user_id uint64, month time.Time) ([]model.Event, error)
}

type Data struct{}

func (d *Data) Create(user_id uint64, e model.Event) (model.Event, error) {
	event := model.Event{
		Id:    1,
		Title: "Mom's birthday",
		Date:  time.Time{},
	}
	return event, nil
}

func (d *Data) Update(user_id uint64, e model.Event) error   { return nil }
func (d *Data) Delete(user_id uint64, event_id uint64) error { return nil }
func (d *Data) GetForDay(user_id uint64, day time.Time) ([]model.Event, error) {
	events := []model.Event{
		{
			Id:    2,
			Title: "Dad's bithday",
			Date:  time.Time{},
		},
		{
			Id:    3,
			Title: "Bro's birthday",
			Date:  time.Time{},
		},
	}
	return events, nil
}
func (d *Data) GetForWeek(user_id uint64, week time.Time) ([]model.Event, error) {
	events := []model.Event{
		{
			Id:    2,
			Title: "Dad's bithday",
			Date:  time.Time{},
		},
		{
			Id:    3,
			Title: "Bro's birthday",
			Date:  time.Time{},
		},
		{
			Id:    4,
			Title: "Sister's birthday",
			Date:  time.Time{},
		},
	}
	return events, nil
}
func (d *Data) GetForMonth(user_id uint64, month time.Time) ([]model.Event, error) {
	events := []model.Event{
		{
			Id:    2,
			Title: "Dad's bithday",
			Date:  time.Time{},
		},
		{
			Id:    3,
			Title: "Bro's birthday",
			Date:  time.Time{},
		},
		{
			Id:    4,
			Title: "Sister's birthday",
			Date:  time.Time{},
		},
		{
			Id:    5,
			Title: "Grandma's birthday",
			Date:  time.Time{},
		},
	}
	return events, nil
}
