// I'm defining this as a different package (todo != todo_test) so the test
// access only the exported types, variables and functions. This practice
// ensures the tests only access the exposed API.
package todo_test

import (
	"os"
	"testing"
	"time"

	"todo"
)

func TestTodoList(t *testing.T) {
	t.Run("starts empty", func(t *testing.T) {
		todoList := todo.TodoList{}
		wantLen := 0
		gotLen := len(todoList)

		if gotLen != wantLen {
			t.Errorf("got: %d, want: %d", gotLen, wantLen)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("create a task with Done status as false", func(t *testing.T) {
		todoList := todo.TodoList{}
		newTask := "study Go"
		todoList.Add(newTask)

		gotDescription := todoList[0].Description
		gotDone := todoList[0].Done

		if gotDescription != newTask || gotDone != false {
			t.Errorf(
				"\n got: %v/%v\nwant: %v/%v",
				gotDescription, gotDone,
				newTask, false,
			)
		}
	})

	t.Run("create a task with a CreatedAt", func(t *testing.T) {
		todoList := todo.TodoList{}
		newTask := "study Go"

		beforeTime := time.Now()
		todoList.Add(newTask)
		afterTime := time.Now()

		gotDescription := todoList[0].Description
		gotCreatedAt := todoList[0].CreatedAt

		assertString(t, gotDescription, newTask)
		assertTimeBetween(t, beforeTime, gotCreatedAt, afterTime)
	})
}

func TestComplete(t *testing.T) {
	t.Run("marks the task's Done status as true", func(t *testing.T) {
		taskDescription := "study Go"
		todoList := todo.TodoList{}

		todoList.Add(taskDescription)
		todoList.Complete(1)

		gotDescription := todoList[0].Description
		gotDone := todoList[0].Done

		if gotDone != true {
			t.Errorf(
				"\n got: %v/%v\nwant: %v/%v",
				gotDescription, gotDone,
				taskDescription, true,
			)
		}
	})

	t.Run("returns error when given ID doesn't exist", func(t *testing.T) {
		todoList := todo.TodoList{}

		err := todoList.Complete(123)

		if err == nil {
			t.Error("expected an error, got none")
		}
	})

	t.Run("assign a timestamp to the CompletedAt attribute", func(t *testing.T) {
		taskDescription := "study Go"
		todoList := todo.TodoList{}
		todoList.Add(taskDescription)

		beforeTime := time.Now()
		todoList.Complete(1)
		afterTime := time.Now()

		gotDescription := todoList[0].Description
		gotCompletedAt := todoList[0].CompletedAt

		if gotDescription != taskDescription {
			t.Errorf("\n got: %v\nwant: %v", gotDescription, taskDescription)
		}

		assertTimeBetween(t, beforeTime, gotCompletedAt, afterTime)
	})
}

func TestDelete(t *testing.T) {
	todoList := todo.TodoList{}

	tasks := []string{
		"study Go",
		"do the dishes",
		"walk the dog",
	}

	for _, t := range tasks {
		todoList.Add(t)
	}

	assertString(t, todoList[0].Description, tasks[0])

	todoList.Delete(2)

	if len(todoList) != 2 {
		t.Errorf("expected list length %d, got %d", 2, len(todoList))
	}

	assertString(t, todoList[1].Description, tasks[2])
}

func TestSaveGet(t *testing.T) {
	list1 := todo.TodoList{}
	list2 := todo.TodoList{}

	taskName := "study Go"
	list1.Add(taskName)

	assertString(t, list1[0].Description, taskName)

	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating temp file: %s", err)
	}
	filename := tempFile.Name()
	defer os.Remove(filename)

	if err := list1.Save(filename); err != nil {
		t.Fatalf("error saving list to file: %s", err)
	}
	if err := list2.Get(filename); err != nil {
		t.Fatalf("error getting list from file: %s", err)
	}

	assertString(t, list1[0].Description, list2[0].Description)
}

// Helper functions for assertions
// ----------------------------------------------------------------------

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("\n got: %q\nwant: %q", got, want)
	}
}

func assertTimeBetween(t testing.TB, beforeTime, currentTime, afterTime time.Time) {
	t.Helper()
	if !beforeTime.Before(currentTime) || !currentTime.Before(afterTime) {
		t.Errorf(
			"expected CreatedAt (%v)\nto be between\n %v\n and\n %v",
			currentTime, beforeTime, afterTime,
		)
	}
}
