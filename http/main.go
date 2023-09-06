package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/stealthrocket/net/http"
	"github.com/stealthrocket/net/wasip1"
)

func main() {
	logRequest := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "hola, desde WASI!")
		fmt.Fprintln(w, "Cabeceras HTTP")
		for header, v := range r.Header {
			fmt.Fprintf(w, "====%s: %s\n", header, v)
		}
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", logRequest)
	addr := ":3000"
	listener, err := wasip1.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Listening and serving at", addr)
	srv := http.Server{
		Handler: mux,
		Addr:    addr,
	}
	err = srv.Serve(listener)
	log.Println("server shutdown", err)
}
