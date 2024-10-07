package main

import (
	"log"

	"github.com/mahinops/secretcli-web/cmd"
)

// Main entry point
func main() {
	app, err := cmd.NewApp(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer app.CloseDatabase()

	// Start the server
	app.StartServer(":8080")
}
