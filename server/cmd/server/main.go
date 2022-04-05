package main

import (
	"context"
	"fmt"

	"github.com/anfelo/streakr/server/internal/users"
	log "github.com/sirupsen/logrus"
)

// App - contain application information
type App struct {
	Name    string
	Version string
}

// Run - Responsible for the instantiation
// and startup of our go application
func (a *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    a.Name,
			"AppVersion": a.Version,
		}).Info("Setting up application")

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
	log.Info("Streakr REST API")
	app := App{
		Name:    "Streakr App",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		log.Error("Error starting up our Web App")
		log.Fatal(err)
	}
}
