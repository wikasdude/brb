package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	SSLMode    string
}

func Connect(cfg *Config) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("Unable to ping DB:", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	return db
}
func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal("Failed to close DB connection:", err)
	}
	log.Println("✅ Closed PostgreSQL connection")
}
func Migrate(db *sql.DB) {
	// Create the vendors table
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS vendors (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		phone VARCHAR(20) NOT NULL
	);
	`)
	if err != nil {
		log.Fatal("Failed to migrate DB:", err)
	}
	log.Println("✅ Migrated PostgreSQL")
}
