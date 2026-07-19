package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type CreateTaskRequest struct {
	Title string `json:"title"`
}

type UpdateTaskRequest struct {
	Done bool `json:"done"`
}

type TaskStore struct {
	mu     sync.Mutex
	nextID int
	tasks  []Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		nextID: 1,
		tasks: []Task{
			{ID: 1, Title: "learn net/http", Done: false},
		},
	}
}

func (s *TaskStore) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := make([]Task, len(s.tasks))
	copy(tasks, s.tasks)
	return tasks
}

func (s *TaskStore) Create(title string) Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++
	task := Task{ID: s.nextID, Title: title, Done: false}
	s.tasks = append(s.tasks, task)
	return task
}

func (s *TaskStore) UpdateDone(id int, done bool) (Task, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks[i].Done = done
			return s.tasks[i], true
		}
	}

	return Task{}, false
}

func (s *TaskStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return true
		}
	}

	return false
}

type Server struct {
	store *TaskStore
}

func NewServer(store *TaskStore) *Server {
	return &Server{store: store}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/tasks", s.handleTasks)
	mux.HandleFunc("/tasks/", s.handleTask)

	return loggingMiddleware(mux)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, s.store.List())
	case http.MethodPost:
		s.createTask(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	var input CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	title := strings.TrimSpace(input.Title)
	if title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	task := s.store.Create(title)
	writeJSON(w, http.StatusCreated, task)
}

func (s *Server) handleTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPatch:
		s.updateTaskDone(w, r)
	case http.MethodDelete:
		s.deleteTask(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func taskIDFromPath(path string) (int, error) {
	idText := strings.TrimPrefix(path, "/tasks/")
	return strconv.Atoi(idText)
}

func (s *Server) updateTaskDone(w http.ResponseWriter, r *http.Request) {
	idText := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idText)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	var input UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	task, ok := s.store.UpdateDone(id, input.Done)
	if !ok {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}

	writeJSON(w, http.StatusOK, task)
}

func (s *Server) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := taskIDFromPath(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	if ok := s.store.Delete(id); !ok {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func serverPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}

	return port
}

func main() {
	server := NewServer(NewTaskStore())
	addr := ":" + serverPort()

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, server.Routes()))
}
