```go
package database

import (
	"batchProcess/config"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var connection *sql.DB

func GetInstance() *sql.DB {
	return connection
}

func SetInstance(conn *sql.DB) {
	connection = conn
}

func InitializeDB(cfg config.SConfig) error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DbName,
		cfg.Database.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	timeout := 5 * time.Second
	done := make(chan error, 1)

	go func() {
		done <- db.Ping()
	}()

	select {
	case <-time.After(timeout):
		return fmt.Errorf("timeout exceeded while connecting to the database")
	case err := <-done:
		if err != nil {
			return fmt.Errorf("failed to verify database status: %w", err)
		}
	}

	log.Println("database connection successfully initialized")
	connection = db
	return nil
}

```