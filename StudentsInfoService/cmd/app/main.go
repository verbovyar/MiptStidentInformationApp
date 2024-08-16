package main

import (
	"StudentsInfoService/internal/app"
)

func main() {
	repo := app.RunRepo()
	app.RunKafka(repo)
	go app.RunHttp(repo)
	app.RunGrpc()
}
