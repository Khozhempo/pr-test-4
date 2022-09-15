package main

import (
	"log"
	"pr-test-4/server"
)

func main() {
	app := server.NewApp()

	if err := app.Run("8080"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
