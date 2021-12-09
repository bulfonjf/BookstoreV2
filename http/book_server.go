package http

import (
	"bookstore/application"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) registerBookRoutes(r *mux.Router) {
	r.HandleFunc("/books", s.getBooks).Methods("GET")
	r.HandleFunc("/books/{bookID}", s.getBook).Methods("GET")
	r.HandleFunc("/books", s.createBook).Methods("POST")
	r.HandleFunc("/books", s.updateBook).Methods("PUT")
	r.HandleFunc("/books/{bookID}", s.deleteBook).Methods("DELETE")
}

func (s *Server) getBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("Accept") {
	case "application/json":
		booksDTO, err := s.BookService.GetBooks()
		if err != nil {
			httpCode := errorStatusCode(application.ErrorCode(err))
			handleErrorAsJson(w, r, httpCode, application.ErrorMessage(err), err)

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(booksDTO)
	default:
		handleNotAcceptable(w, r)

		return
	}
}

func (s *Server) getBook(w http.ResponseWriter, r *http.Request) {
	bookID := mux.Vars(r)["bookID"]
	switch r.Header.Get("Accept") {
	case "application/json":
		if bookID == "" {
			handleErrorAsJson(w, r, http.StatusBadRequest, "path param bookID is required", nil)

			return
		}

		bookDTO, err := s.BookService.GetBookByID(bookID)
		if err != nil {
			httpCode := errorStatusCode(application.ErrorCode(err))
			handleErrorAsJson(w, r, httpCode, application.ErrorMessage(err), err)

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bookDTO)
	default:
		handleNotAcceptable(w, r)

		return
	}
}

func (s *Server) createBook(w http.ResponseWriter, r *http.Request) {
	var createBookDTO application.CreateBookDTO
	switch r.Header.Get("Content-type") {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(&createBookDTO); err != nil {
			handleBadScheme(w, r, err)

			return
		}
	default:
		handleInvalidContentType(w, r)

		return
	}

	switch r.Header.Get("Accept") {
	case "application/json":
		bookDTO, err := s.BookService.CreateBook(createBookDTO)
		if err != nil {
			httpCode := errorStatusCode(application.ErrorCode(err))
			handleErrorAsJson(w, r, httpCode, application.ErrorMessage(err), err)

			return
		}

		s.InventoryService.AddBook(bookDTO)
		if err != nil {
			httpCode := errorStatusCode(application.ErrorCode(err))
			handleErrorAsJson(w, r, httpCode, application.ErrorMessage(err), err)

			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bookDTO)
	default:
		handleNotAcceptable(w, r)

		return
	}
}

func (s *Server) updateBook(w http.ResponseWriter, r *http.Request) {
	var updateBookDTO application.UpdateBookDTO
	switch r.Header.Get("Content-type") {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(&updateBookDTO); err != nil {
			handleBadScheme(w, r, err)

			return
		}
	default:
		handleInvalidContentType(w, r)

		return
	}

	switch r.Header.Get("Accept") {
	case "application/json":
		bookDTO, err := s.BookService.UpdateBook(updateBookDTO)
		if err != nil {
			httpCode := errorStatusCode(application.ErrorCode(err))
			handleErrorAsJson(w, r, httpCode, application.ErrorMessage(err), err)

			return
		}

		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(bookDTO)

		return
	default:
		handleNotAcceptable(w, r)

		return
	}
}

func (s *Server) deleteBook(w http.ResponseWriter, r *http.Request) {
	bookID := mux.Vars(r)["bookID"]
	switch r.Header.Get("Accept") {
	case "application/json":
		if bookID == "" {
			handleErrorAsJson(w, r, http.StatusBadRequest, "path param bookID is required", nil)

			return
		}

		err := s.BookService.DeleteBook(bookID)
		if err != nil {
			httpCode := errorStatusCode(application.ErrorCode(err))
			handleErrorAsJson(w, r, httpCode, application.ErrorMessage(err), err)

			return
		}

		w.WriteHeader(http.StatusNoContent)

		return
	default:
		handleNotAcceptable(w, r)

		return
	}
}
