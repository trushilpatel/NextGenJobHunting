package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm/logger"
)

var envVarsMap = map[string]string{
	"PORT": "",

	"GIN_MODE":       "debug",
	"GORM_LOG_LEVEL": "warn",

	"DB_USERNAME": "",
	"DB_PASSWORD": "",
	"DB_HOST":     "",
	"DB_NAME":     "",
	"DB_PORT":     "",

	"LOG_ENV_VARIABLES": "false",

	"JWT_SECRET": "",
}

var envPath = "../.env"

func LoadEnvVars() {
	getAndLoadEnvVariables()
	if LogEnvVariables() {
		logEnvVars()
	}
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
		log.Printf("%v", missingVars)
	}

	log.Println("Loaded environment variables:")
	for name, value := range envVarsMap {
		log.Printf("%s: %s\n", name, value)
	}

	log.Println("*********** END OF ENV VARIABLES ***********")
}

func getAndLoadEnvVariables() {
	if _, err := os.Stat(envPath); err == nil {
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

func GetPort() string {
	return envVarsMap["PORT"]
}

func GetGinMode() string {

	return envVarsMap["GIN_MODE"]
}

func GetGormLogLevel() logger.LogLevel {
	switch envVarsMap["GIN_LOG_LEVEL"] {
	case "info":
		return logger.Info
	case "warn":
		return logger.Warn
	case "error":
		return logger.Error
	case "silent":
		return logger.Silent
	default:
		return logger.Warn
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

func LogEnvVariables() bool {
	return envVarsMap["LOG_ENV_VARIABLES"] == "true"
}

func GetDBConnectionURL() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Los_Angeles",
		GetDBHost(),
		GetDBUser(),
		GetDBPassword(),
		GetDBName(),
		GetDBPort(),
	)
}

func GetJWTSecret() string {
	return envVarsMap["JWT_SECRET"]
}
