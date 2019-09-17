package manager

import (
	"errors"
	"fmt"
	"github.com/sadysnaat/assignment3/task"
	"sync"
	"time"
)

type Manager struct {
	tasks sync.Map
}

func (m *Manager) Enqueue(t *task.Task) error {
	_, ok := m.tasks.LoadOrStore(t.Id, t)
	if ok {
		fmt.Printf("task %s already in enqueued", t.Id)
		return errors.New(fmt.Sprintf("task %s already in enqueued", t.Id))
	}
	go t.Execute("")
	return nil
}

func (m *Manager) EnqueueWithIntent(t *task.Task, intent string) error {
	_, ok := m.tasks.LoadOrStore(t.Id, t)
	if ok {
		fmt.Printf("task %s already in enqueued", t.Id)
		return errors.New(fmt.Sprintf("task %s already in enqueued", t.Id))
	}
	go t.Execute(intent)
	return nil
}

func (m *Manager) Delete(t *task.Task) {
	m.tasks.Delete(t.Id)
}

func (m *Manager) Purge(wg *sync.WaitGroup) {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("[manager] starting purge")
		remaining := false
		m.tasks.Range(func(key, value interface{}) bool {
			remaining = true
			t := value.(*task.Task)
			status, err := t.Status()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("[manager] task found", t.Id, status)
			if status == "completed" {
				m.Delete(t)
			}

			if status == "failed" || status == "timeout" {
				fmt.Println("[manager] requeue task", t.Id)
				m.Delete(t)
				err := m.Enqueue(t)
				if err != nil {
					fmt.Println("[manager]", err)
				}
			}
			return true
		})

		if !remaining {
			fmt.Println("[manager] no more items remaining")
			wg.Done()
			return
		}

	}
}
