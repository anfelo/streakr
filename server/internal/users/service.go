package users

import (
	"context"
	"net/http"

	"github.com/anfelo/streakr/server/internal/common/transport/errors"
	transportHTTP "github.com/anfelo/streakr/server/internal/common/transport/http"
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
		restErr := errors.NewBadRequestError("invalid user id")
		transportHTTP.RespondJson(w, restErr.Status, restErr)
		return
	}

	user, err := h.Service.GetUserByID(r.Context(), id)
	if err != nil {
		restErr := errors.NewNotFoundError("user not found")
		transportHTTP.RespondJson(w, restErr.Status, restErr)
		return
	}

	transportHTTP.RespondJson(w, http.StatusOK, user)
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
