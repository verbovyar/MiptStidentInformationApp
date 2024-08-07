package main

import (
	"StudentsInfoService/internal/app"
)

func main() {
	repo := app.RunRepo()
	app.RunHttp(repo)
}
