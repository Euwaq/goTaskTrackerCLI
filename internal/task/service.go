package task

import (
	"fmt"
	"gott/internal/storage"
	"strconv"
)

type Service struct {
	repo storage.TaskRepo
}

func NewService(r storage.TaskRepo) Service {
	return Service{repo: r}
}

func (s Service) Cmd(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("Commad is not written")
	}
	var num int
	var err error
	switch args[1] {
	case "mark":
		if len(args) < 4 {
			return fmt.Errorf("Not anough args for mark")
		}
		fallthrough
	case "delete":
		fallthrough
	case "update":
		if len(args) < 3 {
			return fmt.Errorf("Need id to %s", args[1])
		}
		num, err = strconv.Atoi(args[2])
		if err != nil {
			return err
		}
		if num < 0 {
			return fmt.Errorf("Id can't be < 0")
		}

	}
	switch args[1] {
	case "help":
		help()
	case "add":
		_, err = s.repo.AddTask(args[2])
	case "update":
		err = s.repo.UpdateTask(num, args[3])
	case "delete":
		err = s.repo.DeleteTask(num)
	case "mark":
		err = s.repo.MarkTask(num, args[3])
	case "list":
		err = s.repo.ListTasks(args[2])
	default:
		return fmt.Errorf("Unknown commad: %s", args[1])
	}

	return err
}

func help() {
	fmt.Printf("\n Write command after name of program: ")
	fmt.Printf("\n add [descripion] ")
	fmt.Printf("\n mark [id] [status (todo/in-progress/done)]")
	fmt.Printf("\n update [id] [addition]")
	fmt.Printf("\n delete [id]")
	fmt.Printf("\n list [status (all/todo/in-progress/done)]")
}
