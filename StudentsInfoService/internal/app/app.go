package app

import (
	"StudentsInfoService/api/api/ServiceApiPb"
	"StudentsInfoService/internal/grpcHandlers"
	"StudentsInfoService/internal/handlers"
	kafka2 "StudentsInfoService/internal/kafka"
	"StudentsInfoService/internal/kafkaHandlers/addHandler"
	"StudentsInfoService/internal/repositories/db"
	cache2 "StudentsInfoService/pkg/cache"
	"StudentsInfoService/pkg/postgres"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

const port = ":9000"
const connectionString = "postgres://postgres:Verbov3232132121@localhost:5432/studentsDB"

func RunHttp(repo *db.StudentsRepository) {
	h := handlers.New(repo)
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

func RunKafka(repo *db.StudentsRepository) {
	cache, _ := cache2.New()
	producer := kafka2.NewProducer()
	consumerGroup := kafka2.NewConsumerGroup()
	Add := addHandler.NewAddHandler(producer, repo, consumerGroup, cache)
	go addHandler.AddClaim(Add)
}

func RunGrpc(repo *db.StudentsRepository) {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Printf("listen error %v\n", err)
	}
	h := grpcHandlers.New(repo)
	ServiceApiPb.RegisterStudentsInfoServiceServer(grpcServer, h)
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Printf("serve error%v\n", err)
	}
}
