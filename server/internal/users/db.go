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
	ID        string `firestore:"id"`
	FirstName string `firestore:"firstName"`
	LastName  string `firestore:"lastName"`
	Email     string `firestore:"email"`
	Role      string `firestore:"role"`
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
func (db *Database) GetUserByID(ctx context.Context, userID string) (User, error) {
	doc, err := db.userDocumentRef(userID).Get(ctx)
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
	return user, nil
}

// CreateUser - method that creates a new user document
func (db *Database) CreateUser(ctx context.Context, user User) (User, error) {
	newDocRef := db.userDocumentRef(user.ID)
	userDoc := UserDocument{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
	}
	_, err := newDocRef.Create(ctx, userDoc)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// UpdateUser - method that updates a user document
func (db *Database) UpdateUser(
	ctx context.Context,
	userID string,
	user User,
) (User, error) {
	docRef := db.userDocumentRef(userID)
	doc, err := docRef.Get(ctx)
	if err != nil && status.Code(err) != codes.NotFound {
		return User{}, err
	}
	if err != nil && status.Code(err) == codes.NotFound {
		return User{}, err
	}

	var userDoc UserDocument
	err = doc.DataTo(&userDoc)
	if err != nil {
		return User{}, err
	}

	userDoc.FirstName = user.FirstName
	userDoc.LastName = user.LastName
	_, err = docRef.Set(ctx, userDoc)
	if err != nil {
		return User{}, err
	}

	return mapUserDocumentToUser(userDoc), nil
}

// DeleteUser - method that deletes a user document
func (db *Database) DeleteUser(ctx context.Context, userID string) error {
	docRef := db.userDocumentRef(userID)
	_, err := docRef.Get(ctx)
	if err != nil && status.Code(err) != codes.NotFound {
		return err
	}
	if err != nil && status.Code(err) == codes.NotFound {
		return err
	}

	if _, err := docRef.Delete(ctx); err != nil {
		return err
	}
	return nil
}

func (db *Database) usersCollection() *firestore.CollectionRef {
	return db.Client.Collection("users")
}

func (db *Database) userDocumentRef(userID string) *firestore.DocumentRef {
	return db.usersCollection().Doc(userID)
}

func mapUserDocumentToUser(userDoc UserDocument) User {
	return User{
		ID:        userDoc.ID,
		FirstName: userDoc.FirstName,
		LastName:  userDoc.LastName,
		Email:     userDoc.Email,
		Role:      userDoc.Role,
	}
}
