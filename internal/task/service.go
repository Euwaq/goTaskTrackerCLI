package task

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

var tasks = []Task{}

func Cmd(args []string) {
	switch args[1] {
	case "help":
		help()
	case "add":
		add(args[2])
	case "update":
		getTask(args[2]).update(args[3])
	case "delete":
		delete(args[2])
	case "mark":
		getTask(args[2]).mark(args[3])
	case "list":
		list(args[2])
	default:
		log.Fatal(fmt.Errorf("Unknown commad: %s", args[1]))
	}
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil || i > len(tasks) {
		log.Fatal("Write id of task after command")
	}
	return i - 1
}

func getTask(arg string) *Task {
	i := atoi(arg)
	return &tasks[i]
}

func help() {
	fmt.Printf("\n Write command after name of program: ")
	fmt.Printf("\n add [descripion] ")
	fmt.Printf("\n mark [id] [status (todo/in-progress/done)]")
	fmt.Printf("\n update [id] [addition]")
	fmt.Printf("\n delete [id]")
	fmt.Printf("\n list [status (all/todo/in-progress/done)]")
}

func add(d string) {
	t := Task{
		Id:          len(tasks) + 1,
		Description: d,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, t)
}
