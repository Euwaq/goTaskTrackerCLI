package main

import (
	"gott/internal/storage"
	"gott/internal/task"
	"os"
)

func main() {
	s := task.NewService(storage.Read())
	s.Cmd(os.Args)
	storage.Write(s.Data())
}
