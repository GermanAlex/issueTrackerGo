package storage

import "issueTackerGo/pkg/storage/postgres"

type Interface interface {
	Tasks(int, int) ([]postgres.Task, error)
	LabelTasks(int) ([]postgres.Task, error)
	NewTaskTx(postgres.Task) (int, error)
	UpdateTaskTx(int, postgres.Task) error
	DeleteTaskTx(int) error
}
