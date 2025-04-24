# GoLang Portfolio – Tasks Project

This is a simple Go application designed to demonstrate structuring a Go project using a modular and clean architecture. It can serve as a learning resource or a starter template for more complex applications.

## 📁 Project Structure

```
tasks/
├── cmd/        # Application commands (CLI entry points)
├── config/     # Configuration files and logic
├── data/       # Data sources and repositories
├── models/     # Data models and structs
├── main.go     # Main application entry point
├── go.mod      # Go module definition
├── go.sum      # Dependency checksums
```

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.18 or higher

### Running the Project

Clone the repository and run the application:

```bash
git clone https://github.com/Loonguinho/GoLangPortifolio.git
cd GoLangPortifolio/tasks
go run main.go
```

## 🧪 Example CLI Usage

You can run the available CLI commands using:

Example output:

```bash
A CLI for managing tasks

Usage:
  tasks [command]

Available Commands:
  create      Create a new task
  delete      Delete a task
  list        List all tasks
```

Example: creating a task

```bash
go run main.go create "Buy groceries"
```

Example: listing tasks (use '-a' or '--all' to list all tasks)

```bash
go run main.go list
```

## 🧩 Features

- Modular project layout
- Basic CLI commands (create, list, delete)
- Configuration management

## 📦 Dependencies

All dependencies are managed via Go Modules (`go.mod` and `go.sum`).

To install dependencies:

```bash
go mod tidy
```

## 🛠️ Future Improvements

- Add unit tests
- Implement REST API
- Use a database (e.g., SQLite, PostgreSQL)
- Add logging and better error handling

Made by [Loonguinho](https://github.com/Loonguinho)
Inspired by [dreamsofcode-io](https://github.com/dreamsofcode-io)
