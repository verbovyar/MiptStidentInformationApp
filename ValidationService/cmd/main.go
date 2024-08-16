package main

import "Application/ValidationService/internal/app"

func main() {
	go app.RunHttp()
	app.RunGrpc()
}
