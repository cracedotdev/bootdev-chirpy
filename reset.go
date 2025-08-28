package main

import (
	"errors"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, 403, "Not authorized", errors.New("only for dev environment"))
		return
	}

	err := cfg.db.DeleteAllUsers(r.Context())
	err = cfg.db.DeleteAllChirps(r.Context())
	if err != nil {
		respondWithError(w, 500, "Failed to reset db tables", err)
	}
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
