package task

import (
	"net/http"
	"strconv"

	"github.com/victorluisca/go-todo-app/types"
	"github.com/victorluisca/go-todo-app/utils"
)

func RegisterRoutes(router *http.ServeMux, store *Store) {
	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) { handleTasks(w, r, store) })
	router.HandleFunc("/task/{taskID}", func(w http.ResponseWriter, r *http.Request) { handleTask(w, r, store) })
}

func handleTasks(w http.ResponseWriter, r *http.Request, store *Store) {
	switch r.Method {
	case "GET":
		getAllTasks(w, store)
	case "POST":
		createTask(w, r, store)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAllTasks(w http.ResponseWriter, store *Store) {
	tasks, err := store.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createTask(w http.ResponseWriter, r *http.Request, store *Store) {
	var task types.Task
	if err := utils.ParseJSON(r, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := store.CreateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func handleTask(w http.ResponseWriter, r *http.Request, store *Store) {
	taskID, err := strconv.Atoi(r.PathValue("taskID"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		getTask(w, taskID, store)
	case "PUT":
		updateTask(w, r, taskID, store)
	case "DELETE":
		deleteTask(w, taskID, store)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTask(w http.ResponseWriter, taskID int, store *Store) {
	task, err := store.GetTaskByID(taskID)
	if err != nil {
		http.Error(w, "Error fetching task", http.StatusInternalServerError)
		return
	}

	if task == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func updateTask(w http.ResponseWriter, r *http.Request, taskID int, store *Store) {
	var updatedTask types.Task
	if err := utils.ParseJSON(r, &updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTask.ID = taskID
	err := store.UpdateTask(updatedTask)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func deleteTask(w http.ResponseWriter, taskID int, store *Store) {
	err := store.DeleteTask(taskID)
	if err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
