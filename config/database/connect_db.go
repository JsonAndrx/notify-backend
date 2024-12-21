package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"notify-backend/api/utils/debug"

	_ "github.com/lib/pq"
)

var (
    db   *sql.DB
    once sync.Once
)

func GetDBConnection() (*sql.DB, error) {
    var err error
    once.Do(func() {
        fmt.Println("Connecting to database...")
        connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
            os.Getenv("DB_HOST"),
            os.Getenv("DB_PORT"),
            os.Getenv("DB_USER"),
            os.Getenv("DB_PASS"),
            os.Getenv("DB_NAME"))

        db, err = sql.Open("postgres", connStr)
        if err != nil {
            debug.LogError(err)
            return
        }

        db.SetMaxOpenConns(25)
        db.SetMaxIdleConns(25)
        db.SetConnMaxLifetime(5 * 60)

        err = db.Ping()
        if err != nil {
            debug.LogError(err)
            return
        }
    })

    if err != nil {
        return nil, err
    }

    return db, nil
}