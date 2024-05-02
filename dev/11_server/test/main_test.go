package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"dev/11_server/handler"
	"dev/11_server/model"
	"dev/11_server/repository"
	"dev/11_server/service"
)

func TestCreateEvent(t *testing.T) {
	// инициализируем мок репозиторий и мок сервис
	mockRepo := &repository.Data{}
	mockService := service.NewEventService(mockRepo)
	eventHandler := handler.NewEventHandler(mockService)
	event := model.Event{
		Id:    1,
		Title: "Mom's birthday",
		Date:  time.Time{},
	}
	// сериаллизуем событие
	eventJSON, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("Error marshaling event JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "/create_event?user_id=1", bytes.NewBuffer(eventJSON))
	if err != nil {
		t.Fatalf("Error creating HTTP request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(eventHandler.CreateEvent)
	handler.ServeHTTP(rr, req)
	// проверяем код состояния ответа
	if status := rr.Code; status != http.StatusCreated && status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
	// сравниваем ответ
	expected := `{"result":{"id":1,"title":"Mom's birthday","date":"0001-01-01T00:00:00Z"},"error":""}`
	actual := strings.TrimSpace(rr.Body.String())
	if actual != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", actual, expected)
	}
}
