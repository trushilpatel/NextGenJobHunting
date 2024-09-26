package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var envVarsMap = map[string]string{
	"DB_USERNAME": "",
	"DB_PASSWORD": "",
	"DB_HOST":     "",
	"DB_NAME":     "",
	"DB_PORT": "",
}

var envPath = "../.env"

func LoadEnvVars() {
	getAndLoadEnvVariables()
	logEnvVars()
}

func logEnvVars() {
	missingVars := []string{}

	for name, value := range envVarsMap {
		if value == "" {
			missingVars = append(missingVars, name)
		}
	}
	
	log.Println("*********** ENV VARIABLES ***********")
	if len(missingVars) > 0 {
		log.Printf("=> Total missing environment variables : ", len(missingVars))
		log.Printf("%v",missingVars)
	}

	log.Println("Loaded environment variables:")
	for name, value := range envVarsMap {
		log.Printf("%s: %s\n", name, value)
	}

	log.Println("*********** END OF ENV VARIABLES ***********")
}

func getAndLoadEnvVariables() {
	if _, err := os.Stat(envPath); err == nil {
		log.Printf("Env Path: " + envPath)
		if err := godotenv.Load(envPath); err != nil {
			log.Fatalf("Error loading .env file")
		}
	} else {
		log.Fatalf("Env File not found at " + envPath)
	}

	for key := range envVarsMap {
		envVarsMap[key] = os.Getenv(key)
	}
}

func GetDBUser() string {
	return envVarsMap["DB_USERNAME"]
}

func GetDBPassword() string {
	return envVarsMap["DB_PASSWORD"]
}

func GetDBHost() string {
	return envVarsMap["DB_HOST"]
}

func GetDBName() string {
	return envVarsMap["DB_NAME"]
}

func GetDBPort() string {
	return envVarsMap["DB_PORT"]
}