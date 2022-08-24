package api

import (
	"FirstWeek/Internal/Services"
	"FirstWeek/Transaction"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type TransactionController struct {
	TransactionService Services.IService
}

func NewTransactionController(ts Services.IService) {
	port := ":3000"
	c := &TransactionController{
		TransactionService: ts,
	}
	r := chi.NewRouter()

	r.Post("/transaction-db", c.handleCreateTransaction)
	r.Get("/transactions-db", c.handleGetTransactions)

	fmt.Println("Serving on", port)
	http.ListenAndServe(port, r)
}
func (c *TransactionController) handleCreateTransaction(w http.ResponseWriter, r *http.Request) {

	var modelTransaction Transaction.Transaction

	if err := json.NewDecoder(r.Body).Decode(&modelTransaction); err != nil {
		fmt.Fprintf(w, err.Error())
	}

	res, err := c.TransactionService.Create(context.Background(), &modelTransaction)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.NewEncoder(w).Encode(res)
}
func (c *TransactionController) handleGetTransactions(w http.ResponseWriter, r *http.Request) {

	res, err := c.TransactionService.List(context.Background())
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.NewEncoder(w).Encode(res)
}
