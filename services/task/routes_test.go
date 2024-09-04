package task

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/victorluisca/go-todo-app/types"
)

func TestTaskServiceHandlers(t *testing.T) {
	t.Run("get task by id", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/task/", nil)
		req.SetPathValue("taskID", "1")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handleTask)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("create new task", func(t *testing.T) {
		task := types.Task{
			ID:        2,
			Title:     "Tarefa 2",
			Priority:  "Low",
			CreatedAt: time.Now(),
		}
		marshalled, _ := json.Marshal(task)

		req, err := http.NewRequest("POST", "/tasks", bytes.NewReader(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handleTasks)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})

	t.Run("invalid task id", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/task/", nil)
		req.SetPathValue("taskID", "invalid")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handleTask)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("invalid payload", func(t *testing.T) {
		task := types.Task{
			ID:        2,
			Title:     "In",
			Priority:  "Invalid",
			CreatedAt: time.Now(),
		}
		marshalled, _ := json.Marshal(task)

		req, err := http.NewRequest("POST", "/tasks", bytes.NewReader(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handleTasks)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("method not allowed", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/tasks", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handleTasks)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("expected status code %d, got %d", http.StatusMethodNotAllowed, rr.Code)
		}
	})
}
