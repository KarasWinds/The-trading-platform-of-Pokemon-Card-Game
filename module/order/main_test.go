package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	body := gin.H{
		"trader_id":  "1",
		"card_id":    "1",
		"price":      "2.2",
		"order_type": "buy",
	}

	jsonByte, _ := json.Marshal(body)

	router := NewRouter()

	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("POST", "/api/v1/order", bytes.NewReader(jsonByte))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateErrorOrder(t *testing.T) {
	body := gin.H{
		"trader_id":  "1",
		"card_id":    "1",
		"price":      "22.2",
		"order_type": "buy",
	}
	jsonByte, _ := json.Marshal(body)

	router := NewRouter()

	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("POST", "/api/v1/order", bytes.NewReader(jsonByte))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTraderOrder(t *testing.T) {
	router := NewRouter()

	w := httptest.NewRecorder()
	queryString := "?" + "trader_id=1"
	req, _ := http.NewRequest("GET", "/api/v1/orders"+queryString, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
