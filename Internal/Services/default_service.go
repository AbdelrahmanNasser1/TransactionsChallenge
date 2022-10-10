package Services

import (
	"FirstWeek/Internal/Models"
	"FirstWeek/Internal/Repository"
	"FirstWeek/Transaction"
	"context"
	"github.com/dranikpg/dto-mapper"
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
	model := &Models.TransactionModel{}
	dto.Map(model, transaction)

	if err := s.transactionRepo.Create(ctx, model); err != nil {
		return nil, err
	}
	res := &Transaction.Transaction{}
	dto.Map(&res, model)
	return res, nil
}

func (s *DefaultTransactionService) List(ctx context.Context) ([]Transaction.Transaction, error) {
	models, err := s.transactionRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	var result = &[]Transaction.Transaction{}
	dto.Map(result, models)
	return *result, nil
}
func (s *DefaultTransactionService) Update(ctx context.Context, model Models.TransactionModel) ([]Transaction.Transaction, error) {

	err := s.transactionRepo.Update(ctx, &model)
	if err != nil {
		return nil, err
	}
	var result = &[]Transaction.Transaction{}
	dto.Map(result, model)
	return *result, nil
}
