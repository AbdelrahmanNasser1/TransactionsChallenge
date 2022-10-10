package Services

import (
	"FirstWeek/Internal/Models"
	"FirstWeek/Transaction"
	"context"
)

type IService interface {
	Create(ctx context.Context, transaction *Transaction.Transaction) (*Transaction.Transaction, error)
	List(ctx context.Context) ([]Transaction.Transaction, error)
	Update(ctx context.Context, model Models.TransactionModel) ([]Transaction.Transaction, error)
}
