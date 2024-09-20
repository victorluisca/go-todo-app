package task

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/victorluisca/go-todo-app/types"
)

func TestTaskServiceHandlers(t *testing.T) {
	store := &mockTaskStore{}

	t.Run("get all tasks", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handleTasks(w, r, store) })
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("create task", func(t *testing.T) {
		task := `{"title": "Tarefa 1", "priority": "High"}`
		req, err := http.NewRequest(http.MethodPost, "/tasks", strings.NewReader(task))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handleTasks(w, r, store) })
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})

	t.Run("get task by id", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/task/", nil)
		req.SetPathValue("taskID", "1")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handleTask(w, r, store) })
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("update task", func(t *testing.T) {
		task := `{"title": "Tarefa 1", "priority": "Medium"}`
		req, err := http.NewRequest(http.MethodPut, "/task/", strings.NewReader(task))
		req.SetPathValue("taskID", "1")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handleTask(w, r, store) })
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})

	t.Run("delete task", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/task/", nil)
		req.SetPathValue("taskID", "1")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handleTask(w, r, store) })
		handler.ServeHTTP(rr, req)

		assertStatus(t, rr.Code, http.StatusNoContent)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected status code %d, got %d", want, got)
	}
}

type mockTaskStore struct{}

func (m *mockTaskStore) GetAllTasks() ([]*types.Task, error) {
	return []*types.Task{}, nil
}

func (m *mockTaskStore) CreateTask(task types.Task) error {
	return nil
}

func (m *mockTaskStore) GetTaskByID(id int) (*types.Task, error) {
	return &types.Task{}, nil
}

func (m *mockTaskStore) UpdateTask(task types.Task) error {
	return nil
}

func (m *mockTaskStore) DeleteTask(id int) error {
	return nil
}
