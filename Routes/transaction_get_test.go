package Routes

import (
	"FirstWeek/Transaction"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
