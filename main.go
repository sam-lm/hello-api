package main

import (
	"log"
	"net/http"

	"github.com/codeformio/hello-api/api"
	"github.com/kelseyhightower/envconfig"
)

const version = "1.1.0"

func main() {
	var cfg struct {
		Addr    string `envconfig:"ADDR" default:":8080"`
		Message string `envconfig:"MESSAGE" default:"Hello!"`
	}
	envconfig.MustProcess("", &cfg)

	log.Printf("starting to listen on addr %v", cfg.Addr)

	log.Fatal(http.ListenAndServe(cfg.Addr, &api.Handler{Msg: cfg.Message, Version: version}))
}
