package handler

import (
	"dev/11_server/middleware"
	"dev/11_server/model"
	"dev/11_server/service"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// создаем структуры и описываем основные методы
type EventHandler struct {
	eventService *service.EventService
}

func NewEventHandler(eventService *service.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

type Response struct {
	Res any    `json:"result, omitempty"`
	Err string `json:"error, omitempty"`
}

func (h *EventHandler) InitHandler() error {
	http.HandleFunc("/create_event", middleware.Logging(h.CreateEvent))
	http.HandleFunc("/update_event", middleware.Logging(h.UpdateEvent))
	http.HandleFunc("/delete_event", middleware.Logging(h.DeleteEvent))
	http.HandleFunc("/events_for_day", middleware.Logging(h.GetEventsForDay))
	http.HandleFunc("/events_for_week", middleware.Logging(h.GetEventsForWeek))
	http.HandleFunc("/events_for_month", middleware.Logging(h.GetEventsForMonth))

	return http.ListenAndServe(":8080", nil)
}

func ResponseError(w http.ResponseWriter, errStr string, code int) {
	w.WriteHeader(code)
	response := Response{Err: errStr}
	json.NewEncoder(w).Encode(response)
}

func ResponseRes(w http.ResponseWriter, resp any) {
	response := Response{Res: resp}
	json.NewEncoder(w).Encode(response)
}

func GetId(data url.Values, field string) (uint64, error) {
	id := data.Get(field)
	if id == "" {
		return 0, fmt.Errorf("no id")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("invalid id")
	}
	return uint64(idInt), nil
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	userID, err := GetId(r.URL.Query(), "user_id")
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var event model.Event
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdEvent, err := h.eventService.CreateEvent(userID, event)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ResponseRes(w, createdEvent)
}

func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	userID, err := GetId(r.URL.Query(), "user_id")
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var event model.Event
	err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.eventService.UpdateEvent(userID, event)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseRes(w, map[string]string{"result": "Event updated successfully"})
}

func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	userID, err := GetId(r.URL.Query(), "user_id")
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	eventID, err := GetId(r.URL.Query(), "event_id")
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.eventService.DeleteEvent(userID, eventID)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseRes(w, map[string]string{"result": "Event deleted successfully"})
}

func (h *EventHandler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	userID, err := GetId(r.URL.Query(), "user_id")
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	date := r.URL.Query().Get("date")
	day, err := time.Parse("2006-01-02", date)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.eventService.GetEventsForDay(userID, day)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseRes(w, events)
}

func (h *EventHandler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	userID, err := GetId(r.URL.Query(), "user_id")
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	date := r.URL.Query().Get("week")
	week, err := time.Parse("2006-01-02", date)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.eventService.GetEventsForWeek(userID, week)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseRes(w, events)
}

func (h *EventHandler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	userID, err := GetId(r.URL.Query(), "user_id")
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	date := r.URL.Query().Get("month")
	month, err := time.Parse("2006-01", date)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	events, err := h.eventService.GetEventsForMonth(userID, month)
	if err != nil {
		ResponseError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseRes(w, events)
}
