package api

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moq77111113/kite/internal/infra/cli/container"
	"github.com/moq77111113/kite/web"
)

type Server struct {
	container *container.Container
	router    *mux.Router
}

func New(container *container.Container) *Server {
	s := &Server{
		container: container,
	}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.router = mux.NewRouter()

	api := s.router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/health", s.healthCheck).Methods("GET")
	api.HandleFunc("/kits", s.listKits).Methods("GET")
	api.HandleFunc("/kits/{name}", s.getKit).Methods("GET")

	staticFS, err := fs.Sub(web.DistFS, "dist")
	if err != nil {
		panic("failed to load embedded web UI: " + err.Error())
	}

	s.router.PathPrefix("/").Handler(spaHandler(staticFS))
}

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
