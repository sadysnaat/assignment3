package task

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	Id          string
	IsCompleted bool
	status      string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (t *Task) Status() (string, error) {
	if t == nil {
		return "", errors.New("no such task")
	}

	return t.status, nil
}

// Execute is an idempotent function if task is running it's noop
func (t *Task) Execute(intent string) {
	if t.status == "running" {
		return
	}

	t.status = "running"

	// c,f & to values decide task end state
	// c is random wait time to complete
	// f is random wait time to fail
	// to is random wait time to timeout
	var c, to, f int
	switch intent {
	case "shouldComplete":
		// f > to > c
		// task will always complete
		f = rand.Intn(1000)
		to = rand.Intn(f)
		c = rand.Intn(to)
	case "shouldFail":
		// c > to > f
		// task will always fail
		c = rand.Intn(1000)
		to = rand.Intn(c)
		f = rand.Intn(to)
	case "shouldTimeout":
		// c > f > to
		// task will always timeout
		c = rand.Intn(1000)
		f = rand.Intn(c)
		to = rand.Intn(f)
	default:
		// without any intent a task may end up in either of the three states
		c = rand.Intn(1000)
		to = rand.Intn(1000)
		f = rand.Intn(1000)
	}

	fmt.Printf("[task] %s will be completed after %d ms or timeout after %d ms or failed after %dms\n", t.Id, c, to, f)
	select {
	case <-time.After(time.Millisecond * time.Duration(c)):
		fmt.Printf("[task] %s completed\n", t.Id)
		t.IsCompleted = true
		t.status = "completed"
		return
	case <-time.After(time.Millisecond * time.Duration(to)):
		fmt.Printf("[task] %s timeout occured\n", t.Id)
		t.status = "timeout"
		return
	case <-time.After(time.Millisecond * time.Duration(f)):
		fmt.Printf("[task] %s failed\n", t.Id)
		t.status = "failed"
		return
	}
}
