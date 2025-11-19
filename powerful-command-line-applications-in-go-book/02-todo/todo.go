// Package todo have a CRUD implementation for TODO lists
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type task struct {
	Description string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TodoList []task

func (l *TodoList) Add(newTask string) {
	task := task{
		Description: newTask,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, task)
}

func (l *TodoList) Complete(id int) error {
	// TODO: por que isso?
	list := *l
	err := l.isValidID(id)
	if err != nil {
		return err
	}

	task := &list[id-1]
	task.Done = true
	task.CompletedAt = time.Now()

	return nil
}

func (l *TodoList) Delete(id int) error {
	// TODO: essas manobras com ponteiros ainda não estão claras pra mim
	list := *l
	err := l.isValidID(id)
	if err != nil {
		return err
	}

	// TODO: idem
	*l = append(list[:id-1], list[id:]...)

	return nil
}

func (l *TodoList) Save(filename string) error {
	json, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, json, 0o644)
}

func (l *TodoList) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	// TODO: decidir se isso aqui permanece (faz diferença?)
	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *TodoList) isValidID(id int) error {
	list := *l
	if id < 0 || id > len(list) {
		return fmt.Errorf("id %d does not exist", id)
	}

	return nil
}

func (l *TodoList) String() string {
	var formatted string

	for i, task := range *l {
		prefix := "[ ]"
		if task.Done {
			prefix = "[X]"
		}

		formatted += fmt.Sprintf("%s %d: %s\n", prefix, i+1, task.Description)
	}

	return formatted
}
