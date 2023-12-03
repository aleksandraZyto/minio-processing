package main

import (
	"log"

	h "github.com/aleksandraZyto/minio-processing/handlers"
)

func main() {
	handler := h.NewHandler()
	if err := handler.Server.ListenAndServe(); err != nil {
		log.Printf("Error occurred when serving %v\n", err)
	}
}
