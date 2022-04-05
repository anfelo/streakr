package users

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

var (
	ErrFetchingUser   = errors.New("failed to fetch user by id")
	ErrNotImplemented = errors.New("not implemented")
)

// User - a representation of the user structure
type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      string
}

// Store - the main interface that describes
// how to interact with the store or repository layer
type Store interface {
	GetUserByID(context.Context, string) (User, error)
	CreateUser(context.Context, User) (User, error)
	UpdateUser(context.Context, string, User) (User, error)
	DeleteUser(context.Context, string) error
}

// Service - is the struct on which all our logic
// will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetUserByID - returns a user by its ID
func (s *Service) GetUserByID(ctx context.Context, ID string) (User, error) {
	usr, err := s.Store.GetUserByID(ctx, ID)
	if err != nil {
		log.Error("error fetching user by id")
		return User{}, ErrFetchingUser
	}
	return usr, nil
}

// CreateUser - creates a new user
func (s *Service) CreateUser(ctx context.Context, newUsr User) (User, error) {
	insertedUsr, err := s.Store.CreateUser(ctx, newUsr)
	if err != nil {
		log.Error("error creating user")
		return User{}, err
	}
	return insertedUsr, nil
}

// UpdateUser - updates a new user
func (s *Service) UpdateUser(
	ctx context.Context,
	ID string,
	updatedUsr User,
) (User, error) {
	updatedUsr, err := s.Store.UpdateUser(ctx, ID, updatedUsr)
	if err != nil {
		log.Error("error updating user")
		return User{}, err
	}
	return updatedUsr, nil
}

// DeleteUser - deletes a user by its id
func (s *Service) DeleteUser(ctx context.Context, ID string) error {
	return s.Store.DeleteUser(ctx, ID)
}
