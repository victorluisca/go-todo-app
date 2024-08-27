package main

import (
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /{test}", func(w http.ResponseWriter, r *http.Request) {
		test := r.PathValue("test")
		w.Write([]byte(test))
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Println("Server has started")

	return server.ListenAndServe()
}
