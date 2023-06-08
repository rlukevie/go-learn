package todo_test

import (
	"os"
	"testing"

	"github.com/rlukevie/go-learn/rggo/interacting/todo"
)

// TestAdd tests he Add method of the List type.
func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected task name %q, got %q instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}
	l.Add("New Task")

	if l[0].Done {
		t.Errorf("expected task to be incomplete")
	}
	// fmt.Println(l[0].Done)
	l.Complete(1)
	// fmt.Println(l[0].Done)

	if !l[0].Done {
		t.Errorf("expected task to be complete")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, task := range tasks {
		l.Add(task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("expected task name %q, got %q instead", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("expected list length to be 2, got %d instead", len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("expected task name %q, got %q instead", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating tempfile: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("expected task name %q, got %q instead", l1[0].Task, l2[0].Task)
	}
}
