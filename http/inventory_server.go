package http

import (
	"bookstore/application"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) registerInventoryRoutes(r *mux.Router) {
	r.HandleFunc("/inventory", s.getInventory).Methods("GET")
}

func (s *Server) getInventory(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("Accept") {
	case "application/json":
		inventoryDTO, err := s.InventoryService.GetInventory()
		if err != nil {
			httpCode := errorStatusCode(application.ErrorCode(err))
			handleErrorAsJson(w, r, httpCode, application.ErrorMessage(err), err)

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(inventoryDTO)
	default:
		handleNotAcceptable(w, r)

		return
	}
}
