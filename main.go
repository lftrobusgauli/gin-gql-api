package main

import (
	"log"
	"thegame/pkg/db"

	"thegame/pkg/app"
	"thegame/pkg/config"
)

func main() {
	// Create config instance and load the config from the environment
	appconfig := config.New()
	if err := appconfig.LoadFromEnv(); err != nil {
		log.Fatalf("could not load env variables into config: %v", err)
	}

	// Create a new database instance
	database, err := db.Init(appconfig.DBConfig)
	if err != nil {
		log.Fatalf("could not initialize database: %v", err)
	}

	// Load the application dependencies.
	a := &app.App{
		Storage: database,
	}
	// Run the application.
	if err := a.Run(); err != nil {
		log.Fatalf("something went wrong while running the app: %v", err)
	}
}
