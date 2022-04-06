package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/anfelo/streakr/server/internal/common/transport/errors"
	transportHTTP "github.com/anfelo/streakr/server/internal/common/transport/http"
	"github.com/gorilla/mux"
)

// Response - struct used to respond with a simple message
type Response struct {
	Message string `json:"message"`
}

// UserResponse - struct that will be used as a response
// in the rest api
type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

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

	transportHTTP.RespondJson(w, http.StatusOK, mapToUserResponse(user))
}

// CreateUser - method used for creating a new user
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid user model")
		transportHTTP.RespondJson(w, restErr.Status, restErr)
		return
	}

	user, err := h.Service.CreateUser(r.Context(), user)
	if err != nil {
		restErr := errors.NewInternatServerError("internal server error")
		transportHTTP.RespondJson(w, restErr.Status, restErr)
		return
	}

	transportHTTP.RespondJson(w, http.StatusOK, mapToUserResponse(user))
}

// UpdateUser - method used for updating user's data
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		restErr := errors.NewBadRequestError("invalid user id")
		transportHTTP.RespondJson(w, restErr.Status, restErr)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid user model")
		transportHTTP.RespondJson(w, restErr.Status, restErr)
		return
	}

	user, err := h.Service.UpdateUser(r.Context(), id, user)
	if err != nil {
		restErr := errors.NewInternatServerError("internal server error")
		transportHTTP.RespondJson(w, restErr.Status, restErr)
		return
	}

	transportHTTP.RespondJson(w, http.StatusOK, mapToUserResponse(user))
}

// DeleteUser - method used for deleting a user by ID
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		restErr := errors.NewBadRequestError("invalid user id")
		transportHTTP.RespondJson(w, restErr.Status, restErr)
		return
	}

	err := h.Service.DeleteUser(r.Context(), id)
	if err != nil {
		restErr := errors.NewInternatServerError("internal server error")
		transportHTTP.RespondJson(w, restErr.Status, restErr)
		return
	}

	transportHTTP.RespondJson(w, http.StatusOK, Response{Message: "sucessfully deleted"})
}

func mapToUserResponse(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
	}
}
