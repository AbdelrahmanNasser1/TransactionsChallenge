package Routes

import (
	"FirstWeek/Transaction"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
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
