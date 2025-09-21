package storage

import (
	"encoding/json"
	"fmt"
	"gott/internal/model"
	"os"
	"time"
)

type JsonRepo struct {
	FileName string
	MaxId    int
	Tasks    map[int]model.Task
}

func NewJsonRepo(fileName string) (JsonRepo, error) {
	jr := JsonRepo{
		FileName: fileName,
		MaxId:    0,
		Tasks:    make(map[int]model.Task, 0),
	}
	data, err1 := os.ReadFile(fileName)
	if err1 != nil {
		return jr, err1
	}
	err2 := json.Unmarshal(data, &jr)
	return jr, err2
}

func (jr JsonRepo) AddTask(description string) (int, error) {
	jr.MaxId++
	t := model.Task{
		Id:          jr.MaxId,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	jr.Tasks[jr.MaxId] = t
	return t.Id, jr.write()
}

func (jr JsonRepo) ListTasks(status string) (map[int]model.Task, error) {
	switch status {
	case "all":
		return jr.Tasks, nil
	case "todo":
		fallthrough
	case "in-progress":
		fallthrough
	case "done":
		list := make(map[int]model.Task, 0)
		for id, task := range jr.Tasks {
			if task.Status == status {
				list[id] = task
			}
		}
		return list, nil
	default:
		return nil, fmt.Errorf("Unknown status of task: %s", status)
	}
}

func (jr JsonRepo) UpdateTask(id int, dlc string) error {
	task, ok := jr.Tasks[id]
	if !ok {
		return fmt.Errorf("Have not task with id=%d", id)
	}
	task.Description += "\n" + dlc
	task.UpdatedAt = time.Now()
	jr.Tasks[id] = task
	return jr.write()
}

func (jr JsonRepo) MarkTask(id int, status string) error {
	switch status {
	case "todo":
		fallthrough
	case "in-progress":
		fallthrough
	case "done":
		task, ok := jr.Tasks[id]
		if !ok {
			return fmt.Errorf("Have not task with id=%d", id)
		}
		task.Status = status
		jr.Tasks[id] = task
	default:
		return fmt.Errorf("Unknown status of task: %s", status)
	}
	return jr.write()
}

func (jr JsonRepo) DeleteTask(id int) error {
	delete(jr.Tasks, id)
	return jr.write()
}

func (jr JsonRepo) write() error {
	bytes, err := json.Marshal(jr)
	if err != nil {
		return err
	}
	err = os.WriteFile(jr.FileName, bytes, 0644)
	return err

}
