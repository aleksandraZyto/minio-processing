package services

import "log"

type Service interface {
	GetFile(id string) (content string, err error)
}

type FileService struct{}

func (fs *FileService) GetFile(id string) (content string, err error) {
	log.Println(fs)
	return "File from service", nil
}
