package storage

import (
	"encoding/json"
	"fmt"
	"gott/internal/model"
	"log"
	"os"
	"time"
)

type JsonRepo struct {
	tasks []model.Task
}

func NewJsonRepo() JsonRepo {
	return JsonRepo{
		tasks: read(),
	}
}

func (jr JsonRepo) AddTask(description string) (int, error) {
	t := model.Task{
		Id:          len(jr.tasks),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	jr.tasks = append(jr.tasks, t)
	return t.Id, jr.write()
}

func (jr JsonRepo) ListTasks(status string) error {
	switch status {
	case "all":
		for _, task := range jr.tasks {
			task.Print()
		}
	case "todo":
		fallthrough
	case "in-progress":
		fallthrough
	case "done":
		for _, task := range jr.tasks {
			if task.Status == status {
				task.Print()
			}
		}
	default:
		return fmt.Errorf("Unknown status of task: %s", status)
	}
	return nil
}

func (jr JsonRepo) UpdateTask(id int, dlc string) error {

	jr.tasks[id].Description += "\n" + dlc
	jr.tasks[id].UpdatedAt = time.Now()
	return jr.write()
}

func (jr JsonRepo) MarkTask(id int, status string) error {
	switch status {
	case "todo":
		fallthrough
	case "in-progress":
		fallthrough
	case "done":
		jr.tasks[id].Status = status
	default:
		return fmt.Errorf("Unknown status of task: %s", status)
	}
	return jr.write()
}

func (jr JsonRepo) DeleteTask(id int) error {
	jr.tasks = append(jr.tasks[:id], jr.tasks[id+1:]...)
	for k := range jr.tasks {
		jr.tasks[k].Id = k + 1
	}
	return jr.write()
}

func read() []model.Task {
	tasks := make([]model.Task, 0)
	data, err := os.ReadFile("data.json")
	if err != nil {
		_, ok := err.(*os.PathError)
		if ok {
			return tasks
		}
		panic(err)
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal("Data in file is not correct.")
	}
	return tasks
}

func (jr JsonRepo) write() error {
	bytes, err := json.Marshal(jr.tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile("data.json", bytes, 0644)
	if err != nil {
		return err
	}
	return nil

}
