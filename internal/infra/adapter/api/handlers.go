package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// respondJSON writes a JSON response with the given status code
func respondJSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// respondError writes a JSON error response
func respondError(w http.ResponseWriter, message string, status int) {
	respondJSON(w, map[string]string{"error": message}, status)
}

// healthCheck returns a simple health status
func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, map[string]string{"status": "ok"}, http.StatusOK)
}

// listTemplates returns all available templates from the registry
func (s *Server) listTemplates(w http.ResponseWriter, r *http.Request) {
	templates, err := s.service.ListAvailable()
	if err != nil {
		respondError(w, fmt.Sprintf("Failed to fetch templates: %v", err), http.StatusInternalServerError)
		return
	}

	respondJSON(w, map[string]any{
		"templates": templates,
	}, http.StatusOK)
}

// getTemplate returns details for a specific template including files
func (s *Server) getTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if name == "" {
		respondError(w, "Template name is required", http.StatusBadRequest)
		return
	}

	template, err := s.service.GetDetails(name)
	if err != nil {
		respondError(w, fmt.Sprintf("Template '%s' not found: %v", name, err), http.StatusNotFound)
		return
	}

	respondJSON(w, template, http.StatusOK)
}
