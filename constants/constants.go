package constants

const (
	AppAddress= "0.0.0.0:"+AppPortNumber
	AppPortNumber = "3000"
	MinioPortNumber = "9000"
)

const (
	BucketName          = "filestorage"
	MinioInstancePrefix = "amazin-object-storage-node-"
	AccessKey           = "MINIO_ACCESS_KEY"
	SecretKey           = "MINIO_SECRET_KEY"
)

const (
	KeyDoesNotExistErr = "The specified key does not exist."
)
