package task

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Service struct {
	tasks []Task
}

func NewService(t []Task) Service {
	return Service{
		tasks: t,
	}
}

func (s Service) Data() []Task {
	return s.tasks
}

func (s Service) Cmd(args []string) {
	switch args[1] {
	case "help":
		help()
	case "add":
		s.add(args[2])
	case "update":
		s.getTask(args[2]).update(args[3])
	case "delete":
		s.delete(args[2])
	case "mark":
		s.getTask(args[2]).mark(args[3])
	case "list":
		s.list(args[2])
	default:
		log.Fatal(fmt.Errorf("Unknown commad: %s", args[1]))
	}
}

func (s Service) atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil || i > len(s.tasks) {
		log.Fatal("Write id of task after command")
	}
	return i - 1
}

func (s Service) getTask(arg string) *Task {
	i := s.atoi(arg)
	return &s.tasks[i]
}

func help() {
	fmt.Printf("\n Write command after name of program: ")
	fmt.Printf("\n add [descripion] ")
	fmt.Printf("\n mark [id] [status (todo/in-progress/done)]")
	fmt.Printf("\n update [id] [addition]")
	fmt.Printf("\n delete [id]")
	fmt.Printf("\n list [status (all/todo/in-progress/done)]")
}

func (s Service) add(d string) {
	t := Task{
		Id:          len(s.tasks) + 1,
		Description: d,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	s.tasks = append(s.tasks, t)
}

func (s Service) delete(arg string) {
	i := s.atoi(arg)
	s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
	for k := range s.tasks {
		s.tasks[k].Id = k + 1
	}
}

func (s Service) list(which string) {
	switch which {
	case "all":
		for _, task := range s.tasks {
			print(task)
		}
	case "todo":
		fallthrough
	case "in-progress":
		fallthrough
	case "done":
		for _, task := range s.tasks {
			if task.Status == which {
				print(task)
			}
		}
	default:
		log.Fatal(fmt.Errorf("Unknown status of task: %s", which))
	}
}
