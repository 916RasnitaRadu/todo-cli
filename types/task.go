package types

import "strconv"

type Task struct {
	ID          int
	Description string
	CreatedAt   string
	Done        bool
}

func TaskFromCSV(task []string) Task {
	id, _ := strconv.Atoi(task[0])
	done, _ := strconv.ParseBool(task[3])

	return Task{
		ID:          id,
		Description: task[1],
		CreatedAt:   task[2],
		Done:        done,
	}
}
