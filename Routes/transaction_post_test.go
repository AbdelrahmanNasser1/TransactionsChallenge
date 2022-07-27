package Routes

import (
	"FirstWeek/Transaction"
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostAddTransaction(t *testing.T) {
	memory := Transaction.NewMemory()
	transaction := Transaction.Transaction{
		Id:        uuid.New(),
		Amount:    155,
		Currency:  "EGP",
		CreatedAt: "2022-19-7",
	}
	body, _ := json.Marshal(transaction)
	req, err := http.NewRequest("POST", "/AddTransaction", bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	handler := PostAddTransaction(memory)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Transaction has been added successfully!!"

	assert.Equal(t, expected, rr.Body.String(), "Expected not equal to Actual!!!")

}
