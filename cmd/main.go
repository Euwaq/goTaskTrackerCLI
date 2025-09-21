package main

import (
	"gott/internal/storage"
	"gott/internal/task"
	"os"
)

func main() {
	jr, err := storage.NewJsonRepo("data.json")
	if err != nil {
		panic(err)
	}
	s := task.NewService(jr)
	s.Cmd(os.Args)

}
