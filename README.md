# Minio Object Storage API
The Minio Processing API is a RESTful API for retrieving and storing objects.
It spins up multiple MinIO instances and allows you to store obejcts on them.
MinIO instance to store an object is chosen based on the id of the object, so that they are distributed evenly.

## Getting Started
#Prerequisites
Go
Docker

## Running the app
```docker-compose up --build```

## API Endpoints
### Get Object
Request:
```GET /object/{id}```
Response:
```
{
  "content": "object-content"
}
```

### Put Object
```PUT /object/{id}```
Request body:
```
{
	"Content": "testing"
}
```