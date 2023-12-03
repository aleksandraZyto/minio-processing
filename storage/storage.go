package storage

import (
	"bytes"
	"context"
	"hash/fnv"
	"io"
	"log"

	"github.com/minio/minio-go/v7"
	c "github.com/aleksandraZyto/minio-processing/constants"
)

type Storage interface {
	GetFile(ctx context.Context, id string) (string, error)
	PutFile(ctx context.Context, id string, content string) error
}

type FileStorage struct {
	Minio []*minio.Client
}

func (fs *FileStorage) GetFile(ctx context.Context, id string) (string, error) {
	file, err := fs.Minio[0].GetObject(ctx, c.BucketName, id, minio.GetObjectOptions{})
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
	i := getMinioInstance(id, len(fs.Minio))
	contentBytes := []byte(content)

	_, err := fs.Minio[i].PutObject(ctx, c.BucketName, id, bytes.NewReader(contentBytes), int64(len(content)), minio.PutObjectOptions{})
	if err != nil {
		log.Printf("Error putting object to minio bucket: %v", err)
		return err
	}

	log.Printf("Object with id %s successfully uploaded to minio instance %s", id, fs.Minio[i].EndpointURL())
	return nil
}

func getMinioInstance(id string, numStorages int) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(id))
	hashValue := hasher.Sum32()
	return int(hashValue % uint32(numStorages))
}
