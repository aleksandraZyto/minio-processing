package main

import (
	"context"
	"log"

	db "github.com/aleksandraZyto/minio-processing/db"
	d "github.com/aleksandraZyto/minio-processing/docker"
	h "github.com/aleksandraZyto/minio-processing/handlers"
	s "github.com/aleksandraZyto/minio-processing/services"
	st "github.com/aleksandraZyto/minio-processing/storage"
	"github.com/docker/docker/client"
	"github.com/minio/minio-go/v7"
)

func main() {
	ctx := context.Background()

	dockerClient, err := setupDockerClient()
	if err != nil {
		log.Printf("Error setting up the docker client: %v", err)
	}

	minioClients, err := setupMinioClient(ctx, dockerClient)
	if err != nil {
		log.Println("Error configuring minio client")
	}

	storage := &st.FileStorage{
		Minio: minioClients,
	}
	service := &s.FileService{
		Storage: storage,
	}
	handler := h.NewHandler(service)
	if err := handler.Server.ListenAndServe(); err != nil {
		log.Printf("Error occurred when serving %v\n", err)
	}
}

func setupMinioClient(ctx context.Context, dockerClient *client.Client) ([]*minio.Client, error) {
	log.Println("Starting to setup minio clients")
	minioDetails, err := d.GetMinioDetails(dockerClient)
	if err != nil {
		log.Println("Error getting minio instances details")
		return nil, err
	}

	clients, err := db.GenerateClients(ctx, minioDetails)
	if err != nil {
		return clients, err
	}
	return clients, nil
}

func setupDockerClient() (*client.Client, error) {
	log.Println("Starting the setup of docker client")
	dockerClient, err := d.NewClient()
	if err != nil {
		return dockerClient, err
	}
	log.Println("Docker client configured")
	return dockerClient, nil
}
