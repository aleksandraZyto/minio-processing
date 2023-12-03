package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	c "github.com/aleksandraZyto/minio-processing/constants"
	s "github.com/aleksandraZyto/minio-processing/services"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Server  *http.Server
	Service s.Service
}

type Request struct {
	Content string `json:"content"`
}

func NewHandler(service s.Service) *Handler {
	router := mux.NewRouter()
	server := &http.Server{
		Addr:    c.AppAddress,
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
	h.Router.HandleFunc("/file/{id}", h.getFile).Methods("GET")
	h.Router.HandleFunc("/file/{id}", h.putFile).Methods("PUT")
	h.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
	})
}

func (h *Handler) getFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	content, err := h.Service.GetFile(id)
	if err != nil {
		if err.Error() == errors.New(c.KeyDoesNotExistErr).Error() {
			http.Error(w, "This id does not exist", http.StatusNotFound)
		} else {
			log.Printf("Error getting file with id %s: ", id)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

func (h *Handler) putFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var reqObj Request
	err = json.Unmarshal(body, &reqObj)
	if err != nil {
		log.Printf("Error unmarshalling request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(reqObj.Content) == 0 {
		log.Println("Invalid JSON format: 'content' field is required")
		http.Error(w, "'content' field is required", http.StatusBadRequest)
		return
	}

	err = h.Service.PutFile(id, reqObj.Content)
	if err != nil {
		log.Printf("Error putting object with id %s", id)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
