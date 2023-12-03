package storage

import (
	"context"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
)

type Storage interface {
	GetFile(ctx context.Context, id string) (string, error)
	PutFile(ctx context.Context, id string, content string) error
}

type FileStorage struct {
	Minio *minio.Client
}

func (fs *FileStorage) GetFile(ctx context.Context, id string) (string, error) {
	file, err := fs.Minio.GetObject(ctx, "filestorage", "test.txt", minio.GetObjectOptions{})
	if err != nil {
		log.Printf("Getting file from bucket failed: %v", err)
		return "", err
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading the file: %v", err)
		return "", err
	}

	log.Printf("Successfully retrieved file %s", id)
	return string(content), nil
}

func (fs *FileStorage) PutFile(ctx context.Context, id string, content string) error {
	log.Printf("Putting file iin storage %s", id)
	return nil
}
