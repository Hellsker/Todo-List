package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

// New
func New(handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler: handler,
		Addr:    ":8080",
	}
	server := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: 3 * time.Second,
	}
	server.start()
	return server
}
func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}
func (s *Server) Notify() <-chan error {
	return s.notify
}
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}
