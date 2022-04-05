package users

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// UserService - user service interface definition
type UserService interface {
	GetUserByID(context.Context, string) (User, error)
	CreateUser(context.Context, User) (User, error)
	UpdateUser(context.Context, string, User) (User, error)
	DeleteUser(context.Context, string) error
}

// GetUserByID - method that retrieves a user by its ID
func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetUserByID(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		panic(err)
	}
}

// CreateUser
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateUser
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
