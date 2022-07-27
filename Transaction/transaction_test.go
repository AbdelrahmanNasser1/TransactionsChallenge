package Transaction

import (
	"github.com/google/uuid"
	"testing"
)

func TestAdd(t *testing.T) {
	memory := NewMemory()
	memory.Add(Transaction{
		Id:        uuid.New(),
		Amount:    155,
		Currency:  "EGP",
		CreatedAt: "2022-19-7",
	})
	if len(memory.Transactions) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	memory := NewMemory()
	memory.Add(Transaction{
		Id:        uuid.New(),
		Amount:    155,
		Currency:  "EGP",
		CreatedAt: "2022-19-7"})
	results := memory.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
