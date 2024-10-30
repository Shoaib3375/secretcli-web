// main.go

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mahinops/secretcli-web/cmd"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please specify 'api' or 'web' as a command-line argument")
	}

	// Initialize the application based on the mode
	mode := os.Args[1]
	app, err := cmd.NewApp(".env", mode)
	if err != nil {
		log.Fatal(err)
	}
	defer app.CloseDatabase()

	// Determine whether to run the API or web server based on command-line argument
	switch mode {
	case "api":
		startAPI(app)
	case "web":
		startWeb(app)
	default:
		log.Fatalf("Unknown command: %s. Use 'api' or 'web'.", mode)
	}
}

// startAPI launches the API server
func startAPI(app *cmd.App) {
	fmt.Println("Starting API server on :8080")
	app.StartAPIServer(":8080") // Adjust port if needed
}

// startWeb launches the web server with template rendering
func startWeb(app *cmd.App) {
	fmt.Println("Starting Web server on :8081")
	app.StartWebServer(":8081") // Adjust port if needed
}
