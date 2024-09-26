package db

import (
	"fmt"
	"log"
	"next_gen_job_hunting/config/env"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() {
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Los_Angeles",
		env.GetDBHost(),
		env.GetDBUser(),
		env.GetDBPassword(),
		env.GetDBName(),
		env.GetDBPort(),
	)


	customLogger := RegisterLogger()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Database connection successfully established")
}

func GetDB() *gorm.DB {
	return DB
}

func RegisterLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Info,   // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound errors
			Colorful:                  true,          // Enable color in logs
		},
	)
}

// CloseDB closes the database connection (usually for cleanup/shutdown)
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error getting database connection: %v", err)
	}
	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
}


