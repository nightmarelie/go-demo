package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

const todoFile = ".todo.json"

func main() {
	args := os.Args
	if len(args) < 2 {
		usage()
		os.Exit(1)
	}

	switch args[1] {
	case "add":
		if len(args) < 3 {
			fmt.Println("error: missing task text")
			os.Exit(1)
		}
		text := strings.Join(args[2:], " ")
		if err := addTask(text); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	case "list":
		if err := listTasks(); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	case "done":
		if len(args) < 3 {
			fmt.Println("error: missing task id")
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("error: id must be a number")
			os.Exit(1)
		}
		if err := markDone(id); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println(`todo - a tiny CLI

Usage:
  todo add <text>    Add a new task
  todo list          List tasks
  todo done <id>     Mark a task as done`)
}

func todoPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, todoFile), nil
}

func loadTasks() ([]Task, error) {
	path, err := todoPath()
	if err != nil {
		return nil, err
	}
	b, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil // empty list on first run
		}
		return nil, err
	}
	var tasks []Task
	if err := json.Unmarshal(b, &tasks); err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", path, err)
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	path, err := todoPath()
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

func nextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func addTask(text string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	t := Task{
		ID:        nextID(tasks),
		Text:      text,
		Done:      false,
		CreatedAt: time.Now(),
	}
	tasks = append(tasks, t)
	if err := saveTasks(tasks); err != nil {
		return err
	}
	fmt.Printf("added %d: %s\n", t.ID, t.Text)
	return nil
}

func listTasks() error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		fmt.Println("no tasks yet — add one with: todo add <text>")
		return nil
	}
	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%s %d: %s\n", status, t.ID, t.Text)
	}
	return nil
}

func markDone(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("task %d not found", id)
	}
	if err := saveTasks(tasks); err != nil {
		return err
	}
	fmt.Printf("done: %d\n", id)
	return nil
}
