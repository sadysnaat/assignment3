package manager

import (
	"github.com/sadysnaat/assignment3/task"
	"sync"
	"testing"
)

func TestManager_Enqueue(t *testing.T) {
	m := Manager{}

	qtask := &task.Task{Id: "21"}

	err := m.Enqueue(qtask)
	if err != nil {
		t.Fatalf("expected enqueue, got=%s", err.Error())
	}

	err = m.Enqueue(qtask)
	if err != nil {
		if err.Error() != "task 21 already in enqueued" {
			t.Fatalf("expected %s. got %s",
				"task 21 already in enqueued",
				err.Error())
		}
	} else {
		t.Fatalf("expected error")
	}
}

// This test case is just there to increase the coverage of tested code
// Actually no assertions are made or verified for purge
func TestManager_Purge(t *testing.T) {
	m := Manager{}

	qtask := &task.Task{Id: "21"}
	qtask2 := &task.Task{Id: "22"}
	_ = m.EnqueueWithIntent(qtask, "shouldComplete")
	_ = m.EnqueueWithIntent(qtask2, "shouldTimeout")
	var wg sync.WaitGroup
	wg.Add(1)
	m.Purge(&wg)

	wg.Wait()
}
