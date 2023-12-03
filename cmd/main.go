package main

import (
	"context"
	"log"

	db "github.com/aleksandraZyto/minio-processing/db"
	h "github.com/aleksandraZyto/minio-processing/handlers"
	s "github.com/aleksandraZyto/minio-processing/services"
	st "github.com/aleksandraZyto/minio-processing/storage"
	"github.com/minio/minio-go/v7"
)

func main() {
	ctx := context.Background()
	client, err := configureMinioClient(ctx)
	if err != nil {
		log.Println("Error configuring minio client")
	}
	storage := &st.FileStorage{
		Minio: client,
	}
	service := &s.FileService{
		Ctx:    ctx,
		Storage: storage,
	}
	handler := h.NewHandler(service)
	if err := handler.Server.ListenAndServe(); err != nil {
		log.Printf("Error occurred when serving %v\n", err)
	}
}

func configureMinioClient(ctx context.Context) (*minio.Client, error) {
	client, err := db.NewClient(ctx)
	if err != nil {
		return client, err
	}
	return client, nil
}
