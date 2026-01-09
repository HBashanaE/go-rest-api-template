package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Postgres driver
)

// Service represents a service that interacts with a database.
type Service interface {
	Health() string
    // Add other database methods here (e.g., FindUser, CreatePost)
}

type service struct {
	db *sql.DB
}

var (
	dburl      = os.Getenv("DB_URL")
	dbInstance *service
)

func New() Service {
	// Re-use connection if it exists
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("pgx", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or similar
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection
func (s *service) Health() string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := s.db.PingContext(ctx); err != nil {
		return fmt.Sprintf("db down: %v", err)
	}
	return "healthy"
}