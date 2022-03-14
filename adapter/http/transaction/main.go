package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Davi-Arauj/dio-finance/model/transaction"
	"github.com/Davi-Arauj/dio-finance/util"
)
//GetTransaction BLABLABLLA
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2022-03-14T09:34:26"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}
//CreateATransaction blablablablabla
func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll((r.Body))

	_ = json.Unmarshal(body, &res)

	fmt.Println(res)

}
