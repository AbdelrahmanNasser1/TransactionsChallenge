package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

func NewDBConnection() (*bun.DB, error) {
	connectionString := os.Getenv("DATABASE_URL")
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connectionString)
	defer conn.Close(context.Background())
	if err != nil {
		return nil, errors.New(err.Error())
	}
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))
	db := bun.NewDB(sqlDb, pgdialect.New())
	return db, nil
}
