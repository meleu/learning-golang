package main_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "test-todo"
	filename = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		log.Fatalf("cannot build tool %s: %s", binName, err)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(filename)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	newTask := "test task number 1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("add a new task from arguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", newTask)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	newTask2 := "test task number 2"
	t.Run("add a new task from stdin", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdin, newTask2)
		cmdStdin.Close()

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("list all saved tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		want := fmt.Sprintf("[ ] 1: %s\n[ ] 2: %s\n", newTask, newTask2)
		got := string(out)

		if got != want {
			t.Errorf("\n got: %q\nwant: %q", got, want)
		}
	})

	t.Run("mark a task as completed", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "1")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}

		cmd = exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		want := fmt.Sprintf("[X] 1: %s\n[ ] 2: %s\n", newTask, newTask2)
		got := string(out)

		if got != want {
			t.Errorf("\n got: %q\nwant: %q", got, want)
		}
	})
}
