package server

import (
	"github.com/gorilla/mux"
	"github.com/hjoshi123/temporal-loan-app/internal/handler"
	v1 "github.com/hjoshi123/temporal-loan-app/pkg/controllers/v1"
)

func initializeRoutes(r *mux.Router) {
	v1Router := r.PathPrefix("/v1").Subrouter()

	// Banks
	banksController := v1.NewV1BanksController()
	bankRoutes := v1Router.PathPrefix("/banks").Subrouter()
	bankRoutes.Handle("/create", handler.CustomHandler(banksController.CreateBank)).Methods("POST")

	// Transactions
	transactionsController := v1.NewV1TransactionsController()
	transactionRoutes := v1Router.PathPrefix("/transactions").Subrouter()
	transactionRoutes.Handle("/start", handler.CustomHandler(transactionsController.StartTransaction)).Methods("POST")
	transactionRoutes.Handle("/approve", handler.CustomHandler(transactionsController.ApproveTransaction)).Methods("POST")
	transactionRoutes.Handle("/reject", handler.CustomHandler(transactionsController.RejectTransaction)).Methods("POST")

	// Accounts
	accountsController := v1.NewV1AccountsController()
	accountRoutes := v1Router.PathPrefix("/accounts").Subrouter()
	accountRoutes.Handle("/create", handler.CustomHandler(accountsController.CreateAccount)).Methods("POST")
}

func Setup() *mux.Router {
	r := mux.NewRouter()

	initializeRoutes(r)
	return r
}
