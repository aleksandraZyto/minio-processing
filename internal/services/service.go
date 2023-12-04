package services

import (
	"context"
	s "github.com/aleksandraZyto/minio-processing/internal/storage"
	"log"
	"time"
)

type Service interface {
	GetFile(ctx context.Context, id string) (string, error)
	PutFile(ctx context.Context, id string, content string) error
}

type FileService struct {
	Storage s.Storage
}

func (fs *FileService) GetFile(ctx context.Context, id string) (string, error) {
	time.Sleep(6 * time.Second)
	content, err := fs.Storage.GetFile(ctx, id)
	if err != nil {
		log.Println("Error getting file from storage")
		return "", err
	}
	return content, nil
}

func (fs *FileService) PutFile(ctx context.Context, id string, content string) error {
	err := fs.Storage.PutFile(ctx, id, content)
	if err != nil {
		log.Println("Error putting the file into storage")
		return err
	}
	return nil
}
