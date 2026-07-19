package main

import (
	"bufio"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var errMarkNotFound = errors.New("student mark not found")

type StudentMark struct {
	ID          int64     `json:"id"`
	StudentName string    `json:"student_name"`
	RollNumber  string    `json:"roll_number"`
	Subject     string    `json:"subject"`
	Marks       int       `json:"marks"`
	MaxMarks    int       `json:"max_marks"`
	Grade       string    `json:"grade"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type StudentMarkInput struct {
	StudentName string `json:"student_name"`
	RollNumber  string `json:"roll_number"`
	Subject     string `json:"subject"`
	Marks       int    `json:"marks"`
	MaxMarks    int    `json:"max_marks"`
}

type MarkRepository struct {
	db *sql.DB
}

func NewMarkRepository(db *sql.DB) *MarkRepository {
	return &MarkRepository{db: db}
}

func (r *MarkRepository) Create(ctx context.Context, input StudentMarkInput) (StudentMark, error) {
	input = normalizeInput(input)
	grade := calculateGrade(input.Marks, input.MaxMarks)

	const query = `
		INSERT INTO student_marks (student_name, roll_number, subject, marks, max_marks, grade)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, student_name, roll_number, subject, marks, max_marks, grade, created_at, updated_at`

	var mark StudentMark
	err := r.db.QueryRowContext(
		ctx,
		query,
		input.StudentName,
		input.RollNumber,
		input.Subject,
		input.Marks,
		input.MaxMarks,
		grade,
	).Scan(
		&mark.ID,
		&mark.StudentName,
		&mark.RollNumber,
		&mark.Subject,
		&mark.Marks,
		&mark.MaxMarks,
		&mark.Grade,
		&mark.CreatedAt,
		&mark.UpdatedAt,
	)
	return mark, err
}

func (r *MarkRepository) List(ctx context.Context) ([]StudentMark, error) {
	const query = `
		SELECT id, student_name, roll_number, subject, marks, max_marks, grade, created_at, updated_at
		FROM student_marks
		ORDER BY id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var marks []StudentMark
	for rows.Next() {
		var mark StudentMark
		if err := scanMark(rows, &mark); err != nil {
			return nil, err
		}
		marks = append(marks, mark)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return marks, nil
}

func (r *MarkRepository) GetByID(ctx context.Context, id int64) (StudentMark, error) {
	const query = `
		SELECT id, student_name, roll_number, subject, marks, max_marks, grade, created_at, updated_at
		FROM student_marks
		WHERE id = $1`

	var mark StudentMark
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&mark.ID,
		&mark.StudentName,
		&mark.RollNumber,
		&mark.Subject,
		&mark.Marks,
		&mark.MaxMarks,
		&mark.Grade,
		&mark.CreatedAt,
		&mark.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return StudentMark{}, errMarkNotFound
	}

	return mark, err
}

func (r *MarkRepository) Update(ctx context.Context, id int64, input StudentMarkInput) (StudentMark, error) {
	input = normalizeInput(input)
	grade := calculateGrade(input.Marks, input.MaxMarks)

	const query = `
		UPDATE student_marks
		SET student_name = $1,
			roll_number = $2,
			subject = $3,
			marks = $4,
			max_marks = $5,
			grade = $6,
			updated_at = NOW()
		WHERE id = $7
		RETURNING id, student_name, roll_number, subject, marks, max_marks, grade, created_at, updated_at`

	var mark StudentMark
	err := r.db.QueryRowContext(
		ctx,
		query,
		input.StudentName,
		input.RollNumber,
		input.Subject,
		input.Marks,
		input.MaxMarks,
		grade,
		id,
	).Scan(
		&mark.ID,
		&mark.StudentName,
		&mark.RollNumber,
		&mark.Subject,
		&mark.Marks,
		&mark.MaxMarks,
		&mark.Grade,
		&mark.CreatedAt,
		&mark.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return StudentMark{}, errMarkNotFound
	}

	return mark, err
}

func (r *MarkRepository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM student_marks WHERE id = $1`, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errMarkNotFound
	}

	return nil
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanMark(scanner rowScanner, mark *StudentMark) error {
	return scanner.Scan(
		&mark.ID,
		&mark.StudentName,
		&mark.RollNumber,
		&mark.Subject,
		&mark.Marks,
		&mark.MaxMarks,
		&mark.Grade,
		&mark.CreatedAt,
		&mark.UpdatedAt,
	)
}

type Server struct {
	repo *MarkRepository
}

func NewServer(repo *MarkRepository) *Server {
	return &Server{repo: repo}
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.handleHealth)
	mux.HandleFunc("/marks", s.handleMarks)
	mux.HandleFunc("/marks/", s.handleMarkByID)
	return loggingMiddleware(mux)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) handleMarks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		marks, err := s.repo.List(r.Context())
		if err != nil {
			writeError(w, http.StatusInternalServerError, "could not list marks")
			return
		}
		writeJSON(w, http.StatusOK, marks)
	case http.MethodPost:
		s.createMark(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (s *Server) handleMarkByID(w http.ResponseWriter, r *http.Request) {
	id, err := idFromPath(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid mark id")
		return
	}

	switch r.Method {
	case http.MethodGet:
		mark, err := s.repo.GetByID(r.Context(), id)
		if err != nil {
			handleRepositoryError(w, err)
			return
		}
		writeJSON(w, http.StatusOK, mark)
	case http.MethodPut:
		s.updateMark(w, r, id)
	case http.MethodDelete:
		err := s.repo.Delete(r.Context(), id)
		if err != nil {
			handleRepositoryError(w, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (s *Server) createMark(w http.ResponseWriter, r *http.Request) {
	input, ok := decodeAndValidateInput(w, r)
	if !ok {
		return
	}

	mark, err := s.repo.Create(r.Context(), input)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create mark")
		return
	}

	writeJSON(w, http.StatusCreated, mark)
}

func (s *Server) updateMark(w http.ResponseWriter, r *http.Request, id int64) {
	input, ok := decodeAndValidateInput(w, r)
	if !ok {
		return
	}

	mark, err := s.repo.Update(r.Context(), id, input)
	if err != nil {
		handleRepositoryError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, mark)
}

func decodeAndValidateInput(w http.ResponseWriter, r *http.Request) (StudentMarkInput, bool) {
	defer r.Body.Close()

	var input StudentMarkInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return StudentMarkInput{}, false
	}

	if err := validateInput(input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return StudentMarkInput{}, false
	}

	return input, true
}

func validateInput(input StudentMarkInput) error {
	input = normalizeInput(input)

	if input.StudentName == "" {
		return errors.New("student_name is required")
	}
	if input.RollNumber == "" {
		return errors.New("roll_number is required")
	}
	if input.Subject == "" {
		return errors.New("subject is required")
	}
	if input.MaxMarks <= 0 {
		return errors.New("max_marks must be greater than zero")
	}
	if input.Marks < 0 {
		return errors.New("marks cannot be negative")
	}
	if input.Marks > input.MaxMarks {
		return errors.New("marks cannot be greater than max_marks")
	}

	return nil
}

func normalizeInput(input StudentMarkInput) StudentMarkInput {
	input.StudentName = strings.TrimSpace(input.StudentName)
	input.RollNumber = strings.TrimSpace(input.RollNumber)
	input.Subject = strings.TrimSpace(input.Subject)
	return input
}

func calculateGrade(marks, maxMarks int) string {
	if maxMarks <= 0 {
		return "Invalid"
	}

	percentage := float64(marks) / float64(maxMarks) * 100
	switch {
	case percentage >= 90:
		return "A+"
	case percentage >= 80:
		return "A"
	case percentage >= 70:
		return "B"
	case percentage >= 60:
		return "C"
	case percentage >= 50:
		return "D"
	default:
		return "F"
	}
}

func idFromPath(path string) (int64, error) {
	idText := strings.TrimPrefix(path, "/marks/")
	if idText == "" || strings.Contains(idText, "/") {
		return 0, errors.New("invalid id path")
	}

	return strconv.ParseInt(idText, 10, 64)
}

func handleRepositoryError(w http.ResponseWriter, err error) {
	if errors.Is(err, errMarkNotFound) {
		writeError(w, http.StatusNotFound, "mark record not found")
		return
	}

	writeError(w, http.StatusInternalServerError, "server error")
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Println("encode response:", err)
	}
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

func openDB() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/student_marks?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func migrate(ctx context.Context, db *sql.DB) error {
	const schema = `
		CREATE TABLE IF NOT EXISTS student_marks (
			id BIGSERIAL PRIMARY KEY,
			student_name TEXT NOT NULL,
			roll_number TEXT NOT NULL,
			subject TEXT NOT NULL,
			marks INTEGER NOT NULL CHECK (marks >= 0),
			max_marks INTEGER NOT NULL CHECK (max_marks > 0),
			grade TEXT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);

		CREATE INDEX IF NOT EXISTS idx_student_marks_roll_number
		ON student_marks (roll_number);`

	_, err := db.ExecContext(ctx, schema)
	return err
}

func serverAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return fmt.Sprintf(":%s", port)
}

func loadEnvFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.Trim(strings.TrimSpace(value), `"'`)
		if key == "" || os.Getenv(key) != "" {
			continue
		}

		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return scanner.Err()
}

func main() {
	if err := loadEnvFile(".env"); err != nil {
		log.Fatal("load .env:", err)
	}

	db, err := openDB()
	if err != nil {
		log.Fatal("connect database:", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := migrate(ctx, db); err != nil {
		log.Fatal("run migration:", err)
	}

	server := NewServer(NewMarkRepository(db))
	addr := serverAddress()

	log.Println("student marks REST API listening on", addr)
	log.Fatal(http.ListenAndServe(addr, server.Routes()))
}
