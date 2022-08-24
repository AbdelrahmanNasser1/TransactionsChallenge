package Services

import (
	"FirstWeek/Transaction"
	"context"
)

type IService interface {
	Create(ctx context.Context, transaction *Transaction.Transaction) (*Transaction.Transaction, error)
	List(ctx context.Context) ([]Transaction.Transaction, error)
}
