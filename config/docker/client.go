package docker

import (
	"context"
	"fmt"
	db "github.com/aleksandraZyto/minio-processing/config/minio"
	"log"
	"strings"

	cst "github.com/aleksandraZyto/minio-processing/constants"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func NewClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Printf("Error creating docker client: %v", err)
		return cli, err
	}
	log.Println("Successfully created docker client")
	return cli, err
}

func GetMinioDetails(client *client.Client) ([]db.MinioDetails, error) {
	options := types.ContainerListOptions{All: false}
	containers, err := client.ContainerList(context.Background(), options)
	if err != nil {
		fmt.Printf("Error retrieving containers: %v", err)
		return nil, err
	}
	log.Println("Container list retrieved")
	var minioContainerDetails []db.MinioDetails
	for _, c := range containers {
		if strings.Contains(c.Names[0], cst.MinioInstancePrefix) {
			containerDetails, _ := client.ContainerInspect(context.Background(), c.Names[0])
			minioDetails := db.MinioDetails{
				Name:      c.Names[0][1:],
				AccessKey: getProperty(containerDetails, cst.AccessKey),
				SecretKey: getProperty(containerDetails, cst.SecretKey),
			}

			minioContainerDetails = append(minioContainerDetails, minioDetails)
		}
	}
	log.Println("Collected minio details")
	return minioContainerDetails, nil
}

func getProperty(details types.ContainerJSON, name string) string {
	return getEnvVar(details.Config.Env, name)
}

func getEnvVar(envVars []string, key string) string {
	for _, envVar := range envVars {
		if value := parseEnvVar(envVar, key); value != "" {
			return value
		}
	}
	return ""
}

func parseEnvVar(envVar, key string) string {
	parts := strings.SplitN(envVar, "=", 2)
	if len(parts) == 2 && parts[0] == key {
		return parts[1]
	}
	return ""
}
