package db

import (
	"context"
	"database/sql"
	"time"
)

func NewUserDB(dbAddr string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dbAddr)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
