package main

import (
	"log"

	h "github.com/aleksandraZyto/minio-processing/handlers"
	s "github.com/aleksandraZyto/minio-processing/services"
)

func main() {
	service := &s.FileService{}
	handler := h.NewHandler(service)
	if err := handler.Server.ListenAndServe(); err != nil {
		log.Printf("Error occurred when serving %v\n", err)
	}
}
