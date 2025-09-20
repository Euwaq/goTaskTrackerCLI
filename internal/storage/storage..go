package storage

import "gott/internal/model"

type TaskRepo interface {
	AddTask(descripion string) (int, error)
	ListTasks(status string) ([]model.Task, error)
	UpdateTask(id int, dlc string) error
	MarkTask(id int, status string) error
	DeleteTask(id int) error
}
