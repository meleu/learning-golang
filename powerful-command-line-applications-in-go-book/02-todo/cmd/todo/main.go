package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"todo"
)

const TodoFilename = ".todo.json"

func main() {
	// Parsing command line flags
	task := flag.String("task", "", "Task to be included in the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Usage = usage
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
	case *task != "":
		addTask(todoList, *task)
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

func addTask(list *todo.TodoList, task string) {
	list.Add(task)
	if err := list.Save(TodoFilename); err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed by meleu\n", os.Args[0])
	fmt.Fprintln(flag.CommandLine.Output(), "Copyright 2025")
	fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
	flag.PrintDefaults()
}
