package services

import (
	"log"

	s "github.com/aleksandraZyto/minio-processing/storage"
)

type Service interface {
	GetFile(id string) (content string, err error)
}

type FileService struct {
	Storage s.Storage
}

func (fs *FileService) GetFile(id string) (string, error) {
	content, err := fs.Storage.GetFile(id)
	if err != nil {
		log.Println("Error getting file from storage")
		return "", err
	}
	return content, nil
}
