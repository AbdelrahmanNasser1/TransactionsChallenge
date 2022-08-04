package Routes

import (
	"FirstWeek/Internal/Models"
	"FirstWeek/Transaction"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetTransactions(memory Transaction.Getter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		transactions := memory.GetAll()
		json.NewEncoder(writer).Encode(transactions)
	}
}

func GetTransactionsFromDB() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ConnectionString := os.Getenv("DATABASE_URL")
		db, err := Models.ConnectDB(ConnectionString)
		if err != nil {
			fmt.Fprintf(writer, err.Error())
		}
		var Transactions []Models.TransactionModel
		db.NewSelect().Model(&Transactions).Scan(context.Background())
		json.NewEncoder(writer).Encode(Transactions)
	}
}
