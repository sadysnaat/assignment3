package main

import (
	"fmt"
	"github.com/sadysnaat/assignment3/manager"
	"github.com/sadysnaat/assignment3/task"
	"sync"
	"time"
)

func main() {
	m := &manager.Manager{}
	var wg sync.WaitGroup

	wg.Add(1)
	go m.Purge(&wg)
	time.Sleep(time.Millisecond * 100)
	wg.Add(1)
	go m.Purge(&wg)

	for i := 0; i < 10; i++ {
		t := &task.Task{
			Id: fmt.Sprintf("%d", i),
		}
		m.Enqueue(t)
	}

	wg.Wait()
}
