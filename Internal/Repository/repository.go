package Repository

import (
	"FirstWeek/Internal/Models"
	"context"
)

type IRepository interface {
	List(ctx context.Context) ([]Models.TransactionModel, error)
	Create(ctx context.Context, transaction *Models.TransactionModel) error
}
