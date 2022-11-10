package db

import (
	"FirstWeek/Config"
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewDBConnection(configurations Config.Configurations) (*bun.DB, error) {
	connectionString := configurations.Database.ConnectionString
	ctx := context.Background()
	con, err := pgx.Connect(ctx, connectionString)
	defer con.Close(context.Background())
	if err != nil {
		return nil, errors.New(err.Error())
	}
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(connectionString)))
	db := bun.NewDB(sqlDb, pgdialect.New())
	return db, nil
}
