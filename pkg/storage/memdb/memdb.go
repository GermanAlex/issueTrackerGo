package memdb

import "issueTackerGo/pkg/storage/postgres"

type DB []postgres.Task

func (db DB) Tasks(int, int) ([]postgres.Task, error) {
	return db, nil
}

func (db DB) LabelTasks(int) ([]postgres.Task, error) {
	return db, nil
}

func (db DB) NewTaskTx(postgres.Task) (int, error) {
	return 0, nil
}

func (db DB) UpdateTaskTx(int, postgres.Task) error {
	return nil
}

func (db DB) DeleteTaskTx(int) error {
	return nil
}
