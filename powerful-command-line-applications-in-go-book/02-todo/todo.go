package todo

import (
	"fmt"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (list *List) Add(task string) {
	newTask := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*list = append(*list, newTask)
}

func (list *List) Complete(i int) error {
	newList := *list

	if i <= 0 || i > len(newList) {
		return fmt.Errorf("item %d does not exist", i)
	}

	newList[i-1].Done = true
	newList[i-1].CompletedAt = time.Now()

	return nil
}
