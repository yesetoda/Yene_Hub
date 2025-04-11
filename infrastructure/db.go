package infrastructure

import (
	"fmt"
	"log"
	"os"

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
			LogLevel: logger.Info,
		},
	)

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

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
