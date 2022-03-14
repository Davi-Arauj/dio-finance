package http

import (
	"net/http"

	"github.com/Davi-Arauj/dio-finance/adapter/http/actuator"
	"github.com/Davi-Arauj/dio-finance/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Init bçalsjflakp
func Init() {
	http.HandleFunc("/transactions", transaction.GetTransaction)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)
	http.HandleFunc("/health", actuator.Health)
	http.Handle("/metrics", promhttp.Handler())


	http.ListenAndServe(":8080", nil)

}
