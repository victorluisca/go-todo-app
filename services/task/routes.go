package task

import (
	"net/http"
	"strconv"
	"time"

	"github.com/victorluisca/go-todo-app/types"
	"github.com/victorluisca/go-todo-app/utils"
)

var tasks = []types.Task{
	{ID: 1, Title: "Tarefa 1", Priority: "Medium", CreatedAt: time.Now()},
}

func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/tasks", handleTasks)
	router.HandleFunc("/task/{taskID}", handleTask)
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllTasks(w)
	case "POST":
		createTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAllTasks(w http.ResponseWriter) {
	if err := utils.WriteJSON(w, http.StatusOK, tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task types.Task
	if err := utils.ParseJSON(r, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.ID = len(tasks) + 1
	task.CreatedAt = time.Now()
	tasks = append(tasks, task)

	if err := utils.WriteJSON(w, http.StatusCreated, task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleTask(w http.ResponseWriter, r *http.Request) {
	taskID, err := strconv.Atoi(r.PathValue("taskID"))
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		getTask(w, taskID)
	case "PUT":
		updateTask(w, r, taskID)
	case "DELETE":
		deleteTask(w, taskID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTask(w http.ResponseWriter, taskID int) {
	for _, task := range tasks {
		if task.ID == taskID {
			if err := utils.WriteJSON(w, http.StatusOK, task); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func updateTask(w http.ResponseWriter, r *http.Request, taskID int) {
	var updatedTask types.Task
	if err := utils.ParseJSON(r, &updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := utils.Validate.Struct(updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			updatedTask.ID = taskID
			tasks[i] = updatedTask
			if err := utils.WriteJSON(w, http.StatusOK, updatedTask); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}

func deleteTask(w http.ResponseWriter, taskID int) {
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Task not found", http.StatusNotFound)
}
