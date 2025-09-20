package model

import (
	"fmt"
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t Task) Print() {
	layout := "15:04___02.01"
	fmt.Printf("\n%d)_ %s \n %s \n Created at: %s \n Updated at: %s \n \n", t.Id, t.Status, t.Description, t.CreatedAt.Format(layout), t.UpdatedAt.Format(layout))
}
