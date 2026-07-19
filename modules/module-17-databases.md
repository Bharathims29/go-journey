# Module 17: Databases

## Goal

Learn how Go applications store and retrieve data using SQL databases.

By the end of this module, you should understand:
- `database/sql`
- SQL basics
- Connecting to PostgreSQL or SQLite
- Queries and commands
- Scanning rows
- Transactions
- Migrations
- Repository pattern basics

## 1. Why Use a Database?

Files are useful for simple persistence, but databases are better when you need:
- Fast lookup
- Many records
- Filtering and sorting
- Concurrent access
- Transactions
- Reliable updates

Most Go web applications put database code behind a small repository layer.

## 2. The `database/sql` Package

Go's standard library has a generic SQL package:

```go
import "database/sql"
```

`database/sql` does not include a database driver by itself.

For real projects, you add a driver such as:
- SQLite driver
- PostgreSQL driver
- MySQL driver

The common shape is:

```go
db, err := sql.Open("sqlite", "tasks.db")
if err != nil {
	return err
}
defer db.Close()
```

The exact driver name depends on the driver you install.

## 3. SQL Basics

SQL is the language used to work with relational databases.

Create a table:

```sql
CREATE TABLE tasks (
	id INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	done BOOLEAN NOT NULL
);
```

Insert a row:

```sql
INSERT INTO tasks (title, done) VALUES (?, ?);
```

Read rows:

```sql
SELECT id, title, done FROM tasks ORDER BY id;
```

Update a row:

```sql
UPDATE tasks SET done = ? WHERE id = ?;
```

Delete a row:

```sql
DELETE FROM tasks WHERE id = ?;
```

## 4. Queries and Commands

Use `ExecContext` for commands that do not return rows:

```go
_, err := db.ExecContext(ctx, "DELETE FROM tasks WHERE id = ?", id)
```

Use `QueryContext` for many rows:

```go
rows, err := db.QueryContext(ctx, "SELECT id, title, done FROM tasks")
```

Use `QueryRowContext` for one row:

```go
err := db.QueryRowContext(ctx, query, id).Scan(&task.ID, &task.Title, &task.Done)
```

Always close rows:

```go
defer rows.Close()
```

## 5. Scanning Rows

Scanning copies database columns into Go variables.

```go
var task Task
err := rows.Scan(&task.ID, &task.Title, &task.Done)
```

The order of `Scan` arguments must match the order of selected columns.

## 6. Transactions

A transaction groups multiple database operations.

Either all operations succeed, or none of them are saved.

```go
tx, err := db.BeginTx(ctx, nil)
if err != nil {
	return err
}
defer tx.Rollback()

_, err = tx.ExecContext(ctx, "UPDATE tasks SET done = ? WHERE id = ?", true, id)
if err != nil {
	return err
}

return tx.Commit()
```

`Rollback` after `Commit` is harmless. The `defer tx.Rollback()` pattern protects you if an earlier step fails.

## 7. Migrations

A migration is a database change saved as code.

Examples:
- Create a table
- Add a column
- Add an index
- Rename a table

Small projects can start with a simple `schema.sql` file.

Production projects usually use a migration tool.

## 8. Repository Pattern

A repository hides storage details behind methods.

```go
type TaskRepository interface {
	Create(ctx context.Context, title string) (Task, error)
	List(ctx context.Context) ([]Task, error)
	UpdateDone(ctx context.Context, id int, done bool) (Task, error)
	Delete(ctx context.Context, id int) error
}
```

Your HTTP handlers can depend on this interface instead of raw SQL.

This makes code easier to test.

## 9. Run the Example

Open this example folder:

[module-17-databases](../examples/module-17-databases)

From the module example folder, run:

```bash
cd examples/module-17-databases
go run .
```

Run the tests:

```bash
go test -v
```

The example includes:
- Repository pattern
- Create, list, update, and delete operations
- A `database/sql` repository implementation shape
- Transaction code with `BeginTx`, `ExecContext`, `Commit`, and `Rollback`
- An in-memory repository used for runnable tests without an external driver

## 10. Practice Tasks

1. Create a `tasks` table with `id`, `title`, and `done`.
2. Write a repository method to create a task.
3. Write a repository method to list tasks.
4. Write a repository method to update `done`.
5. Write a repository method to delete a task.
6. Add a transaction that updates a task and inserts an audit row.
7. Connect Module 16's HTTP API to a repository interface.

## Checkpoint

You are ready for the next module when:
- You understand why `database/sql` needs a driver.
- You can write basic `SELECT`, `INSERT`, `UPDATE`, and `DELETE` statements.
- You can scan query results into Go structs.
- You can use a transaction for multi-step changes.
- You can place database code behind a repository.
