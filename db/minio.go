package db

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewClient(ctx context.Context) (*minio.Client, error) {
	endpoint := "amazin-object-storage-node-1:9000"
	accessKeyID := "ring"
	secretAccessKey := "treepotato"
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
