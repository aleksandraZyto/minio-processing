package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	s "github.com/aleksandraZyto/minio-processing/services"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Server  *http.Server
	Service s.Service
}

func NewHandler(service s.Service) *Handler {
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: router,
	}
	handler := &Handler{
		Router:  router,
		Server:  server,
		Service: service,
	}

	handler.registerHandlers()

	return handler
}

func (h *Handler) registerHandlers() {
	h.Router.HandleFunc("/file/{id}", h.GetFile).Methods("GET")
	h.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
	})
}

func (h *Handler) GetFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		log.Println("Provided empty id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	content, err := h.Service.GetFile(id)
	if err != nil {
		log.Printf("Error getting file with id %s", id)
		return
	}

	respJSON, err := json.Marshal(content)
	if err != nil {
		log.Printf("Error parsing the response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully retrieved file with id %s", id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJSON)
}
