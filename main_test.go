package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHealthzRoute(t *testing.T) {
	router := setupRouter()
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/healthz", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "", response.Body.String())
}

func TestGetVehiclesRoute(t *testing.T) {
	router := setupRouter()
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/vehicles", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	expectedBytes, _ := json.Marshal(vehicles)
	actual := response.Body.String()
	require.JSONEq(t, string(expectedBytes), actual)
}

func TestGetVehicleByIDRoute(t *testing.T) {
	router := setupRouter()
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/vehicles/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	expectedBytes, _ := json.Marshal(vehicles[0])
	actual := response.Body.String()
	require.JSONEq(t, string(expectedBytes), actual)
}

func TestPostVehiclesRoute(t *testing.T) {
	router := setupRouter()
	response := httptest.NewRecorder()
	var expectedBytes = []byte(`{
		"id": "6",
		"model": "Kia e-Niro",
		"maker": "Kia"
	}`)
	request, _ := http.NewRequest("POST", "/vehicles", bytes.NewBuffer(expectedBytes))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)
	actual := response.Body.String()
	require.JSONEq(t, string(expectedBytes), actual)
}
