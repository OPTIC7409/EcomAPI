package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/OPTIC7409/EcomApi/service/user"
	"github.com/gorilla/mux"
)

type APIService struct {
	addr string
	db   *sql.DB
}

func NewApiService(addr string, db *sql.DB) *APIService {
	return &APIService{addr: addr, db: db}
}

func (s *APIService) Start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Starting server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
