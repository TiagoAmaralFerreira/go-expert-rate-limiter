package main

import (
	"log"
	"net/http"

	"github.com/TiagoAmaralFerreira/go-expert-rate-limiter/configs"
	handler "github.com/TiagoAmaralFerreira/go-expert-rate-limiter/internal/handlers"
	middleware "github.com/TiagoAmaralFerreira/go-expert-rate-limiter/internal/middlewares"
	"github.com/TiagoAmaralFerreira/go-expert-rate-limiter/internal/repository"
	"github.com/TiagoAmaralFerreira/go-expert-rate-limiter/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	// Carregar configuração.
	cfg, err := configs.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Initialize Redis repository
	redisRepo := repository.NewRedisRepository(cfg.RedisClient)

	// Initialize rate limiter service
	rateLimiterService := service.NewRateLimiterService(redisRepo, cfg.MaxRequestsPerSecond, cfg.BlockDuration)

	// Initialize handler
	limiterHandler := handler.NewLimiterHandler()

	// Initialize middleware
	rateLimitMiddleware := middleware.NewRateLimitMiddleware(rateLimiterService)

	// Setup routes with middleware
	router := mux.NewRouter()
	router.Use(rateLimitMiddleware.Handle) // Apply middleware to all routes
	router.HandleFunc("/", limiterHandler.HandleRequest).Methods("GET")

	// Start HTTP server
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
