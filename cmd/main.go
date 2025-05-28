package main

import (
	"api-contracts-check/cmd/app"
)

func main() {
	application := app.BuildApplication()
	if err := application.Run(); err != nil {
		panic(err)
	}
}
