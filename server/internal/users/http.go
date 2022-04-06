package users

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anfelo/streakr/server/internal/common/transport/logs"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service UserService
	Server  *http.Server
}

func NewHandler(service UserService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.Router.Use(logs.LoggingMiddleware)
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/api/v1/users/{id}", h.GetUserByID).Methods("GET")
	h.Router.HandleFunc("/api/v1/users", h.CreateUser).Methods("POST")
	h.Router.HandleFunc("/api/v1/users/{id}", h.UpdateUser).Methods("PUT")
	h.Router.HandleFunc("/api/v1/users/{id}", h.DeleteUser).Methods("DELETE")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := h.Server.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("shut down gracefully")
	return nil
}
