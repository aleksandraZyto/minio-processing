package services

import (
	"context"
	"log"

	s "github.com/aleksandraZyto/minio-processing/storage"
	db "github.com/aleksandraZyto/minio-processing/db"
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
	minioDetails, ok := fs.Ctx.Value(db.MinioKey("minioDetails")).(db.MinioDetails)
	log.Println("MINIO:::::::")
	log.Println(minioDetails)
	log.Println(ok)
	content, err := fs.Storage.GetFile(fs.Ctx, id)
	if err != nil {
		log.Println("Error getting file from storage")
		return "", err
	}
	return content, nil
}

func (fs *FileService) PutFile(id string, content string) error {
	err := fs.Storage.PutFile(fs.Ctx, id, content)
	if err != nil {
		log.Println("Error putting the file into storage")
		return err
	}
	return nil
}
