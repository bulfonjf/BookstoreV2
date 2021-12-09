package http

import (
	"bookstore/application"
	"context"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
)

const ShutdownTimeout = 1 * time.Second

type Server struct {
	ln     net.Listener
	server *http.Server
	router *mux.Router

	Addr   string
	Domain string

	BookService *application.BookService
	InventoryService *application.InventoryService
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		router: mux.NewRouter(),
	}

	s.router.Use(reportPanic)

	s.router.NotFoundHandler = http.HandlerFunc(handleNotFound)

	router := s.router.PathPrefix("/").Subrouter()
	s.registerBookRoutes(router)
	s.registerInventoryRoutes(router)

	s.server.Handler = router

	return s
}

func (s *Server) Open() (err error) {
	if s.Domain != "" {
		s.ln = autocert.NewListener(s.Domain)
	} else {
		if s.ln, err = net.Listen("tcp", s.Addr); err != nil {
			return err
		}
	}

	go s.server.Serve(s.ln)

	return nil
}

func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}

func reportPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func handleInvalidContentType(w http.ResponseWriter, r *http.Request) {
	handleErrorAsJson(w, r, http.StatusUnsupportedMediaType, "Content-type header must be application/json", nil)
}

func handleNotAcceptable(w http.ResponseWriter, r *http.Request) {
	handleErrorAsJson(w, r, http.StatusNotAcceptable, "The client doesn't accept application/json", nil)
}

func handleBadScheme(w http.ResponseWriter, r *http.Request, err error) {
	handleErrorAsJson(w, r, http.StatusBadRequest, "Invalid scheme", err)
}
