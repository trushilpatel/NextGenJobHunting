package database

import (
	"log"
	"next-gen-job-hunting/jobs"
	"next-gen-job-hunting/user"
)

var models = []interface{}{&user.User{}, &jobs.JobUrl{}}

func RunAutoDBMigrations() {
	db := GetDB()
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("failed to auto migrate %T: %v", model, err)
		}
	}
}
