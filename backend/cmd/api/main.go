package main

import (
	"gonote/internal/app"
	"gonote/internal/config"
	"gonote/internal/db"
	"log"
)

func main() {
	// err := godotenv.Load()
	//
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	cfg := config.NewConfig()

	DB, err := db.New(&cfg.DB)

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	appContext := app.NewAppContext(DB)

	a := app.NewApplication(cfg, appContext)

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
