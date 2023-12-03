package handlers

import (
	"encoding/json"
	"io/ioutil"
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

type Request struct {
	Content string `json:"content"`
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
	h.Router.HandleFunc("/file/{id}", h.getFile).Methods("GET")
	h.Router.HandleFunc("/file/{id}", h.putFile).Methods("Put")
	h.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
	})
}

func (h *Handler) getFile(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) putFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		log.Println("Passed id is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(reqObj.Content) == 0 { // what if content is empty???
		log.Println("Invalid JSON format: 'content' field is required")
		w.WriteHeader(http.StatusBadRequest)
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
