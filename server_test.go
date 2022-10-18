package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCustomer(t *testing.T) {
	r := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET","/customers",nil)

	r.ServeHTTP(w,req)
	
	assert.Equal(t,http.StatusOK,w.Code)
	assert.NotNil(t,w.Body)



}