package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l List) Add(task string) List {
	return append(l, item{Task: task, CreatedAt: time.Now()})
}

func (l List) Complete(i int) error {
	if i <= 0 || i > len(l) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	(l)[i-1].Done = true
	(l)[i-1].CompletedAt = time.Now()

	return nil
}

func (l List) Delete(i int) (List, error) {
	if i <= 0 || i > len(l) {
		return nil, fmt.Errorf("Item %d does not exist", i)
	}

	l = append(l[:i-1], l[i:]...)

	return l, nil
}

func (l List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return fmt.Errorf("marshaling list: %w", err)
	}

	return os.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("Error reading file: %w", err)
	}

	return json.Unmarshal(file, &l)
}
