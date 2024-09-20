package task

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/victorluisca/go-todo-app/types"
	"github.com/victorluisca/go-todo-app/utils"
)

func RegisterRoutes(router *http.ServeMux, store types.TaskStore) {
	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) { handleTasks(w, r, store) })
	router.HandleFunc("/task/{taskID}", func(w http.ResponseWriter, r *http.Request) { handleTask(w, r, store) })
}

func handleTasks(w http.ResponseWriter, r *http.Request, store types.TaskStore) {
	switch r.Method {
	case http.MethodGet:
		getAllTasks(w, store)
	case http.MethodPost:
		createTask(w, r, store)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAllTasks(w http.ResponseWriter, store types.TaskStore) {
	tasks, err := store.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, tasks); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func createTask(w http.ResponseWriter, r *http.Request, store types.TaskStore) {
	var task types.Task
	if err := utils.ParseJSON(r, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(task); err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, errors.Error(), http.StatusBadRequest)
		return
	}

	if err := store.CreateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// w.WriteHeader(http.StatusCreated) -- superfluous call
	if err := utils.WriteJSON(w, http.StatusCreated, task); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func handleTask(w http.ResponseWriter, r *http.Request, store types.TaskStore) {
	taskID, err := strconv.Atoi(r.PathValue("taskID"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTask(w, taskID, store)
	case http.MethodPut:
		updateTask(w, r, taskID, store)
	case http.MethodDelete:
		deleteTask(w, taskID, store)
	case http.MethodPatch:
		updateTaskCompletion(w, taskID, store)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTask(w http.ResponseWriter, taskID int, store types.TaskStore) {
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
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func updateTask(w http.ResponseWriter, r *http.Request, taskID int, store types.TaskStore) {
	existingTask, err := store.GetTaskByID(taskID)
	if err != nil {
		http.Error(w, "Error fetching task", http.StatusInternalServerError)
		return
	}
	if existingTask == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	var updatedTask types.Task
	if err := utils.ParseJSON(r, &updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(updatedTask); err != nil {
		errors := err.(validator.ValidationErrors)
		http.Error(w, errors.Error(), http.StatusBadRequest)
		return
	}

	updatedTask.ID = taskID
	err = store.UpdateTask(updatedTask)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, updatedTask); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

func deleteTask(w http.ResponseWriter, taskID int, store types.TaskStore) {
	existingTask, err := store.GetTaskByID(taskID)
	if err != nil {
		http.Error(w, "Error fetching task", http.StatusInternalServerError)
		return
	}
	if existingTask == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	err = store.DeleteTask(taskID)
	if err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateTaskCompletion(w http.ResponseWriter, taskID int, store types.TaskStore) {
	existingTask, err := store.GetTaskByID(taskID)
	if err != nil {
		http.Error(w, "Error fetching task", http.StatusInternalServerError)
		return
	}
	if existingTask == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	err = store.ToggleTaskCompletion(taskID)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Error updating the task completion status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
