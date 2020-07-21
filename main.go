package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

const version = 1

func main() {
	var cfg struct {
		Addr    string `envconfig:"ADDR" default:":8080"`
		Message string `envconfig:"MESSAGE" default:"Hello!"`
	}
	envconfig.MustProcess("", &cfg)

	log.Printf("starting to listen on addr %v", cfg.Addr)

	log.Fatal(http.ListenAndServe(cfg.Addr, &handler{msg: cfg.Message}))
}

type response struct {
	Msg     string `json:"message"`
	Version int    `json:"version"`
}

type handler struct {
	msg string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("received request: %s", r.URL.Path)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Msg: h.msg, Version: version})
}
