package model

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Event struct {
	Id    uint64    `json:"id, omitempty"`
	Title string    `json:"title, omitempty"`
	Date  time.Time `json:"date, omitempty"`
}

func (e *Event) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

func (e *Event) Parsing(data url.Values) error {
	title := data.Get("title")
	e.Title = title
	date := data.Get("date")
	datedate, err := time.Parse("2000-01-01", date)
	if err != nil {
		return fmt.Errorf("error parsing date")
	}
	e.Date = datedate
	if err != nil {

	}
	return nil
}
