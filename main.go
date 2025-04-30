package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"teste/handlers"
)

var (
	logger *zap.Logger
	port   = getEnv("PORT", "8080")
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		logger.Info("request processed",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.Duration("duration", time.Since(start)),
		)
	})
}

func main() {
	router := mux.NewRouter()

	// Middleware
	router.Use(loggingMiddleware)

	// API v1 routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Initialize handlers
	monsterHandler := handlers.NewMonsterHandler(logger)

	// Monster routes
	api.HandleFunc("/monsters", monsterHandler.GetMonsters).Methods("GET")
	api.HandleFunc("/monsters/{id}", monsterHandler.GetMonster).Methods("GET")

	// Other routes
	api.HandleFunc("/items", monsterHandler.GetItems).Methods("GET")
	api.HandleFunc("/maps", monsterHandler.GetMaps).Methods("GET")
	api.HandleFunc("/skills", monsterHandler.GetSkills).Methods("GET")
	api.HandleFunc("/npcs", monsterHandler.GetNPCs).Methods("GET")

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	// Server configuration
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Info("starting server",
		zap.String("port", port),
	)

	if err := server.ListenAndServe(); err != nil {
		logger.Fatal("server failed to start",
			zap.Error(err),
		)
	}
}
