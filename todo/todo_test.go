package todo_test

import (
	"github.com/watife/todo/todo"
	"os"
	"testing"
)

func TestList_Add(t *testing.T) {
	var l todo.List
	taskName := "new task"

	l = l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %s, got %s", taskName, l[0].Task)
	}
}

func TestList_Complete(t *testing.T) {
	var l todo.List
	taskName := "new task"

	l = l.Add(taskName)

	if l[0].Done {
		t.Errorf("expected %s to be incomplete", taskName)
	}

	err := l.Complete(1)
	if err != nil {
		return
	}

	if !l[0].Done {
		t.Errorf("expected %s to be complete", taskName)
	}
}

func TestList_Delete(t *testing.T) {
	var l todo.List

	tasks := []string{"task 1", "task 2", "task 3"}
	for _, task := range tasks {
		l = l.Add(task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("expected %s, got %s", tasks[0], l[0].Task)
	}

	l, _ = l.Delete(2)

	if len(l) != 2 {
		t.Errorf("expected %d, got %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("expected %s, got %s", tasks[2], l[1].Task)

	}

}

func TestList_SaveGet(t *testing.T) {
	var l1, l2 todo.List
	taskName := "new task"

	l1 = l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("expected %s, got %s", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Fatalf("Error removing temp file: %s", err)
		}
	}(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("expected %s, got %s", l1[0].Task, l2[0].Task)
	}

}
