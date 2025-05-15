package httpserver

import (
	"context"
	"errors"
	"log"
	"net/http"
)

type httpServer struct {
	server *http.Server
}

func New(address string, handler *http.ServeMux) *httpServer {
	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	return &httpServer{
		server: server,
	}
}

func (s *httpServer) Start() {
	log.Printf("Starting server on %q...\n", s.server.Addr)
	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server failed: %s", err)
	}
}

func (s *httpServer) Stop() {
	log.Println("Shutting down HTTP server gracefully...")
	if err := s.server.Shutdown(context.TODO()); err != nil {
		log.Fatalf("Graceful shutdown failed: %v", err)
	}
}
