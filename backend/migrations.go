package main

import (
	"log"

	"next-gen-job-hunting/api/jobs"
	"next-gen-job-hunting/api/user"

	"next-gen-job-hunting/config/database"
)

var models = []interface{}{&user.User{}, &jobs.JobUrl{}}

func RunAutoDBMigrations() {
	db := database.NewDB()
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("failed to auto migrate %T: %v", model, err)
		}
	}
}
