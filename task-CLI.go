package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var tasks = []Task{}

func main() {
	read()
	cmd(os.Args)
	write()
}

func cmd(args []string) {
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

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func help() {
	fmt.Printf("\n Write command after name of program: ")
	fmt.Printf("\n add [descripion] ")
	fmt.Printf("\n mark [id] [status (todo/in-progress/done)]")
	fmt.Printf("\n update [id] [addition]")
	fmt.Printf("\n delete [id]")
	fmt.Printf("\n list [status (all/todo/in-progress/done)]")
}

func delete(arg string) {
	i := atoi(arg)
	tasks = append(tasks[:i], tasks[i+1:]...)
	for k := range tasks {
		tasks[k].Id = k + 1
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

func (t *Task) update(dlc string) {
	t.Description += "\n" + dlc
	t.UpdatedAt = time.Now()
}

func (t *Task) mark(stat string) {
	switch stat {
	case "todo":
		fallthrough
	case "in-progress":
		fallthrough
	case "done":
		t.Status = stat
	default:
		log.Fatal(fmt.Errorf("Unknown status of task: %s", stat))
	}
}

func list(which string) {
	switch which {
	case "all":
		for _, task := range tasks {
			print(task)
		}
	case "todo":
		fallthrough
	case "in-progress":
		fallthrough
	case "done":
		for _, task := range tasks {
			if task.Status == which {
				print(task)
			}
		}
	default:
		log.Fatal(fmt.Errorf("Unknown status of task: %s", which))
	}
}

func (t Task) print() {
	layout := "15:04___02.01"
	fmt.Printf("\n%d)_ %s \n %s \n Created at: %s \n Updated at: %s \n \n", t.Id, t.Status, t.Description, t.CreatedAt.Format(layout), t.UpdatedAt.Format(layout))
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

func read() {
	data, err := os.ReadFile("data.json")
	if err != nil {
		_, ok := err.(*os.PathError)
		if ok {
			return
		}
		panic(err)
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Fatal("Data in file is not correct.")
	}
}

func write() {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	os.WriteFile("data.json", bytes, 0644)

}
