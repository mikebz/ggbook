package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Go Guest Book Starting")
	startFlag := flag.Bool("start", true, "Start web server")
	migrateFlag := flag.Bool("migrate", false, "Database migration")

	flag.Parse()

	err := openDb()
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
		listenAndServe()
		return
	}
}
