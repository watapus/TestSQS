package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateMessageFromFIS(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		json.NewEncoder(w).Encode("OKOK")
	}

	template := "transaction_processed"

	for i := 0; i < 10; i++ {
		correlationID := fmt.Sprintf("00%v", i)
		borrowerName := fmt.Sprintf("Test Borrower %v", i)
		CreateMessage(template, correlationID, borrowerName)
	}

	w.WriteHeader(http.StatusOK)
}
