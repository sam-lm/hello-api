package main

import (
	"log"
	"net/http"

	"github.com/codeformio/hello-api/api"
	"github.com/kelseyhightower/envconfig"
)

const version = "1.3.1"

func main() {
	var cfg struct {
		Addr    string `envconfig:"ADDR" default:":8080"`
		Message string `envconfig:"MESSAGE" default:"Hello!"`
	}
	envconfig.MustProcess("", &cfg)

	log.Printf("starting to listen on addr %v", cfg.Addr)
	log.Println("Hi from Nick!")

	log.Println("Hi from Dan")
	log.Fatal(http.ListenAndServe(cfg.Addr, &api.Handler{Message: cfg.Message, Version: version}))
}
