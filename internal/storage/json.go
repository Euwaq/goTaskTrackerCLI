package storage

import (
	"encoding/json"
	"gott/internal/task"
	"log"
	"os"
)

func Read() []task.Task {
	tasks := make([]task.Task, 0)
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

func Write(tasks []task.Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	os.WriteFile("data.json", bytes, 0644)

}
