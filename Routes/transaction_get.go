package Routes

import (
	"FirstWeek/Transaction"
	"encoding/json"
	"net/http"
)

func GetTransactions(memory Transaction.Getter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		transactions := memory.GetAll()
		json.NewEncoder(writer).Encode(transactions)
	}
}
