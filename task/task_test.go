package task

import (
	"testing"
	"time"
)

func TestTask_Execute_Completed(t *testing.T) {
	tt := []struct {
		task        *Task
		executeMode string
		expected    string
	}{
		{
			task:        &Task{Id: "42"},
			executeMode: "shouldComplete",
			expected:    "completed",
		},
		{
			task:        &Task{Id: "52"},
			executeMode: "shouldFail",
			expected:    "failed",
		},
		{
			task:        &Task{Id: "62"},
			executeMode: "shouldTimeout",
			expected:    "timeout",
		},
	}

	for _, test := range tt {
		go test.task.Execute(test.executeMode)

		time.Sleep(time.Millisecond * 200)
		status, _ := test.task.Status()
		//fmt.Println(status)
		if status != test.expected {
			t.Fatalf("task should be %s. got=%s", test.expected, status)
		}
	}
}
