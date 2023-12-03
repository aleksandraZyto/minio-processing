package db

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

)

type MinioKey string

type MinioDetails struct {
	Name      string
	AccessKey string
	SecretKey string
}

func GenerateClients(ctx context.Context, minioDetails []MinioDetails) ([]*minio.Client, error) {
	log.Println("Generating clients")
	var minioClients []*minio.Client
	for _, md := range minioDetails {
		client, err := newClient(ctx, md)
		if err != nil {
			log.Printf("Error creating client %s", md.Name)
		}
		log.Printf("Successfully created client %s", md.Name)
		minioClients = append(minioClients, client)
	}
	return minioClients, nil
}

func newClient(ctx context.Context, minioDetails MinioDetails) (*minio.Client, error) {
	endpoint := minioDetails.Name+":9000"
	accessKeyID := minioDetails.AccessKey
	secretAccessKey := minioDetails.SecretKey
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Printf("Error creating minio client: %v", err)
		return minioClient, err
	}

	err = newBucket(ctx, minioClient)
	if err != nil {
		log.Println("Error creating minio bucket")
		return minioClient, err
	}

	log.Println("Succussfully created minio client")
	return minioClient, nil
}

func newBucket(ctx context.Context, mc *minio.Client) error {
	log.Println("Creating new minio bucket")
	bucketName := "filestorage"
	err := mc.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := mc.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket '%s' already exists", bucketName)
			return nil
		} else {
			log.Printf("Error making bucket: %v", err)
			return err
		}
	}
	log.Printf("Succussfully created bucket %s", bucketName)
	return nil
}
