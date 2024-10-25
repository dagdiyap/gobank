package main

func (s *APIServer) routes() {
	s.router.HandleFunc("/login", makeHTTPHandleFunc(s.handleLogin))
	s.router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	s.router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccountByID))
	s.router.HandleFunc("/transfer", makeHTTPHandleFunc(s.handleTransfer))
}
