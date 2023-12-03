package db

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	c "github.com/aleksandraZyto/minio-processing/constants"
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
		log.Printf("Client created: %s", md.Name)
		minioClients = append(minioClients, client)
	}
	return minioClients, nil
}

func newClient(ctx context.Context, minioDetails MinioDetails) (*minio.Client, error) {
	endpoint := minioDetails.Name+c.PortNumber
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

	log.Printf("Succussfully created minio client %s", minioDetails.Name)
	return minioClient, nil
}

func newBucket(ctx context.Context, mc *minio.Client) error {
	log.Printf("Creating new minio bucket at endpoint %s", mc.EndpointURL())
	bucketName := c.BucketName
	err := mc.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := mc.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket '%s' already exists at andpoint %s", bucketName, mc.EndpointURL())
			return nil
		} else {
			log.Printf("Error making bucket: %v", err)
			return err
		}
	}
	log.Printf("Succussfully created bucket %s", bucketName)
	return nil
}
