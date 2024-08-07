package app

import (
	"Application/ValidationService/internal/handlers"
	"Application/ValidationService/kafka"
	"fmt"
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
