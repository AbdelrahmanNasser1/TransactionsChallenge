package main

import (
	"FirstWeek/Internal/Repository"
	"FirstWeek/Internal/Services"
	"FirstWeek/Internal/adapter/api"
	"FirstWeek/Internal/adapter/db"
	"FirstWeek/Transaction"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	//port := ":3000"
	memory := Transaction.NewMemory()
	memory.Add(Transaction.Transaction{
		Id:        uuid.New(),
		Amount:    155,
		Currency:  "EGP",
		CreatedAt: "2022-19-7",
	})
	memory.Add(Transaction.Transaction{
		Id:        uuid.New(),
		Amount:    43687,
		Currency:  "MXN",
		CreatedAt: "2022-08-04",
	})

	//Create Routes
	/*r := chi.NewRouter()

	r.Get("/transactions", Routes.GetTransactions(memory))
	r.Get("/transactions-from-DB", Routes.GetTransactionsFromDB())
	r.Post("/transaction", Routes.PostAddTransaction(memory))
	r.Post("/transaction-to-DB", Routes.PostAddTransactionToDB())

	fmt.Println("Serving on", port)
	http.ListenAndServe(port, r)*/
	conn, err := db.NewDBConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	tranRepo := Repository.NewDefaultRepository(conn)
	tranSer := Services.NewDefaultService(tranRepo)
	api.NewTransactionController(tranSer)
}
