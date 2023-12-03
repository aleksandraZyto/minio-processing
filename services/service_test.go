package services

import (
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
)

type StorageMock struct {
	mock.Mock
}

func (sm *StorageMock) GetFile(id string) (string, error){
	args := sm.Called(id)
	return args.Get(0).(string), args.Error(1)
}

func TestGetFileService(t *testing.T) {
	mockStorage := &StorageMock{}
	exp := "Content from storage"
	mockStorage.On("GetFile", "1").Return(exp, nil)

	service := &FileService{
		Storage: mockStorage,
	}

	content, err := service.GetFile("1")
	
	assert.Equal(t, exp, content)
	assert.Equal(t, nil, err)
}