package Repository

import (
	"FirstWeek/Internal/Models"
	"context"
	"github.com/uptrace/bun"
)

type databaseRepo struct {
	conn *bun.DB
}

func NewDefaultRepository(conn *bun.DB) *databaseRepo {
	return &databaseRepo{
		conn: conn,
	}
}

func (db *databaseRepo) List(ctx context.Context) ([]Models.TransactionModel, error) {
	var models []Models.TransactionModel
	err := db.conn.NewSelect().Model(&models).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (db *databaseRepo) Create(ctx context.Context, model *Models.TransactionModel) error {
	_, err := db.conn.NewInsert().Model(model).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
