# Student Marks REST API Code Explanation

This backend is a simple REST API for a student mark portal.

It uses:
- `net/http` for the HTTP server
- `database/sql` for database access
- PostgreSQL as the database
- `github.com/lib/pq` as the PostgreSQL driver

## Files

- `main.go`: API server, handlers, repository, validation, migration, and app startup.
- `main_test.go`: Small unit tests for grade calculation, validation, and path ID parsing.
- `schema.sql`: PostgreSQL table and index.
- `CURL_COMMANDS.md`: Commands to test the API manually.

## Data Model

`StudentMark` is the response model.

It contains:
- `id`: Database ID
- `student_name`: Student name
- `roll_number`: Student roll number
- `subject`: Subject name
- `marks`: Marks scored
- `max_marks`: Maximum marks
- `grade`: Calculated grade
- `created_at`: Record creation time
- `updated_at`: Record update time

`StudentMarkInput` is the request model used for create and update.

The client sends name, roll number, subject, marks, and max marks. The server calculates the grade.

## Database Table

The `student_marks` table stores each subject mark record.

Important columns:
- `id BIGSERIAL PRIMARY KEY`: Auto-generated ID.
- `student_name TEXT NOT NULL`: Required student name.
- `roll_number TEXT NOT NULL`: Required roll number.
- `marks INTEGER NOT NULL`: Scored marks.
- `max_marks INTEGER NOT NULL`: Maximum possible marks.
- `grade TEXT NOT NULL`: Grade calculated by Go.

The app runs `migrate` on startup to create the table if it does not already exist.

## Repository Functions

`MarkRepository` owns all database queries.

`Create(ctx, input)`:
- Validates and normalizes input before the handler calls it.
- Calculates the grade.
- Inserts a row into PostgreSQL.
- Returns the newly created record using `RETURNING`.

`List(ctx)`:
- Reads all mark records.
- Orders by `id`.
- Scans rows into `StudentMark` structs.

`GetByID(ctx, id)`:
- Reads one mark record by ID.
- Returns `errMarkNotFound` when no row exists.

`Update(ctx, id, input)`:
- Replaces the student mark fields.
- Recalculates grade.
- Updates `updated_at`.
- Returns the updated row.

`Delete(ctx, id)`:
- Deletes one record by ID.
- Uses `RowsAffected` to detect missing records.

## Handler Functions

`Routes()`:
- Creates the HTTP router.
- Registers `/health`, `/marks`, and `/marks/{id}`.
- Wraps all routes with logging middleware.

`handleHealth`:
- Handles `GET /health`.
- Returns `{"status":"ok"}`.

`handleMarks`:
- Handles collection routes.
- `GET /marks` lists all records.
- `POST /marks` creates a record.

`handleMarkByID`:
- Handles single-record routes.
- `GET /marks/{id}` reads one record.
- `PUT /marks/{id}` updates one record.
- `DELETE /marks/{id}` deletes one record.

`createMark`:
- Decodes JSON.
- Validates input.
- Calls `repo.Create`.
- Returns `201 Created`.

`updateMark`:
- Decodes JSON.
- Validates input.
- Calls `repo.Update`.
- Returns `200 OK`.

## Helper Functions

`decodeAndValidateInput`:
- Reads request JSON.
- Validates required fields and mark limits.
- Writes `400 Bad Request` if input is invalid.

`validateInput`:
- Checks required fields.
- Checks `max_marks > 0`.
- Checks `marks >= 0`.
- Checks `marks <= max_marks`.

`calculateGrade`:
- Converts marks to percentage.
- Returns grade:
- `A+` for 90 and above
- `A` for 80 and above
- `B` for 70 and above
- `C` for 60 and above
- `D` for 50 and above
- `F` below 50

`idFromPath`:
- Extracts the ID from paths like `/marks/1`.
- Returns an error for invalid IDs.

`writeJSON`:
- Sets `Content-Type: application/json`.
- Writes the status code.
- Encodes a Go value as JSON.

`writeError`:
- Sends JSON errors like `{"error":"message"}`.

`openDB`:
- Reads `DATABASE_URL`.
- Uses a default local PostgreSQL URL if it is missing.
- Pings the database before starting the server.

`loadEnvFile`:
- Reads `.env` from the current folder.
- Loads simple `KEY=value` lines into environment variables.
- Does not overwrite values already exported in the shell.

`serverAddress`:
- Reads `PORT`.
- Defaults to `8080`.

## API Routes

- `GET /health`: Check server status.
- `POST /marks`: Create student mark record.
- `GET /marks`: List all mark records.
- `GET /marks/{id}`: Get one mark record.
- `PUT /marks/{id}`: Update one mark record.
- `DELETE /marks/{id}`: Delete one mark record.
