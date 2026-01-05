package user

import (
	"fmt"
	"net/http"

	"github.com/HadeedTariq/go-ecommerce-api/types"
	"github.com/HadeedTariq/go-ecommerce-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
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

	_, err = h.store.GetUserByEmail(payload.Email)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with this %s email already exist", payload.Email))
	}

	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		Email:     payload.Email,
		LastName:  payload.LastName,
		Password:  payload.Password,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusCreated, nil)
}
