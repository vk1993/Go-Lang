package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const creditScoreMin = 500
const creditScoreMax = 900

type credit_rating struct {
	CreditRating int `json:"credit_rating"`
}

func getCreditScore(w http.ResponseWriter, r *http.Request) {
	// Introduce a slight delay to simulate processing time
	time.Sleep(time.Millisecond * 300) // Adjust delay as needed
	var creditRating = credit_rating{
		CreditRating: (rand.Intn(creditScoreMax-creditScoreMin) + creditScoreMin),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(creditRating)
}

func handleRequests() {
	http.Handle("GET /stock", http.HandlerFunc(getCreditScore))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
