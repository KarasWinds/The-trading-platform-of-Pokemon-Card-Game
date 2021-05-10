package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardTrade(t *testing.T) {
	router := NewRouter()

	w := httptest.NewRecorder()
	queryString := "?" + "card_id=1"
	req, _ := http.NewRequest("GET", "/api/v1/trades"+queryString, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
