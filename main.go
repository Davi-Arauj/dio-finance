package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/transactions", getTransaction)
	http.HandleFunc("/transactions/create", createATransaction)

	http.ListenAndServe(":8080", nil)
}

type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type Transactions []Transaction

func getTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2006-01-02T15:04:05")

	var transactions = Transactions{
		Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func createATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = Transactions{}
	var body,_ = ioutil.ReadAll((r.Body))

	_ = json.Unmarshal(body,&res)

	fmt.Println(res)

}