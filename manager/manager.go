package manager

import (
	"fmt"
	"github.com/sadysnaat/assignment3/task"
	"sync"
	"time"
)

type Manager struct {
	tasks sync.Map
}

func (m *Manager) Enqueue(t *task.Task) {
	_, ok := m.tasks.LoadOrStore(t.Id, t)
	if ok {
		fmt.Printf("task %s already in enqueued", t.Id)
		return
	}
	go t.Execute("")
}

func (m *Manager) Delete(t *task.Task) {
	m.tasks.Delete(t.Id)
}

func (m *Manager) Purge(wg *sync.WaitGroup) {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("starting purge")
		remaining := false
		m.tasks.Range(func(key, value interface{}) bool {
			remaining = true
			t := value.(*task.Task)
			status, err := t.Status()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("task found", t.Id, status)
			if status == "completed" {
				m.Delete(t)
			}

			if status == "failed" || status == "timeout" {
				fmt.Println("requeue task", t.Id)
				m.Delete(t)
				m.Enqueue(t)
			}
			return true
		})

		if !remaining {
			fmt.Println("no more items remaining")
			wg.Done()
		}

	}
}
