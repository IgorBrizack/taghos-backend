package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	conn *gorm.DB
}

var (
	instance *Database
	once     sync.Once
)

func NewDatabase() (*Database, error) {
	var err error
	once.Do(func() {
		dbHost := os.Getenv("DB_PROD_HOST")
		dbPort := os.Getenv("DB_PROD_PORT")
		dbName := os.Getenv("DB_PROD_NAME")
		dbUser := os.Getenv("DB_PROD_USER")
		dbPass := os.Getenv("DB_PROD_PASS")

		if dbHost == "" || dbPort == "" || dbName == "" || dbUser == "" || dbPass == "" {
			err = fmt.Errorf("missing required environment variables for database configuration")
			return
		}

		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPass, dbName)

		conn, connErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if connErr != nil {
			err = fmt.Errorf("failed to connect to database: %w", connErr)
			return
		}

		sqlDB, connErr := conn.DB()
		if connErr != nil {
			err = fmt.Errorf("failed to get sql.DB from GORM: %w", connErr)
			return
		}

		if pingErr := sqlDB.Ping(); pingErr != nil {
			err = fmt.Errorf("failed to ping database: %w", pingErr)
			return
		}

		log.Println("Successfully connected to the database using GORM")
		instance = &Database{conn: conn}
	})

	return instance, err
}

func GetInstance() (*Database, error) {
	if instance == nil {
		return NewDatabase()
	}
	return instance, nil
}

func (db *Database) GetConnection() *gorm.DB {
	return db.conn
}

func (db *Database) ValidateConnection() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from GORM: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Println("Database connection lost. Reconnecting...")
		newDB, err := NewDatabase()
		if err != nil {
			return fmt.Errorf("failed to reconnect to database: %w", err)
		}
		*db = *newDB
	}
	return nil
}

func (db *Database) Close() error {
	sqlDB, err := db.conn.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from GORM: %w", err)
	}

	if sqlDB != nil {
		log.Println("Closing database connection")
		return sqlDB.Close()
	}
	return nil
}
