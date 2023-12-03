package storage

import (
	"bytes"
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
	Minio []*minio.Client
}

func (fs *FileStorage) GetFile(ctx context.Context, id string) (string, error) {
	file, err := fs.Minio[0].GetObject(ctx, "filestorage", "test.txt", minio.GetObjectOptions{})
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

func (s *FileStorage) PutFile(ctx context.Context, id string, content string) error {
	contentBytes := []byte(content)

	_, err := s.Minio[0].PutObject(ctx, "filestorage", id, bytes.NewReader(contentBytes), int64(len(content)), minio.PutObjectOptions{})
	if err != nil {
		log.Printf("Error putting object to minio bucket: %v", err)
		return err
	}

	log.Printf("Object with id %s successfully uploaded to minio bucket", id)
	return nil
}
