package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"teste/config"
	"teste/handlers"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func TestMonsterEndpoint(t *testing.T) {
	logger, _ := zap.NewProduction()
	db, _ := config.NewDatabase(logger)
	handler := handlers.NewMonsterHandler(logger, db)

	req, _ := http.NewRequest("GET", "/api/v1/monsters/1001", nil)
	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/monsters/{id}", handler.GetMonster)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
