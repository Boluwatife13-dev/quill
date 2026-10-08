# Quill Backend

Go (Gin) + MongoDB API for **Quill** — a classic Notepad with a simple built-in To-Do system.

The backend provides authentication, note management, task management, and protected API routes for the Quill application.

## Requirements

* Go 1.27 or newer
* MongoDB running locally
* MongoDB default URI: `mongodb://localhost:27017`

## Run it

```bash
go mod tidy
go run ./cmd/server
```

The API runs on:

```text
http://localhost:6000
```

## Environment

Create a `.env` file in the project root:

```env
PORT=6000
MONGO_URI=mongodb://localhost:27017
JWT_SECRET=quill_secret_key
```

### Environment variables

| Variable     | Description                               |
| ------------ | ----------------------------------------- |
| `PORT`       | Port used by the backend                  |
| `MONGO_URI`  | MongoDB connection string                 |
| `JWT_SECRET` | Secret used to sign and verify JWT tokens |

Do not expose `JWT_SECRET` publicly.

## Authentication

Quill uses **JWT authentication** for protected routes.

After logging in, the client sends the token using:

```text
Authorization: Bearer YOUR_TOKEN
```

The authentication middleware checks the token before allowing access to protected endpoints.

## Endpoints

All responses are JSON.

### Authentication

| Method   | Path                 | Who       |
| -------- | -------------------- | --------- |
| `POST`   | `/api/auth/register` | Anyone    |
| `POST`   | `/api/auth/login`    | Anyone    |
| `GET`    | `/api/auth/me`       | Signed in |
| `PUT`    | `/api/auth/me`       | Signed in |
| `DELETE` | `/api/auth/me`       | Signed in |

### Notes

| Method   | Path             | Who       |
| -------- | ---------------- | --------- |
| `POST`   | `/api/notes`     | Signed in |
| `GET`    | `/api/notes`     | Signed in |
| `GET`    | `/api/notes/:id` | Signed in |
| `PUT`    | `/api/notes/:id` | Signed in |
| `DELETE` | `/api/notes/:id` | Signed in |

### Tasks

| Method   | Path             | Who       |
| -------- | ---------------- | --------- |
| `POST`   | `/api/tasks`     | Signed in |
| `GET`    | `/api/tasks`     | Signed in |
| `GET`    | `/api/tasks/:id` | Signed in |
| `PUT`    | `/api/tasks/:id` | Signed in |
| `DELETE` | `/api/tasks/:id` | Signed in |

## Features

### Notes

* Create notes
* View notes
* Edit notes
* Delete notes
* Store note titles and content
* Track creation and update dates

### To-Do

* Create tasks
* Edit tasks
* Complete tasks
* Delete tasks
* Task priorities
* Due dates
* Categories
* Completion status

### User Accounts

* User registration
* User login
* User profile
* Update profile
* Delete account
* JWT authentication

## Rules

* Users must authenticate before accessing protected resources.
* Passwords should never be stored as plain text.
* JWT tokens are required for protected endpoints.
* Users should only access their own notes and tasks.
* Invalid or expired JWT tokens are rejected.
* Unauthorized requests receive an HTTP `401` response.
* MongoDB stores users, notes, and tasks.

## Database

Quill uses MongoDB.

The database name is:

```text
quill
```

Main collections:

```text
users
notes
tasks
```

### User document

Users contain:

* ID
* Name
* Email
* Password
* CreatedAt
* UpdatedAt

### Note document

Notes contain:

* ID
* UserID
* Title
* Content
* CreatedAt
* UpdatedAt

### Task document

Tasks contain:

* ID
* UserID
* Title
* Description
* Priority
* DueDate
* Category
* Completed
* CreatedAt
* UpdatedAt

## Structure

```text
quill-backend/

├── cmd/
│   └── server/
│       └── main.go

├── config/
│   └── config.go

├── database/
│   └── database.go

├── handlers/
│   ├── note_handler.go
│   ├── task_handler.go
│   └── user_handler.go

├── middleware/
│   └── auth.go

├── models/
│   ├── note.go
│   ├── task.go
│   └── user.go

├── routes/
│   └── routes.go

├── utils/
│   └── response.go

├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Technologies

* **Go** — Backend language
* **Gin** — HTTP web framework
* **MongoDB** — Database
* **MongoDB Go Driver** — MongoDB connection
* **JWT** — Authentication
* **bcrypt** — Password hashing
* **godotenv** — Environment configuration

## Development Status

### Completed

* Project structure
* Go module
* Dependencies
* Environment configuration
* MongoDB connection
* User model
* Note model
* Task model
* API routes
* Authentication middleware

### In Progress

* User handlers
* Note handlers
* Task handlers
* User authentication logic
* Password hashing
* JWT generation
* MongoDB CRUD operations
* Request validation
* Error handling

## Project Goal

Quill is designed to keep the simplicity of a classic Notepad while adding a lightweight To-Do system.

The goal is to provide a clean application where users can:

**Write → Organize → Complete**

without the unnecessary complexity of a large productivity platform.

---

**Quill — Simple notes. Simple tasks. Nothing complicated.**
