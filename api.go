package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Create a type struct APiServer
type APIServer struct {
	listenAddr string
	store      Storage
}

// Initialize a new APIServer
func newAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()                                          //Create new instance of router
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount)) //Route path to
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountById))
	log.Println("JSON API server is running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		return s.handleGetAccount(w, r)
	case "POST":
		return s.handleCreateAccount(w, r)
	case "DELETE":
		return s.handleDeleteAccount(w, r)
	default:
		return fmt.Errorf("method not allowed %s", r.Method)
	}
}

// An origin HandleFunc will not return an error
// Ex: func handleGetAccount(w  http.ResponseWriter, r *http.Request){}
// With this we'll have to handle the error inside of each HandleFunc and it will get messy.
// So in order to reduce the code we'll just handle the error inside of a single
// makeHandleFunc function
func (s *APIServer) handleGetAccount(responsewriter http.ResponseWriter, request *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}

	return WriteJSON(responsewriter, http.StatusOK, accounts)
}
func (s *APIServer) handleGetAccountById(responsewriter http.ResponseWriter, request *http.Request) error {
	vars := mux.Vars(request)
	fmt.Println(vars)
	return WriteJSON(responsewriter, http.StatusFound, &Account{})
}
func (s *APIServer) handleCreateAccount(responsewriter http.ResponseWriter, request *http.Request) error {
	createAccountReq := new(CreateAccountRequest)
	account := NewAccount(createAccountReq.FirstName, createAccountReq.LastName)
	if err := decodeFromRequestBodyToAccount(request, &account); err != nil {
		return err
	}
	if err := s.store.CreateAccount(&account); err != nil {
		return err
	}
	return WriteJSON(responsewriter, http.StatusOK, account)
}
func decodeFromRequestBodyToAccount(request *http.Request, account *Account) error {
	err := json.NewDecoder(request.Body).Decode(account)
	return err
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Sub Func
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
