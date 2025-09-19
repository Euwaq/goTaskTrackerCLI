package task

import (
	"fmt"
	"log"
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) update(dlc string) {
	t.Description += "\n" + dlc
	t.UpdatedAt = time.Now()
}

func (t *Task) mark(stat string) {
	switch stat {
	case "todo":
		fallthrough
	case "in-progress":
		fallthrough
	case "done":
		t.Status = stat
	default:
		log.Fatal(fmt.Errorf("Unknown status of task: %s", stat))
	}
}

func (t Task) print() {
	layout := "15:04___02.01"
	fmt.Printf("\n%d)_ %s \n %s \n Created at: %s \n Updated at: %s \n \n", t.Id, t.Status, t.Description, t.CreatedAt.Format(layout), t.UpdatedAt.Format(layout))
}
