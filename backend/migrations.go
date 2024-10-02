package main

import (
	"log"

	"next-gen-job-hunting/api/joburl"
	"next-gen-job-hunting/api/token"
	"next-gen-job-hunting/api/user"
	userjobpost "next-gen-job-hunting/api/user-job-post"

	"next-gen-job-hunting/api/jobpost"
	"next-gen-job-hunting/config/database"
)

var models = []interface{}{
	&user.User{},
	&joburl.JobUrl{},
	&jobpost.JobPost{},
	&token.Token{},
	&userjobpost.UserJobPost{},
}

var sqlScriptsList = [][]string{
	userjobpost.UserJobPostScripts,
	jobpost.SqlScripts,
}

func RunAutoDBMigrations() {
	var sqlScripts = []string{}
	for _, scripts := range sqlScriptsList {
		sqlScripts = append(sqlScripts, scripts...)
	}

	db := database.NewDB()

	// Execute SQL scripts to create enums
	for _, script := range sqlScripts {
		if err := db.Exec(script).Error; err != nil {
			log.Fatalf("failed to execute script: %v", err)
		}
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("failed to auto migrate %T: %v", model, err)
		}
	}

}
