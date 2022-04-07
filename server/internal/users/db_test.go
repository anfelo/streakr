package users

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/require"
)

func TestDatabase_CreateUser(t *testing.T) {
	t.Parallel()
	db := newDatabase(t)
	testCases := []struct {
		Name            string
		UserConstructor func(t *testing.T) User
	}{
		{
			Name:            "base_user",
			UserConstructor: newUser,
		},
	}
	for _, c := range testCases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			expectedUser := c.UserConstructor(t)

			_, err := db.CreateUser(ctx, expectedUser)
			require.NoError(t, err)
		})
	}
}

func newDatabase(t *testing.T) *Database {
	// TODO: Run firebase emulator instead of testing against prod db
	// Or create another DB for testing
	firestoreClient, err := firestore.NewClient(context.Background(), os.Getenv("GCP_PROJECT"))
	require.NoError(t, err)

	return &Database{Client: firestoreClient}
}

func newUser(t *testing.T) User {
	return User{
		ID:        "some_id",
		FirstName: "Andres",
		LastName:  "Osorio",
		Email:     "test@test.com",
		Role:      "test",
	}
}
