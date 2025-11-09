# Task Manager API

This single-file documentation gives a concise overview of the API, model, and how to run and test it.

Base URL: http://localhost:8080

Endpoints

- GET /tasks
  - Returns: 200 OK — JSON array of tasks.

- GET /tasks/:id
  - Returns: 200 OK — JSON object for the task, or 400 Bad Request when not found.

- POST /tasks
  - Request: JSON Task object. Returns 200 OK with the created task on success.

- PUT /tasks/:id
  - Request: JSON with fields to update (partial updates applied for non-empty strings). Returns 200 OK with the updated task or 400 on error.

- DELETE /tasks/:id
  - Returns: 200 OK with the deleted task or 400 when not found.

Model (Task)

- id: string
- title: string
- description: string
- status: string (examples: Pending, In Progress, Completed)
- due_date: RFC3339 timestamp (JSON maps to Go time.Time)