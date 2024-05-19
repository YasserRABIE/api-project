package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *APIserver) run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleAccount))

	log.Println("json api is running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIserver) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch {
	case r.Method == "GET":
		return s.handleGetAccount(w, r)
	case r.Method == "POST":
		return s.handleCreateAccount(w, r)
	case r.Method == "DELETE":
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIserver) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println(id)

	return WriteJSON(w, http.StatusOK, &Account{})
}

func (s *APIserver) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIserver) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func makeHTTPHandleFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

type APIserver struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIserver {
	return &APIserver{
		listenAddr: listenAddr,
	}
}
