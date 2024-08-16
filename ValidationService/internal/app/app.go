package app

import (
	"Application/ValidationService/api/api/ValidationApiPb"
	"Application/ValidationService/internal/grpcHandlers"
	"Application/ValidationService/internal/handlers"
	"Application/ValidationService/kafka"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
)

const port = ":3030"

func RunHttp() {
	p := kafka.NewProducer()
	c := kafka.NewConsumer()
	h := handlers.New(p, c)
	http.Handle("/", h.Router)

	fmt.Println("Service 3030 is listening...")

	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

func RunGrpc() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("listen error %v\n", err)
	}
	connect, err := grpc.Dial("9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := ValidationApiPb.NewValidationServiceClient(connect)
	h := grpcHandlers.New(client)
	ValidationApiPb.RegisterValidationServiceServer(grpcServer, h)
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Printf("serve error%v\n", err)
	}
}
