package user

import (
	"net/http"

	"github.com/HadeedTariq/go-ecommerce-api/types"
	"github.com/HadeedTariq/go-ecommerce-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegiser).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegiser(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload

	err := utils.ParseJson(r, payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// ~ ok so over there we have to check that the user exist with in the db or not and for that we have to introduce the databases

}
