package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArchishmanSengupta/expense-tracker/api/models"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	transactionInstance := models.Transaction{}

	json.NewDecoder(r.Body).Decode(&transactionInstance)

	// transactionInstance.CreatedAt = time.Now()
	// transactionInstance.UpdatedAt = time.Now()

	transaction, err := transactionInstance.Insert()

	//Error Handling
	if err != nil {
		fmt.Println("Error found while", err)
	}
	//content type
	w.Header().Set("Content-Type", "application/json")

	//encode
	json.NewEncoder(w).Encode(transaction)

}
