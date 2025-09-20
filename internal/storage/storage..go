package storage

type TaskRepo interface {
	AddTask(descripion string) (int, error)
	ListTasks(status string) error
	UpdateTask(id int, dlc string) error
	MarkTask(id int, status string) error
	DeleteTask(id int) error
}
