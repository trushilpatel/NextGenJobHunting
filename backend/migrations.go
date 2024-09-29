package main

import (
	"log"

	"next-gen-job-hunting/api/joburl"
	"next-gen-job-hunting/api/token"
	"next-gen-job-hunting/api/user"

	"next-gen-job-hunting/config/database"
)

var models = []interface{}{
	&user.User{},
	&joburl.JobUrl{},
	&token.Token{},
}

func RunAutoDBMigrations() {
	db := database.NewDB()
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("failed to auto migrate %T: %v", model, err)
		}
	}
}
