package http

import (
	"encoding/json"
	"net/http"

	"github.com/echovl/pinger"
	"github.com/gorilla/mux"
)

func Handler(core *pinger.Core) http.Handler {
	mux := mux.NewRouter()
	mux.Handle("/health", handleHealth()).
		Methods(http.MethodGet)
	mux.Handle("/hosts", handleCreateHost(core)).
		Methods(http.MethodPost)
	mux.Handle("/hosts", handleGetHosts(core)).
		Methods(http.MethodGet)
	mux.Handle("/hosts/{id:[0-9]+}", handleGetHost(core)).
		Methods(http.MethodGet)
	mux.Handle("/hosts/{id:[0-9]+}", handleRemoveHost(core)).
		Methods(http.MethodDelete)

	return mux
}

func respondError(w http.ResponseWriter, status int, err error) {
	type response struct {
		Errors []string `json:"errors"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := response{Errors: make([]string, 0, 1)}
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
	}

	json.NewEncoder(w).Encode(resp)
}

func respondOk(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json")
	if body == nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	json.NewEncoder(w).Encode(body)
}
