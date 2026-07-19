package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	handler := NewServer(NewTaskStore()).Routes()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
}

func TestListTasks(t *testing.T) {
	handler := NewServer(NewTaskStore()).Routes()
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var tasks []Task
	if err := json.NewDecoder(rec.Body).Decode(&tasks); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("got %d tasks, want 1", len(tasks))
	}
}

func TestCreateTask(t *testing.T) {
	handler := NewServer(NewTaskStore()).Routes()
	body := bytes.NewBufferString(`{"title":"learn handlers"}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusCreated)
	}

	var task Task
	if err := json.NewDecoder(rec.Body).Decode(&task); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if task.Title != "learn handlers" {
		t.Fatalf("title = %q, want %q", task.Title, "learn handlers")
	}
}

func TestCreateTaskRequiresTitle(t *testing.T) {
	handler := NewServer(NewTaskStore()).Routes()
	body := bytes.NewBufferString(`{"title":" "}`)
	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestUpdateTaskDone(t *testing.T) {
	handler := NewServer(NewTaskStore()).Routes()
	body := bytes.NewBufferString(`{"done":true}`)
	req := httptest.NewRequest(http.MethodPatch, "/tasks/1", body)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var task Task
	if err := json.NewDecoder(rec.Body).Decode(&task); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if !task.Done {
		t.Fatal("task should be done")
	}
}

func TestDeleteTask(t *testing.T) {
	handler := NewServer(NewTaskStore()).Routes()
	req := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNoContent)
	}
}

func TestDeleteTaskNotFound(t *testing.T) {
	handler := NewServer(NewTaskStore()).Routes()
	req := httptest.NewRequest(http.MethodDelete, "/tasks/999", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}
