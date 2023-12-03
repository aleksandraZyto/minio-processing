package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	// "github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (sm *ServiceMock) GetFile(id string) (content string, err error) {
	args := sm.Called(id)
	return args.Get(0).(string), args.Error(1)
}

func TestRootEndpoint(t *testing.T) {
	mockService := &ServiceMock{}
	h := NewHandler(mockService)
	h.registerHandlers()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	h.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetFileHandler(t *testing.T) {
	expResp := "test-content"
	mockService := &ServiceMock{}
	mockService.On("GetFile", "1").Return(expResp, nil)

	h := NewHandler(mockService)
	h.registerHandlers()

	req, err := http.NewRequest("GET", "/file/1", nil)
	if err != nil {
		log.Println("Error adding test GET handler")
	}

	rr := httptest.NewRecorder()
	h.Router.ServeHTTP(rr, req)

	actResp := strings.Trim(rr.Body.String(), "\"")
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expResp, actResp)
}
