package main

import (
	"context"
	"errors"
	"testing"
)

func TestMemoryTaskRepositoryCRUD(t *testing.T) {
	ctx := context.Background()
	repo := NewMemoryTaskRepository()

	task, err := repo.Create(ctx, "learn SQL")
	requireNoError(t, err)

	if task.ID != 1 {
		t.Fatalf("task ID = %d, want 1", task.ID)
	}

	updated, err := repo.UpdateDone(ctx, task.ID, true)
	requireNoError(t, err)

	if !updated.Done {
		t.Fatal("task should be marked done")
	}

	tasks, err := repo.List(ctx)
	requireNoError(t, err)

	if len(tasks) != 1 {
		t.Fatalf("got %d tasks, want 1", len(tasks))
	}

	err = repo.Delete(ctx, task.ID)
	requireNoError(t, err)

	tasks, err = repo.List(ctx)
	requireNoError(t, err)

	if len(tasks) != 0 {
		t.Fatalf("got %d tasks, want 0", len(tasks))
	}
}

func TestMemoryTaskRepositoryDeleteMissingTask(t *testing.T) {
	repo := NewMemoryTaskRepository()

	err := repo.Delete(context.Background(), 99)
	if !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("error = %v, want %v", err, ErrTaskNotFound)
	}
}

func requireNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
