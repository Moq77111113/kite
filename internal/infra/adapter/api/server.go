package api

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moq77111113/kite/internal/application/template"
	"github.com/moq77111113/kite/web"
)

// Server handles HTTP requests for the Kite API and web UI
type Server struct {
	service *template.Service
	router  *mux.Router
}

// New creates a new Server instance
func New(svc *template.Service) *Server {
	s := &Server{
		service: svc,
	}
	s.setupRoutes()
	return s
}

// setupRoutes configures all HTTP routes
func (s *Server) setupRoutes() {
	s.router = mux.NewRouter()

	api := s.router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/templates", s.listTemplates).Methods("GET")
	api.HandleFunc("/templates/{name}", s.getTemplate).Methods("GET")

	staticFS, err := fs.Sub(web.DistFS, "dist")
	if err != nil {
		panic("failed to load embedded web UI: " + err.Error())
	}

	s.router.PathPrefix("/").Handler(spaHandler(staticFS))
}

// spaHandler serves the SPA and falls back to index.html for unknown routes
func spaHandler(staticFS fs.FS) http.Handler {
	fileServer := http.FileServer(http.FS(staticFS))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		f, err := staticFS.Open(path[1:])
		if err == nil {
			f.Close()
			fileServer.ServeHTTP(w, r)
			return
		}

		r.URL.Path = "/"
		fileServer.ServeHTTP(w, r)
	})
}

func (s *Server) Start(port string) error {
	addr := ":" + port
	return http.ListenAndServe(addr, s.router)
}
