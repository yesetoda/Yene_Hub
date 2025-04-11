package mock_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewMockDB creates a new mock database connection
func NewMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, error) {
	// Create SQL mock database
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create mock database: %v", err)
	}

	// Configure GORM logger for testing
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)

	// Create GORM DB instance with mock database
	dialector := postgres.New(postgres.Config{
		Conn: sqlDB,
	})
	
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open gorm db: %v", err)
	}

	return db, mock, nil
}

// CloseDB closes the mock database connection
func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// ResetMock resets all mock expectations and call history
func ResetMock(mock sqlmock.Sqlmock) {
	mock.ExpectationsWereMet()
}

// MockDBConn provides a mock database connection for testing
type MockDBConn struct {
	DB   *gorm.DB
	Mock sqlmock.Sqlmock
}

// NewMockDBConn creates a new MockDBConn instance
func NewMockDBConn(t *testing.T) (*MockDBConn, error) {
	db, mock, err := NewMockDB(t)
	if err != nil {
		return nil, err
	}

	return &MockDBConn{
		DB:   db,
		Mock: mock,
	}, nil
}

// Close closes the mock database connection
func (m *MockDBConn) Close() error {
	return CloseDB(m.DB)
}

// GetDB returns the mock GORM DB instance
func (m *MockDBConn) GetDB() *gorm.DB {
	return m.DB
}

// GetMock returns the SQL mock instance
func (m *MockDBConn) GetMock() sqlmock.Sqlmock {
	return m.Mock
}

// GetSQLDB returns the underlying SQL DB instance
func (m *MockDBConn) GetSQLDB() (*sql.DB, error) {
	return m.DB.DB()
}
