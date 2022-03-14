package actuator

import (
	"encoding/json"
	"net/http"
)
//HealthBody  blabalbla
type HealthBody struct {
	Status string `json:"status"`
}
//Health blabla
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var status = HealthBody{"alive"}

	_ = json.NewEncoder(w).Encode(status)
}
