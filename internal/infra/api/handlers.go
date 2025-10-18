package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/gorilla/mux"
	"github.com/moq77111113/kite/internal/application/describe"
	"github.com/moq77111113/kite/internal/application/list"
)

func respondJSON(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func respondError(w http.ResponseWriter, message string, status int) {
	respondJSON(w, map[string]string{"error": message}, status)
}

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, map[string]string{"status": "ok"}, http.StatusOK)
}

func (s *Server) listKits(w http.ResponseWriter, r *http.Request) {

	listSvc := list.New(s.container.Repository, s.container.InstallationRegistry)
	kits, err := listSvc.Execute()

	if err != nil {
		respondError(w, fmt.Sprintf("Failed to fetch kits: %v", err), http.StatusInternalServerError)
		return
	}

	tag := r.URL.Query().Get("tag")
	if tag != "" {
		filtered := make([]any, 0)
		for _, k := range kits {
			if slices.Contains(k.Tags, tag) {
				filtered = append(filtered, k)
			}
		}
		respondJSON(w, map[string]any{"kits": filtered}, http.StatusOK)
		return
	}

	respondJSON(w, map[string]any{
		"kits": kits,
	}, http.StatusOK)
}

func (s *Server) getKit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if name == "" {
		respondError(w, "Kit name is required", http.StatusBadRequest)
		return
	}

	desc := describe.New(s.container.Repository, s.container.InstallationRegistry)

	item, err := desc.Execute(name)
	if err != nil {
		if err.Error() == fmt.Sprintf("kit '%s' not found", name) {
			respondError(w, err.Error(), http.StatusNotFound)
			return
		}
		respondError(w, fmt.Sprintf("Failed to fetch kit: %v", err), http.StatusInternalServerError)
		return
	}

	respondJSON(w, item, http.StatusOK)
}
