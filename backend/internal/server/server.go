package server

import (
	"net/http"
	"peakly/internal/handler"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router      *chi.Mux
	taskHandler *handler.TaskHandler
}

func New() *Server {
	s := &Server{
		router: chi.NewRouter(),
	}
	s.registerRoutes()
	return s
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
