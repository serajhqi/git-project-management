package app

import (
	"database/sql"
	"fmt"
	"git-project-management/config"
	"git-project-management/internal/activity"
	"log"

	"github.com/danielgtaylor/huma/v2"
	_ "github.com/lib/pq"
)

func Setup(api *huma.API, cfg config.Config) {
	// init database

	db, err := sql.Open("postgres", cfg.CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")

	activity.Setup(api, db)
}
