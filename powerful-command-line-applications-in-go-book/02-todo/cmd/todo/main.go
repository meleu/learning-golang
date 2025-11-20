package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"todo"
)

const TodoFilename = ".todo.json"

const usage = `
Developed by meleu
Copyright 2025

Usage information:
`

func main() {
	// Parsing command line flags
	// task := flag.String("task", "", "Task to be included in the ToDo list")
	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Usage = func() { fmt.Print(usage); flag.PrintDefaults() }
	flag.Parse()

	todoList := &todo.TodoList{}

	if err := todoList.Get(TodoFilename); err != nil {
		log.Fatal(err)
	}

	switch {
	case *list:
		fmt.Print(todoList)
	case *complete > 0:
		completeTask(todoList, *complete)
	case *add:
		addTask(todoList)
	default:
		log.Fatal("Invalid option")
	}
}

func completeTask(list *todo.TodoList, i int) {
	if err := list.Complete(i); err != nil {
		log.Fatal(err)
	}
	if err := list.Save(TodoFilename); err != nil {
		log.Fatal(err)
	}
}

func addTask(list *todo.TodoList) {
	task, err := readTask(os.Stdin, flag.Args()...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	list.Add(task)
	if err := list.Save(TodoFilename); err != nil {
		log.Fatal(err)
	}
}

func readTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}

	return s.Text(), nil
}
