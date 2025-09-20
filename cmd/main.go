package main

import (
	"gott/internal/storage"
	"gott/internal/task"
	"os"
)

func main() {
	s := task.NewService(storage.NewJsonRepo())
	s.Cmd(os.Args)

}
