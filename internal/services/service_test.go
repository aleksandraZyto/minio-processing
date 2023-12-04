package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type StorageMock struct {
	mock.Mock
}

func (sm *StorageMock) GetFile(ctx context.Context, id string) (string, error) {
	args := sm.Called(ctx, id)
	return args.Get(0).(string), args.Error(1)
}

func (sm *StorageMock) PutFile(ctx context.Context, id string, content string) error {
	args := sm.Called(ctx, id, content)
	return args.Error(0)
}

func TestGetFileService(t *testing.T) {
	mockStorage := &StorageMock{}
	exp := "Content from storage"
	mockStorage.On("GetFile", context.Background(), "1").Return(exp, nil)

	service := &FileService{
		Storage: mockStorage,
	}

	content, err := service.GetFile(context.Background(),"1")

	assert.Equal(t, exp, content)
	assert.Equal(t, nil, err)
}

func TestPutObjectService(t *testing.T) {
	mockStorage := &StorageMock{}
	mockStorage.On("PutFile", context.Background(), "1", "test-content").Return(nil)

	service := &FileService{Storage: mockStorage}
	err := service.PutFile(context.Background(),"1", "test-content")

	assert.Equal(t, nil, err)
}
