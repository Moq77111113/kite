package server

import (
	"net/http"

	"github.com/gorilla/mux"
	registry "github.com/moq77111113/kite/api/registry/v1"
)

// Server handles HTTP requests for the Kite API and web UI
type Server struct {
	client registry.Client
	router *mux.Router
}

// New creates a new Server instance with the given registry client
func New(client registry.Client) *Server {
	s := &Server{
		client: client,
	}
	s.setupRoutes()
	return s
}

// setupRoutes configures all HTTP routes
func (s *Server) setupRoutes() {
	s.router = mux.NewRouter()

	// API routes under /api prefix
	api := s.router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/templates", s.listTemplates).Methods("GET")
	api.HandleFunc("/templates/{name}", s.getTemplate).Methods("GET")

	// Health check endpoint
	s.router.HandleFunc("/", s.healthCheck).Methods("GET")
}

// Start begins listening on the specified port
func (s *Server) Start(port string) error {
	addr := ":" + port
	return http.ListenAndServe(addr, s.router)
}
