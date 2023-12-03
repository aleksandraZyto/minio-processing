package main

import (
	"log"

	h "github.com/aleksandraZyto/minio-processing/handlers"
	s "github.com/aleksandraZyto/minio-processing/services"
	st "github.com/aleksandraZyto/minio-processing/storage"
)

func main() {
	storage := &st.FileStorage{}	
	service := &s.FileService{
		Storage: storage,
	}
	handler := h.NewHandler(service)
	if err := handler.Server.ListenAndServe(); err != nil {
		log.Printf("Error occurred when serving %v\n", err)
	}
}
