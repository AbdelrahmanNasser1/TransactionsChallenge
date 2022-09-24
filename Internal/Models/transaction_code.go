package Models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type TransactionModel struct {
	bun.BaseModel `bun:"table:transaction"`
	ID            uuid.UUID `bun:"id,notnull,pk,type:uuid,default:gen_random_uuid()"`
	Amount        int64     `bun:"amount,notnull"`
	Currency      string    `bun:"currency,notnull"`
	CreatedAt     string    `bun:"createdat,notnull"`
	Status        bool      `bun:"status"`
}

func ConnectDB(connectionString string) (*bun.DB, error) {
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
