// main package for the go guest book application
// this application demonstrates a simple golang web application
// using sqlite as the database and gorm as the orm.
// the application has a simple api for creating, reading, updating and deleting resources.
// the application is designed to be easy to deploy and run.  It uses a single binary
// and has minimal dependencies.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Go Guest Book Starting")
	startFlag := flag.Bool("start", true, "Start web server")
	migrateFlag := flag.Bool("migrate", false, "Database migration")

	flag.Parse()

	// load up all the necessary environment variables
	dbUrlEnv := os.Getenv("DB_URL")
	server := os.Getenv("SERVER")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := openDb(dbUrlEnv)
	if err != nil {
		panic(err)
	}

	// the actions below to migrate and serve
	// are mutually exclusive on purpose.
	// generally migration is done one time and we don't wnat to mix it
	// with running the regular service

	if *migrateFlag {
		err = migrate()
		if err != nil {
			panic(err)
		}
		return
	}

	if *startFlag {
		listenAndServe(server + ":" + port)
		return
	}
}
