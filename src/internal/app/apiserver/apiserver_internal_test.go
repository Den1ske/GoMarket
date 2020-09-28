package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleProduct(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/product", nil)
	s.handleProduct().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}