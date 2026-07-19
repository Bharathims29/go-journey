# Student Marks REST API Curl Commands

Start PostgreSQL and create a database named `student_marks`.

Create or update `.env` with your real PostgreSQL password:

```text
DATABASE_URL=postgres://postgres:YOUR_PASSWORD@localhost:5432/student_marks?sslmode=disable
PORT=8040
```

Run the API:

```bash
cd server/rest
go run .
```

The server uses `PORT` from `.env`. With the example above, it runs on `http://localhost:8040`.

## Health Check

```bash
curl http://localhost:8040/health
```

## Create Student Mark

```bash
curl -X POST http://localhost:8040/marks \
  -H "Content-Type: application/json" \
  -d '{
    "student_name": "Bharathi",
    "roll_number": "R001",
    "subject": "Mathematics",
    "marks": 88,
    "max_marks": 100
  }'
```

## List All Marks

```bash
curl http://localhost:8040/marks
```

## Get One Mark Record

Replace `1` with the record ID.

```bash
curl http://localhost:8040/marks/1
```

## Update Mark Record

Replace `1` with the record ID.

```bash
curl -X PUT http://localhost:8040/marks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "student_name": "Bharathi",
    "roll_number": "R001",
    "subject": "Mathematics",
    "marks": 94,
    "max_marks": 100
  }'
```

## Delete Mark Record

Replace `1` with the record ID.

```bash
curl -X DELETE http://localhost:8040/marks/1
```

Successful delete returns `204 No Content`.

## Validation Error Example

This returns `400 Bad Request` because marks cannot be greater than max marks.

```bash
curl -X POST http://localhost:8040/marks \
  -H "Content-Type: application/json" \
  -d '{
    "student_name": "Bharath",
    "roll_number": "R001",
    "subject": "Science",
    "marks": 120,
    "max_marks": 100
  }'
```
