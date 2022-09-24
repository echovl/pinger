package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/echovl/pinger"
	"github.com/gorilla/mux"
)

func handleCreateHost(core *pinger.Core) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var host pinger.Host
		err := json.NewDecoder(r.Body).Decode(&host)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		if err = pinger.ValidateHost(&host); err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		err = core.UpsertHost(r.Context(), &host)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		respondOk(w, host)
	})
}

func handleGetHosts(core *pinger.Core) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			limit int = 50
			skip  int
		)

		if err := r.ParseForm(); err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		if q := r.Form.Get("limit"); q != "" {
			limit, _ = strconv.Atoi(q)
		}
		if q := r.Form.Get("skip"); q != "" {
			skip, _ = strconv.Atoi(q)
		}

		hosts, err := core.GetHosts(r.Context(), limit, skip)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}
		respondOk(w, hosts)
	})
}

func handleGetHost(core *pinger.Core) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
		}

		host, err := core.GetHost(r.Context(), hostID)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}
		respondOk(w, host)
	})
}

func handleRemoveHost(core *pinger.Core) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hostID, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
		}

		err = core.RemoveHost(r.Context(), hostID)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}
		respondOk(w, nil)
	})
}
