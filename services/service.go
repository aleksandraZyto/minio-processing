package services

import (
	"context"
	"log"

	s "github.com/aleksandraZyto/minio-processing/storage"
)

type Service interface {
	GetFile(id string) (string, error)
	PutFile(id string, content string) error
}

type FileService struct {
	Ctx     context.Context
	Storage s.Storage
}

func (fs *FileService) GetFile(id string) (string, error) {
	content, err := fs.Storage.GetFile(fs.Ctx, id)
	if err != nil {
		log.Println("Error getting file from storage")
		return "", err
	}
	return content, nil
}

func (fs *FileService) PutFile(id string, content string) error {
	log.Println(content)
	return nil
}
