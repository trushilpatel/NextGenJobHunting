package database

import (
	"log"
	"next-gen-job-hunting/config/env"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func connectDB() (*gorm.DB, error) {
	dsn := env.GetDBConnectionURL()
	customLogger := RegisterLogger()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: customLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}
	log.Println("Database connection successfully established")

	return db, nil
}

func NewDB() *gorm.DB {
	once.Do(func() {
		var err error
		DB, err = connectDB()
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
	})
	return DB
}

func RegisterLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level (Silent, Error, Warn, Info)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound errors
			Colorful:                  true,        // Enable color in logs
		},
	)
}

// CloseDB closes the database connection (usually for cleanup/shutdown)
func CloseDB() {
	if DB == nil {
		return
	}
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Error getting database connection: %v", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}
