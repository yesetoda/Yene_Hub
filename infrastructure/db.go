package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"a2sv.org/hub/Domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewDBConnection creates a new PostgreSQL database connection
func NewDBConnection() (*gorm.DB, error) {
	// Get database connection string from environment variable
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	// Configure GORM logger
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Error, // Only log errors
			IgnoreRecordNotFoundError: true,        // Ignore record not found errors
			Colorful:                  false,        // Disable colors for better performance
			SlowThreshold:            time.Second,  // Only log queries that take more than 1 second
		},
	)

	// Open database connection with optimized configuration
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
		PrepareStmt: true,                // Enable prepared statement cache
		SkipDefaultTransaction: true,     // Skip default transaction for better performance
	})
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	
	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)           // Maximum number of idle connections
	sqlDB.SetMaxOpenConns(100)          // Maximum number of open connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Maximum lifetime of a connection

	// Auto migrate database models
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.Country{},
		&entity.Group{},
		&entity.Track{},
		&entity.Problem{},
		&entity.Attendance{},
		&entity.Contest{},
		&entity.SuperGroup{},
		&entity.Session{},
		&entity.Submission{},
		&entity.Comment{},
		&entity.Vote{},
		&entity.Post{},
		&entity.PostToTag{},
		&entity.PostTag{},
		&entity.Invite{},
		&entity.SuperToGroup{},
		&entity.DailyProblem{},
		&entity.Exercise{},
		&entity.ProblemTrack{},
		&entity.GoogleOAuth{},
		&entity.GroupSession{},
		&entity.HOA{},
		&entity.Fund{},
		&entity.Stipend{},
		&entity.RecentAction{},
		&entity.APIToken{},
	)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
