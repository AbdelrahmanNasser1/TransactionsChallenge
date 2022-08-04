package main

import (
	"FirstWeek/Routes"
	"FirstWeek/Transaction"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	port := ":3000"
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
	r := chi.NewRouter()

	r.Get("/transactions", Routes.GetTransactions(memory))
	r.Get("/transactionsFromDB", Routes.GetTransactionsFromDB())
	r.Post("/AddTransaction", Routes.PostAddTransaction(memory))
	r.Post("/AddTransactionToDB", Routes.PostAddTransactionToDB())

	fmt.Println("Serving on", port)
	http.ListenAndServe(port, r)
}
