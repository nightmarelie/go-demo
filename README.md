# Usfull links
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with...)  
  A practical guide to learning Go through test-driven development.

- [Learn X in Y Minutes: Go](https://learnxinyminutes.com/docs/go/)  
  Great for quickly learning Go syntax.

- [Effective Go](https://go.dev/doc/effective_go)  
  A must-read for writing idiomatic Go code.

# Tiny TODO CLI (Go)

A minimal command-line task manager written in Go. Tasks are stored as JSON in your home directory at `~/.todo.json`.

## Prerequisites

- Go installed (`go version` should work). Install from `https://golang.org/dl/` if needed.

## Installation

```bash
# Run directly
go run .

# Or build a binary
go build -o todo

# Or install to $(go env GOPATH)/bin
go install .
```

## Usage

```text
todo - a tiny CLI

Usage:
  todo add <text>    Add a new task
  todo list          List tasks
  todo done <id>     Mark a task as done
```

Examples:

```bash
# Add tasks
todo add "buy milk"
todo add "write blog post"

# List tasks
todo list
# [ ] 1: buy milk
# [ ] 2: write blog post

# Complete a task
todo done 1

# List again
todo list
# [x] 1: buy milk
# [ ] 2: write blog post
```

## Data Location

- Tasks are stored in `~/.todo.json` (created automatically on first run).

## Project Structure

```
go-demo/
├── go.mod
├── main.go
└── README.md
```

## Development

- Format: `go fmt ./...`
- Tidy modules: `go mod tidy`
- Test (if tests exist): `go test ./...`

## Troubleshooting

- If the binary won't run: `chmod +x ./todo`
- If you see module errors: `go mod tidy`
- If `~/.todo.json` is missing/corrupt, delete it and re-run; it will be recreated.

## License

MIT
