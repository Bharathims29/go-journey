package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"sync"
)

var ErrTaskNotFound = errors.New("task not found")

type Task struct {
	ID    int
	Title string
	Done  bool
}

type TaskRepository interface {
	Create(ctx context.Context, title string) (Task, error)
	List(ctx context.Context) ([]Task, error)
	UpdateDone(ctx context.Context, id int, done bool) (Task, error)
	Delete(ctx context.Context, id int) error
}

type MemoryTaskRepository struct {
	mu     sync.Mutex
	nextID int
	tasks  map[int]Task
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		nextID: 0,
		tasks:  make(map[int]Task),
	}
}

func (r *MemoryTaskRepository) Create(ctx context.Context, title string) (Task, error) {
	if err := ctx.Err(); err != nil {
		return Task{}, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.nextID++
	task := Task{ID: r.nextID, Title: title, Done: false}
	r.tasks[task.ID] = task
	return task, nil
}

func (r *MemoryTaskRepository) List(ctx context.Context) ([]Task, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	tasks := make([]Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return tasks, nil
}

func (r *MemoryTaskRepository) UpdateDone(ctx context.Context, id int, done bool) (Task, error) {
	if err := ctx.Err(); err != nil {
		return Task{}, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := r.tasks[id]
	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Done = done
	r.tasks[id] = task
	return task, nil
}

func (r *MemoryTaskRepository) Delete(ctx context.Context, id int) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.tasks[id]; !ok {
		return ErrTaskNotFound
	}

	delete(r.tasks, id)
	return nil
}

type SQLTaskRepository struct {
	db *sql.DB
}

func NewSQLTaskRepository(db *sql.DB) *SQLTaskRepository {
	return &SQLTaskRepository{db: db}
}

func (r *SQLTaskRepository) Create(ctx context.Context, title string) (Task, error) {
	const query = `INSERT INTO tasks (title, done) VALUES (?, ?) RETURNING id, title, done`

	var task Task
	err := r.db.QueryRowContext(ctx, query, title, false).Scan(&task.ID, &task.Title, &task.Done)
	return task, err
}

func (r *SQLTaskRepository) List(ctx context.Context) ([]Task, error) {
	const query = `SELECT id, title, done FROM tasks ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *SQLTaskRepository) UpdateDone(ctx context.Context, id int, done bool) (Task, error) {
	const query = `UPDATE tasks SET done = ? WHERE id = ? RETURNING id, title, done`

	var task Task
	err := r.db.QueryRowContext(ctx, query, done, id).Scan(&task.ID, &task.Title, &task.Done)
	if errors.Is(err, sql.ErrNoRows) {
		return Task{}, ErrTaskNotFound
	}

	return task, err
}

func (r *SQLTaskRepository) Delete(ctx context.Context, id int) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM tasks WHERE id = ?`, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (r *SQLTaskRepository) MarkDoneWithAudit(ctx context.Context, id int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(ctx, `UPDATE tasks SET done = ? WHERE id = ?`, true, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return ErrTaskNotFound
	}

	_, err = tx.ExecContext(ctx, `INSERT INTO task_audit (task_id, message) VALUES (?, ?)`, id, "marked done")
	if err != nil {
		return err
	}

	return tx.Commit()
}

func main() {
	ctx := context.Background()
	repo := NewMemoryTaskRepository()

	first, _ := repo.Create(ctx, "learn database/sql")
	second, _ := repo.Create(ctx, "write repository tests")
	updated, _ := repo.UpdateDone(ctx, first.ID, true)
	_ = repo.Delete(ctx, second.ID)
	tasks, _ := repo.List(ctx)

	fmt.Println("Module 17: Databases")
	fmt.Println("updated task:", updated)
	fmt.Println("remaining tasks:", tasks)
}
