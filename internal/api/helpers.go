package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpError struct {
	Error error `json:"error"`
}

func SendErrJSON(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	err = json.NewEncoder(w).Encode(HttpError{Error: err})
	if err != nil {
		log.Println(err)
	}
	//	todo add errors mapper
}

func SendJSON(w http.ResponseWriter, resp any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println(err)
	}
}
