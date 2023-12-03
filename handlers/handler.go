package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
	Server *http.Server
}

func NewHandler() *Handler {
	router := mux.NewRouter()
	server := &http.Server{
		Addr: "0.0.0.0:3000",
		Handler: router,
	}
	handler := &Handler{
		Router: router,
		Server: server,
	}

	handler.registerHandlers()

	return handler
}

func (h *Handler) registerHandlers(){
	h.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Println("Hello World!")
	})
}