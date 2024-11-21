package config

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

func InitDB(host, port, user, password, dbname string) *sql.DB {
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
        host, port, user, password, dbname)
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    return db
}