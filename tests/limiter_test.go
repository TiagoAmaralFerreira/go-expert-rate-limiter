package tests

import (
	"net/http"
	"net/http/httptest"

	"testing"
	"time"

	"github.com/TiagoAmaralFerreira/go-expert-rate-limiter/configs"
	middleware "github.com/TiagoAmaralFerreira/go-expert-rate-limiter/internal/middlewares"
	"github.com/TiagoAmaralFerreira/go-expert-rate-limiter/internal/repository"
	"github.com/TiagoAmaralFerreira/go-expert-rate-limiter/internal/service"
	"github.com/go-redis/redis/v8"
)

func TestRateLimiterMiddleware(t *testing.T) {
	// Configuração do Redis para teste
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer redisClient.FlushDB(redisClient.Context())

	// Configurações do Rate Limiter
	config := configs.Config{
		MaxRequestsPerSecond: 5,
		BlockDuration:        2 * time.Second,
		RedisClient:          redisClient,
	}

	repo := repository.NewRedisRepository(redisClient)
	rateLimiter := service.NewRateLimiterService(repo, config.MaxRequestsPerSecond, config.BlockDuration)
	rateLimitMiddleware := middleware.NewRateLimitMiddleware(rateLimiter)

	// Test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	server := rateLimitMiddleware.Handle(handler)

	// Enviar 6 requisições para testar o limite
	for i := 1; i <= 6; i++ {
		req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		if i <= 5 && w.Code != http.StatusOK {
			t.Errorf("Request %d: esperava status 200, mas recebeu %d", i, w.Code)
		}

		if i == 6 && w.Code != http.StatusTooManyRequests {
			t.Errorf("Request %d: esperava status 429, mas recebeu %d", i, w.Code)
		}
	}
}

func TestBlockDuration(t *testing.T) {
	// Configuração do Redis para teste
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer redisClient.FlushDB(redisClient.Context())

	// Configurações do Rate Limiter
	config := configs.Config{
		MaxRequestsPerSecond: 5,
		BlockDuration:        2 * time.Second,
		RedisClient:          redisClient,
	}

	repo := repository.NewRedisRepository(redisClient)
	rateLimiter := service.NewRateLimiterService(repo, config.MaxRequestsPerSecond, config.BlockDuration)
	rateLimitMiddleware := middleware.NewRateLimitMiddleware(rateLimiter)

	// Test handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	server := rateLimitMiddleware.Handle(handler)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)

	// Exceder o limite
	for i := 1; i <= 6; i++ {
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)
	}

	// Aguardar duração do bloqueio
	time.Sleep(3 * time.Second)

	// Nova requisição após o bloqueio
	w := httptest.NewRecorder()
	server.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Esperava status 200 após o bloqueio, mas recebeu %d", w.Code)
	}
}
