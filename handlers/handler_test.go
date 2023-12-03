package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	c "github.com/aleksandraZyto/minio-processing/constants"
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

func (sm *ServiceMock) PutFile(id string, content string) error {
	args := sm.Called(id, content)
	return args.Error(0)
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
	testCases := []struct {
		testName     string
		id           string
		expResp      string
		serviceErr   error
		expectedCode int
	}{
		{
			testName:     "Happy path",
			id:           "1",
			expResp:      "test-content",
			serviceErr:   nil,
			expectedCode: http.StatusOK,
		},
		{
			testName:     "Return StatusNotFound if key does not exist",
			id:           "1",
			expResp:      "",
			serviceErr:   errors.New(c.KeyDoesNotExistErr),
			expectedCode: http.StatusNotFound,
		},
		{
			testName:     "Return InternalServerError if service fails",
			id:           "1",
			expResp:      "",
			serviceErr:   errors.New("Some error"),
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			mockService := &ServiceMock{}
			mockService.On("GetFile", tc.id).Return(tc.expResp, tc.serviceErr)

			h := NewHandler(mockService)
			h.registerHandlers()

			req, err := http.NewRequest("GET", "/file/"+tc.id, nil)
			if err != nil {
				t.Fatal("Error adding test GET handler")
			}

			rr := httptest.NewRecorder()
			h.Router.ServeHTTP(rr, req)

			assert.Equal(t, tc.expectedCode, rr.Code)

			if tc.expectedCode == http.StatusOK {
				actResp := strings.Trim(rr.Body.String(), "\"")
				assert.Equal(t, tc.expResp, actResp)
			}
		})
	}
}

func TestPutFileHandler(t *testing.T) {
    testCases := []struct {
        testName      string
        id            string
        content       string
        serviceErr    error
        expectedCode  int
    }{
        {
            testName:     "Happy path",
            id:           "1",
            content:      "test-content",
            serviceErr:   nil,
            expectedCode: http.StatusOK,
        },
        {
            testName:     "Return InternalServerError if service fails",
            id:           "1",
            content:      "test-content",
            serviceErr:   errors.New("Error from service"),
            expectedCode: http.StatusInternalServerError,
        },
        {
            testName:     "Return BadRequest if content is empty",
            id:           "1",
            content:      "",
            serviceErr:   nil,
            expectedCode: http.StatusBadRequest,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.testName, func(t *testing.T) {
            mockService := &ServiceMock{}
            mockService.On("PutFile", tc.id, tc.content).Return(tc.serviceErr)

            h := NewHandler(mockService)
            h.registerHandlers()

            requestBody := []byte(`{"content": "` + tc.content + `"}`)
            req, err := http.NewRequest("PUT", "/file/"+tc.id, bytes.NewBuffer(requestBody))
            if err != nil {
                t.Fatal(err)
            }

            rr := httptest.NewRecorder()
            h.Router.ServeHTTP(rr, req)

            assert.Equal(t, tc.expectedCode, rr.Code)
        })
    }
}
