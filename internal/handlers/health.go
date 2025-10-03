package handlers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	resp := response{
		Message: "OK",
		Status:  http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)
	json.NewEncoder(w).Encode(resp)
}
