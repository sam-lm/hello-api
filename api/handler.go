package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Msg     string `json:"message"`
	Version string `json:"version"`
}

type Handler struct {
	Msg     string
	Version string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request: %s", r.URL.Path)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Msg: h.Msg, Version: h.Version})
}
