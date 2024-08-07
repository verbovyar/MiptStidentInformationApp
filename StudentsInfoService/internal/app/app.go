package app

import (
	"StudentsInfoService/internal/handlers"
	"StudentsInfoService/internal/kafka"
	"StudentsInfoService/internal/repositories/db"
	"StudentsInfoService/pkg/postgres"
	"fmt"
	"net/http"
)

const port = ":9000"
const connectionString = "postgres://postgres:Verbov3232132121@localhost:5432/studentsDB"

func RunHttp(repo *db.StudentsRepository) {
	p := kafka.NewProducer()
	cg := kafka.NewConsumerGroup()
	h := handlers.New(repo, p, cg)
	http.Handle("/", h.Router)

	fmt.Println("Service 9000 is listening...")

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

func RunRepo() *db.StudentsRepository {
	p := postgres.New(connectionString)
	repo := db.New(p.Pool)

	return repo
}
