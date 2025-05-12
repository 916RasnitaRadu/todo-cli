package types

import (
	"fmt"
	"strconv"
	"time"
)

type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	Done        bool
}

func TaskFromCSV(task []string) Task {
	id, _ := strconv.Atoi(task[0])
	done, _ := strconv.ParseBool(task[3])
	createdAt, err := time.Parse(time.DateTime, task[2])
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	return Task{
		ID:          id,
		Description: task[1],
		CreatedAt:   createdAt,
		Done:        done,
	}
}
