package main

import (
	"context"
	"fmt"

	"github.com/anfelo/streakr/server/internal/users"
)

// Run - Responsible for the instantiation
// and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")

	ctx := context.Background()
	db, err := users.NewDatabase(ctx)
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	userService := users.NewService(db)
	httpHandler := users.NewHandler(userService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Streakr REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
