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

func (t *Task) Execute(intent string) {
	if t.status == "running" {
		return
	}

	var c, to, f int
	switch intent {
	case "shouldComplete":
		f = rand.Intn(100)
		to = rand.Intn(f)
		c = rand.Intn(to)
	case "shouldFail":
		c = rand.Intn(100)
		to = rand.Intn(c)
		f = rand.Intn(to)
	case "shouldTimeout":
		c = rand.Intn(100)
		f = rand.Intn(c)
		to = rand.Intn(f)
	default:
		c = rand.Intn(100)
		to = rand.Intn(100)
		f = rand.Intn(100)
	}

	fmt.Printf("task %s will be completed after %d ms\n or timeout after %d ms\n or failed after %dms\n", t.Id, c, to, f)
	select {
	case <-time.After(time.Millisecond * time.Duration(c)):
		fmt.Printf("task %s completed\n", t.Id)
		t.IsCompleted = true
		t.status = "completed"
		return
	case <-time.After(time.Millisecond * time.Duration(to)):
		fmt.Printf("task %s timeout occured\n", t.Id)
		t.status = "timeout"
		return
	case <-time.After(time.Millisecond * time.Duration(f)):
		fmt.Printf("task %s failed\n", t.Id)
		t.status = "failed"
		return
	}
}
