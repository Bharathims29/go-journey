# Module 16: HTTP and APIs

## Goal

Learn how to build a basic REST-style API in Go using the standard `net/http` package.

By the end of this module, you should understand:
- `net/http`
- Handlers
- Routing basics
- Request and response handling
- JSON APIs
- HTTP status codes
- Middleware basics
- Environment variables
- Testing handlers with `httptest`

## 1. What Is an HTTP API?

An HTTP API lets other programs communicate with your program over HTTP.

Common API actions:
- `GET` reads data
- `POST` creates data
- `PUT` or `PATCH` updates data
- `DELETE` removes data

Go can build HTTP APIs using the standard library. You do not need a framework to start.

## 2. The `net/http` Package

The `net/http` package provides servers, handlers, requests, and responses.

Minimal server:

```go
func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
```

`http.ResponseWriter` writes the response.

`*http.Request` contains the method, path, body, headers, and query values.

## 3. Handlers

A handler is code that responds to an HTTP request.

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
```

Most APIs use handlers to:
- Check the HTTP method
- Decode input
- Call application logic
- Encode output
- Set the right status code

## 4. Routing Basics

Routing means matching a request path to the right handler.

Go's standard library includes `http.ServeMux`:

```go
mux := http.NewServeMux()
mux.HandleFunc("/tasks", tasksHandler)
```

For beginner APIs, a simple `ServeMux` is enough.

Inside a handler, you can check the method:

```go
switch r.Method {
case http.MethodGet:
	listTasks(w, r)
case http.MethodPost:
	createTask(w, r)
default:
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}
```

## 5. JSON Responses

APIs commonly send JSON.

```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(value)
```

For request bodies:

```go
var input CreateTaskRequest
err := json.NewDecoder(r.Body).Decode(&input)
```

Always handle JSON decode errors.

## 6. Status Codes

HTTP status codes tell the client what happened.

Common codes:
- `200 OK` for successful reads
- `201 Created` for successful creates
- `204 No Content` for successful deletes with no body
- `400 Bad Request` for invalid input
- `404 Not Found` for missing resources
- `405 Method Not Allowed` for wrong methods
- `500 Internal Server Error` for unexpected server problems

Status codes are part of your API contract.

## 7. Middleware Basics

Middleware wraps a handler to add shared behavior.

Example logging middleware:

```go
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
```

Middleware is useful for:
- Logging
- Authentication
- Request IDs
- Panic recovery
- CORS

## 8. Environment Variables

Environment variables let you configure the app without changing code.

```go
port := os.Getenv("PORT")
if port == "" {
	port = "8080"
}
```

This is useful when running the same app locally, in Docker, or on a server.

## 9. Testing HTTP Handlers

Use `httptest` to test handlers without starting a real server.

```go
req := httptest.NewRequest(http.MethodGet, "/health", nil)
rec := httptest.NewRecorder()

handler.ServeHTTP(rec, req)

if rec.Code != http.StatusOK {
	t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
}
```

This keeps API tests fast and reliable.

## 10. Run the Example

Open this example folder:

[module-16-http-apis](../examples/module-16-http-apis)

From the module example folder, run:

```bash
cd examples/module-16-http-apis
go run .
```

The server starts on port `8080` by default.

Try it from another terminal:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/tasks
curl -X POST http://localhost:8080/tasks -d '{"title":"learn handlers"}'
curl -X PATCH http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"done":true}'
curl -X DELETE http://localhost:8080/tasks/1
```

Run the tests:

```bash
go test -v
```

The example includes:
- A health endpoint
- A JSON task list endpoint
- A JSON create task endpoint
- A JSON update task endpoint
- A delete task endpoint
- Request validation
- Status codes
- Logging middleware
- Handler tests with `httptest`

## 11. Practice Tasks

1. Add `GET /health` that returns `{"status":"ok"}`.
2. Add `GET /tasks` that returns a list of tasks.
3. Add `POST /tasks` that validates a JSON body.
4. Return `400 Bad Request` for empty task titles.
5. Add logging middleware.
6. Read the server port from the `PORT` environment variable.
7. Test your handlers with `httptest`.

## Checkpoint

You are ready for the next module when:
- You can create a handler function.
- You can route requests with `http.ServeMux`.
- You can read JSON request bodies.
- You can write JSON responses.
- You can choose basic HTTP status codes.
- You can wrap handlers with middleware.
- You can test handlers without starting a real server.
