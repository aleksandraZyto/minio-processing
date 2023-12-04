package storage

import (
	"bytes"
	"context"
	db "github.com/aleksandraZyto/minio-processing/config/minio"
	"io"
	"log"

	c "github.com/aleksandraZyto/minio-processing/constants"
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
	i := db.GetMinioInstance(id, len(fs.Minio))
	file, err := fs.Minio[i].GetObject(ctx, c.BucketName, id, minio.GetObjectOptions{})
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
	i := db.GetMinioInstance(id, len(fs.Minio))
	contentBytes := []byte(content)

	_, err := fs.Minio[i].PutObject(ctx, c.BucketName, id, bytes.NewReader(contentBytes), int64(len(content)), minio.PutObjectOptions{})
	if err != nil {
		log.Printf("Error putting object to minio bucket: %v", err)
		return err
	}

	log.Printf("Object with id %s successfully uploaded to minio instance %s", id, fs.Minio[i].EndpointURL())
	return nil
}
