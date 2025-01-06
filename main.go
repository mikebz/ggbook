// main package for the go guest book application
// this application demonstrates a simple golang web application
// using sqlite as the database and gorm as the orm.
// the application has a simple api for creating, reading, updating and deleting resources.
// the application is designed to be easy to deploy and run.  It uses a single binary
// and has minimal dependencies.
package main

import (
	"context"
	"flag"
	"log"
	"os"
)

var logger = log.Default()

func main() {
	ctx := context.Background()
	logger.Println("Go Guest Book Starting")
	startFlag := flag.Bool("start", true, "Start web server")
	migrateFlag := flag.Bool("migrate", false, "Database migration")

	flag.Parse()

	// load up all the necessary environment variables
	dbUrlEnv := os.Getenv("DB_URL")
	aiApiKey := os.Getenv("GEMINI_API_KEY")
	server := os.Getenv("SERVER")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := openDb(dbUrlEnv)
	if err != nil {
		log.Fatal(err)
	}

	// the actions below to migrate and serve
	// are mutually exclusive on purpose.
	// generally migration is done one time and we don't wnat to mix it
	// with running the regular service

	if *migrateFlag {
		err = migrate()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if *startFlag {

		// setting up the AI related objects.
		client, err := createAiClient(ctx, aiApiKey)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		configureAiModel(client)

		// now starting the web server
		listenAndServe(server + ":" + port)
		return
	}
}
