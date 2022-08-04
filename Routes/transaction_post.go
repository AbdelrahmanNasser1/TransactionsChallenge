package Routes

import (
	"FirstWeek/Internal/Models"
	"FirstWeek/Transaction"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
	"os"
)

func PostAddTransaction(memory *Transaction.Memory) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		reqBody, _ := ioutil.ReadAll(request.Body)
		var transaction Transaction.Transaction
		json.Unmarshal(reqBody, &transaction)
		validate := validator.New()
		err := validate.Struct(transaction)
		if err != nil {
			fmt.Fprintf(writer, err.Error())
			return
		} else {
			arr := memory.GetAll()
			for _, e := range arr {
				if e.Id == transaction.Id {
					fmt.Fprint(writer, "Enter another UUID!!")
					return
				}
			}
		}
		memory.Add(transaction)
		writer.Write([]byte("Transaction has been added successfully!!"))
		return
	}
}

func PostAddTransactionToDB() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ConnectionString := os.Getenv("DATABASE_URL")
		db, _ := Models.ConnectDB(ConnectionString)

		reqBody, _ := ioutil.ReadAll(request.Body)
		var transaction Transaction.Transaction
		json.Unmarshal(reqBody, &transaction)
		validate := validator.New()
		err := validate.Struct(transaction)
		if err != nil {
			fmt.Fprintf(writer, err.Error())
			return
		} else {
			var Transactions []Models.TransactionModel
			db.NewSelect().Model(&Transactions).Scan(context.Background())

			for _, e := range Transactions {
				if e.ID == transaction.Id {
					fmt.Fprint(writer, "Enter another UUID!!")
					return
				}
			}
		}
		model := Models.TransactionModel{
			ID:        transaction.Id,
			Amount:    int64(transaction.Amount),
			Currency:  transaction.Currency,
			CreatedAt: transaction.CreatedAt,
		}
		db.NewInsert().Model(&model).Exec(context.Background())
		writer.Write([]byte("Transaction has been added successfully!!"))
		return
	}
}
