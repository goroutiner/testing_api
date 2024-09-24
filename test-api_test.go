package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing_api/api"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api.MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotNil(t, responseRecorder.Body)
}

// Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestInvalidCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=kazan", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api.MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(api.MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	stringWithCafe := responseRecorder.Body.String()
	listOfCafe := strings.Split(stringWithCafe, ",")
	assert.Len(t, listOfCafe, totalCount)
}
