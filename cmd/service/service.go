package service

import (
	"fmt"
	"issueTackerGo/pkg/storage"
	"issueTackerGo/pkg/storage/postgres"
	"log"
	"os"
)

// Интерфейс БД.
var db storage.Interface

func main() {
	var err error
	pwd := os.Getenv("dbpass")
	if pwd == "" {
		os.Exit(1)
	}
	connstr :=
		"postgres://postgres:" +
			pwd + "@ubuntu-server.northeurope.cloudapp.azure.com/tasks"
	// присвоение переменной типа интерфейс конкретной реализации БД
	db, err = postgres.New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
}
