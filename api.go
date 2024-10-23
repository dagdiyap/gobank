package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type APIServer struct {
	ListenAddr string
	store Storage
	router *mux.Router
	server *http.Server
}

func NewAPIServer(ListenAddr string, store Storage) *APIServer{
	router := mux.NewRouter()
	return &APIServer{
		ListenAddr: ListenAddr,
		store :  store,
		router: router,
	}
}

func Run(s *APIServer) {
	s.server = &http.Server{
		Addr: s.ListenAddr,
		Handler: s.router,
		ReadTimeout: 10,	
		WriteTimeout: 10,
		IdleTimeout: 60,	
	}

	s.routes()

	if err := s.server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}