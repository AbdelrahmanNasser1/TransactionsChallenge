package Transaction

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Transaction struct {
	Id        uuid.UUID `json:"id" validate:"required"`
	Amount    float64   `json:"amount" validate:"gte=0,lte=100000"`
	Currency  string    `json:"currency" validate:"len=3"`
	CreatedAt string    `json:"createdAt"`
	Status    bool      `json:"status"`
}

type Getter interface {
	GetAll() []Transaction
}
type Setter interface {
	Add(transaction Transaction)
}

type Memory struct {
	Transactions []Transaction
}

func NewMemory() *Memory {
	return &Memory{
		Transactions: []Transaction{},
	}
}

func (m *Memory) GetAll() []Transaction {
	return m.Transactions
}

func (m *Memory) Add(transaction Transaction) {
	validate := validator.New()
	err := validate.Struct(transaction)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		m.Transactions = append(m.Transactions, transaction)
	}
}
