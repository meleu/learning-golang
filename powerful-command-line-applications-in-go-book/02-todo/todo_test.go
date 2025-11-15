package todo_test

import (
	"testing"
	"todo"
)

func TestAdd(t *testing.T) {
	list := todo.List{}
	taskName := "New Task"
	list.Add(taskName)

	got := list[0].Task

	if got != taskName {
		t.Errorf("got %q, want %q", got, taskName)
	}
}

func TestComplete(t *testing.T) {
	list := todo.List{}
	taskName := "New Task"
	list.Add(taskName)

	if list[0].Done {
		t.Errorf("New task should not be completed.")
	}

	list.Complete(1)

	if !list[0].Done {
		t.Errorf("New task should be completed.")
	}
}
