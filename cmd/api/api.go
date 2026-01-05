package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/HadeedTariq/go-ecommerce-api/services/user"
	"github.com/gorilla/mux"
)

// ~ so this is the entry point for the go application
type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// ~ ok so over there the user service is registered for handling the request

	// ~ so this is how the depenedency injection works out with in the golang
	userStore := user.NewStore(s.db)
	userService := user.NewHandler(userStore)
	userService.RegisterRoutes(subRouter)

	log.Println("listening on : ", s.addr)

	return http.ListenAndServe(s.addr, nil)
}
