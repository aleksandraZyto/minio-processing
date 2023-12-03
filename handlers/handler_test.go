package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetObjectHandler(t *testing.T) {
	h := NewHandler()
	h.registerHandlers()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	h.Router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
