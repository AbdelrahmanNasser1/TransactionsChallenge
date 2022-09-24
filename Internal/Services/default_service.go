package Services

import (
	"FirstWeek/Internal/Models"
	"FirstWeek/Internal/Repository"
	"FirstWeek/Transaction"
	"context"
	"github.com/go-playground/validator/v10"
)

type DefaultTransactionService struct {
	transactionRepo Repository.IRepository
}

func NewDefaultService(repo Repository.IRepository) *DefaultTransactionService {
	return &DefaultTransactionService{
		transactionRepo: repo,
	}
}

func (s *DefaultTransactionService) Create(ctx context.Context, transaction *Transaction.Transaction) (*Transaction.Transaction, error) {
	validate := validator.New()
	err := validate.Struct(transaction)
	if err != nil {
		return nil, err
	}
	model := Models.TransactionModel{
		ID:        transaction.Id,
		Amount:    int64(transaction.Amount),
		Currency:  transaction.Currency,
		CreatedAt: transaction.CreatedAt,
		Status:    transaction.Status,
	}
	if err := s.transactionRepo.Create(ctx, &model); err != nil {
		return nil, err
	}
	res := Transaction.Transaction{
		Id:        model.ID,
		Amount:    float64(model.Amount),
		Currency:  model.Currency,
		CreatedAt: model.CreatedAt,
		Status:    model.Status,
	}
	return &res, nil
}

func (s *DefaultTransactionService) List(ctx context.Context) ([]Transaction.Transaction, error) {
	models, err := s.transactionRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	var temp Transaction.Transaction
	var result []Transaction.Transaction
	for _, e := range models {
		temp.Id = e.ID
		temp.Amount = float64(e.Amount)
		temp.Currency = e.Currency
		temp.CreatedAt = e.CreatedAt
		temp.Status = e.Status
		result = append(result, temp)
	}
	return result, nil
}
