package Routes

import (
	"FirstWeek/Internal/Models"
	"FirstWeek/Transaction"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetTransactions(t *testing.T) {
	memory := Transaction.NewMemory()
	memory.Add(Transaction.Transaction{
		Id:        uuid.New(),
		Amount:    155,
		Currency:  "EGP",
		CreatedAt: "2022-19-7",
	})

	req, err := http.NewRequest("GET", "/transactions", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	handler := GetTransactions(memory)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	Body, _ := ioutil.ReadAll(rr.Body)
	BodyString := strings.TrimSpace(string(Body))

	MemoryString, _ := json.Marshal(memory.Transactions)

	assert.Equal(t, BodyString, string(MemoryString), "Expected not equal to Actual!!!")

}
func TestGetTransactionsFromDB(t *testing.T) {
	ConnectionString := os.Getenv("DATABASE_URL")
	db, err := Models.ConnectDB(ConnectionString)

	if err != nil {
		t.Error(err)
	}
	var Transactions []Models.TransactionModel
	db.NewSelect().Model(&Transactions).Scan(context.Background())

	req, err := http.NewRequest("GET", "/transaction", nil)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	handler := GetTransactionsFromDB()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	Body, _ := ioutil.ReadAll(rr.Body)
	BodyString := strings.TrimSpace(string(Body))
	DBString, _ := json.Marshal(Transactions)

	assert.Equal(t, BodyString, string(DBString), "Expected not equal to Actual!!!")

}
