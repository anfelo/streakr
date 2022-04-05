package users

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserDocument - struct that defines the firebase user document model
type UserDocument struct {
	FirstName string
	LastName  string
	Email     string
	Role      string
}

// Database - struct that defines the users db
type Database struct {
	Client *firestore.Client
}

// NewDatabase - function that creates a new database instance
func NewDatabase(ctx context.Context) (*Database, error) {
	firestoreClient, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to the database: %w", err)
	}
	return &Database{Client: firestoreClient}, nil
}

// GetUserByID - method that returns a user by id from the db
func (d *Database) GetUserByID(ctx context.Context, userID string) (User, error) {
	doc, err := d.userDocumentRef(userID).Get(ctx)
	if err != nil && status.Code(err) != codes.NotFound {
		return User{}, err
	}
	if err != nil && status.Code(err) == codes.NotFound {
		return User{}, err
	}

	var user User
	err = doc.DataTo(&user)
	if err != nil {
		return User{}, err
	}
	user.ID = doc.Ref.ID
	return user, nil
}

// CreateUser
func (d *Database) CreateUser(ctx context.Context, user User) (User, error) {
	return User{}, nil
}

// UpdateUser
func (d *Database) UpdateUser(ctx context.Context, userID string, user User) (User, error) {
	return User{}, nil
}

// DeleteUser
func (d *Database) DeleteUser(ctx context.Context, userID string) error {
	return nil
}

func (db *Database) usersCollection() *firestore.CollectionRef {
	return db.Client.Collection("users")
}

func (db *Database) userDocumentRef(userID string) *firestore.DocumentRef {
	return db.usersCollection().Doc(userID)
}
